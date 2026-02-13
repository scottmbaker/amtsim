package cira

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"amtsim/internal/amt"
	"amtsim/internal/amt/consts"
	"amtsim/internal/logging"
	"amtsim/internal/wsman"
)

const CredsFile = "cira_creds.json"

// APF Constants
const (
	APF_SERVICE_REQUEST  = 5
	APF_SERVICE_ACCEPT   = 6
	APF_USERAUTH_REQUEST = 50
	APF_USERAUTH_SUCCESS = 52
	APF_USERAUTH_FAILURE = 51
	APF_CHANNEL_OPEN     = 90 // Server opening channel (e.g. 16992)
)

type Client struct {
	MPSAddress    string
	Username      string
	Password      string
	Uuid          string
	stopChan      chan struct{}
	mu            sync.RWMutex
	channelMap    map[uint32]uint32 // LocalID -> ServerID
	authenticated bool              // Connection-wide auth state
	nextChannelID uint32
	currentConn   *tls.Conn
}

func NewClient(addr, user, pass, uuid string) *Client {
	return &Client{
		MPSAddress:    addr,
		Username:      user,
		Password:      pass,
		Uuid:          uuid,
		stopChan:      make(chan struct{}),
		channelMap:    make(map[uint32]uint32),
		authenticated: false,
		nextChannelID: 1, // Start at 1
	}
}

func (c *Client) Start() {
	logging.Info("Initializing CIRA client for %s", c.MPSAddress)

	// Wire up password change callback from AMT to CIRA
	amt.OnCiraCredentialsChange = c.SetCredentials

	// Wire up Unprovision callback
	amt.OnUnprovision = c.ForceDisconnect

	go c.Run()
}

func (c *Client) ForceDisconnect() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.currentConn != nil {
		logging.Info("Forcing CIRA Disconnection due to Unprovisioning.")
		if err := c.currentConn.Close(); err != nil {
			logging.Error("Failed to close CIRA connection: %v", err)
		}
		c.currentConn = nil
		c.authenticated = false
		amt.SetRemoteAccessStatus(consts.RemoteAccessStatusNotConnected) // Ensure status is updated
	}
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) SetCredentials(user, pass string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Username = user
	c.Password = pass
	logging.Info("CIRA Client credentials updated. User: %s", user)
	c.saveCredentials(user, pass)
}

func (c *Client) saveCredentials(user, pass string) {
	creds := Credentials{Username: user, Password: pass}
	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		logging.Error("Failed to marshal credentials: %v", err)
		return
	}
	if err := os.WriteFile(CredsFile, data, 0600); err != nil {
		logging.Error("Failed to save credentials to %s: %v", CredsFile, err)
	} else {
		logging.Info("Credentials saved to %s", CredsFile)
	}
}

func LoadCredentials() (*Credentials, error) {
	data, err := os.ReadFile(CredsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // No file, not an error
		}
		return nil, err
	}
	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, err
	}
	return &creds, nil
}

func (c *Client) Stop() {
	close(c.stopChan)
}

func (c *Client) Run() {
	logging.Info("Starting CIRA client connecting to %s", c.MPSAddress)
	for {
		select {
		case <-c.stopChan:
			return
		default:
			if amt.GetProvisioningState() != consts.ProvisioningStatePost {
				logging.Info("Device is not provisioned. CIRA will start after activation.")
				time.Sleep(5 * time.Second)
				continue
			}
			if err := c.connectAndServe(); err != nil {
				logging.Error("CIRA error: %v. Retrying in 5s...", err)
				amt.SetRemoteAccessStatus(consts.RemoteAccessStatusNotConnected) // Not Connected
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func (c *Client) connectAndServe() error {
	logging.Info("Connecting to MPS at %s...", c.MPSAddress)
	conn, err := tls.Dial("tcp", c.MPSAddress, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return err
	}

	c.mu.Lock()
	c.currentConn = conn
	c.mu.Unlock()

	defer func() {
		if err := conn.Close(); err != nil {
			logging.Error("Failed to close connection: %v", err)
		}
		c.mu.Lock()
		if c.currentConn == conn {
			c.currentConn = nil
		}
		c.mu.Unlock()
	}()

	// Reset channel map on new connection
	c.mu.Lock()
	c.channelMap = make(map[uint32]uint32)
	c.nextChannelID = 1
	c.authenticated = false
	c.mu.Unlock()

	// Update RAS Status to Connecting
	amt.SetRemoteAccessStatus(consts.RemoteAccessStatusConnecting)

	if err := sendServiceRequest(conn, "auth@amt.intel.com"); err != nil {
		return fmt.Errorf("failed to send service request: %v", err)
	}

	if err := expectServiceAccept(conn, "auth@amt.intel.com"); err != nil {
		return fmt.Errorf("auth service accept failed: %v", err)
	}

	if err := sendProtocolVersion(conn, c.Uuid); err != nil {
		return fmt.Errorf("failed to send protocol version: %v", err)
	}

	c.mu.RLock()
	pass := c.Password
	user := c.Username
	c.mu.RUnlock()

	if pass == "" {
		return fmt.Errorf("CIRA password not set")
	}
	if user == "" {
		return fmt.Errorf("CIRA username not set")
	}

	logging.Info("Attempting CIRA Auth with User: %s, Pass: %s", user, pass)

	if err := sendUserAuthRequest(conn, user, pass); err != nil {
		return fmt.Errorf("failed to send auth request: %v", err)
	}

	if err := expectAuthSuccess(conn); err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	logging.Info("CIRA Authenticated Successfully.")
	// Update RAS Status to Connected
	amt.SetRemoteAccessStatus(consts.RemoteAccessStatusConnected)

	// Buffer for accumulating APF data
	accumulator := make([]byte, 0, 65535)
	readBuf := make([]byte, 4096)

	// 65 seconds to allow for some jitter/latency on the 60s keepalive from MPS
	const readTimeout = 65 * time.Second

	for {
		// Set read deadline to detect broken connections
		if err := conn.SetReadDeadline(time.Now().Add(readTimeout)); err != nil {
			logging.Error("Failed to set read deadline: %v", err)
			return err
		}

		// log.Println("Waiting for data from MPS...")
		n, err := conn.Read(readBuf)
		if err != nil {
			if err == io.EOF {
				logging.Info("MPS closed connection")
				return nil
			}
			// Check for timeout
			if os.IsTimeout(err) {
				logging.Info("CIRA Connection timed out (no data/keepalive in %v). Reconnecting...", readTimeout)
				return err
			}
			logging.Error("Read error: %v", err)
			return err
		}
		if n > 0 {
			logging.Debug("READ %d bytes from MPS", n)
			accumulator = append(accumulator, readBuf[:n]...)

			// Process accumulator
			for len(accumulator) > 0 {
				if len(accumulator) < 1 {
					break
				}
				code := accumulator[0]
				needed := 0

				logging.Debug("Processing Opcode: %d (AccLen: %d)", code, len(accumulator))

				switch code {
				case 90: // CHANNEL_OPEN
					if len(accumulator) < 5 {
						needed = -1 // Wait
					} else {
						// 1 (code) + 4 (typeLen) + typeLen + 4 (chan) + 4 (win) + 4 (max)
						typeLen := binary.BigEndian.Uint32(accumulator[1:])
						minHeader := 1 + 4 + int(typeLen) + 12
						if len(accumulator) < minHeader {
							needed = -1
						} else if len(accumulator) < minHeader+4 { // Wait for TargetLen
							needed = -1
						} else {
							targetLen := binary.BigEndian.Uint32(accumulator[minHeader:])
							intermediate := minHeader + 4 + int(targetLen) + 4
							if len(accumulator) < intermediate+4 { // Wait for SourceLen
								needed = -1
							} else {
								sourceLen := binary.BigEndian.Uint32(accumulator[intermediate:])
								needed = intermediate + 4 + int(sourceLen) + 4
							}
						}
					}

				case 94: // CHANNEL_DATA
					if len(accumulator) < 9 {
						needed = -1
					} else {
						dataLen := binary.BigEndian.Uint32(accumulator[5:])
						needed = 9 + int(dataLen)
					}

				case 97: // CHANNEL_CLOSE
					needed = 5

				case 93: // CHANNEL_WINDOW_ADJUST
					needed = 9

				case 208, 209: // KEEPALIVE
					needed = 5

				case APF_SERVICE_REQUEST, APF_SERVICE_ACCEPT, APF_USERAUTH_REQUEST, APF_USERAUTH_SUCCESS, APF_USERAUTH_FAILURE:
					logging.Debug("Unexpected Handshake Opcode in loop: %d", code)
					needed = 1 // Skip 1 byte to try validation/recovery? Or just break.

				default:
					logging.Error("Unknown APF Opcode: %d. Dropping byte.", code)
					needed = 1
				}

				if needed == -1 || len(accumulator) < needed {
					logging.Debug("Need more data for Opcode %d (Have %d, Need %d+)", code, len(accumulator), needed)
					break // Wait for more data
				}

				logging.Debug("Orchestrating Opcode %d (PacketLen: %d)", code, needed)

				// Extract Packet
				packet := accumulator[:needed]
				accumulator = accumulator[needed:]

				// Helper to invoke handlers
				switch code {
				case 90:
					c.handleChannelOpen(conn, packet)
				case 94:
					logging.Debug("Invoking handleChannelData...")
					c.handleChannelData(conn, packet)
					logging.Debug("Returned from handleChannelData")
				case 97:
					logging.Debug("Received Channel Close (97)")
				case 93:
					logging.Debug("Received Window Adjust (93)")
				case 208:
					// Keepalive Request, send Reply (209)
					reply := []byte{209, 0, 0, 0, 0} // Cookie 0
					if _, err := conn.Write(reply); err != nil {
						logging.Error("Failed to write keepalive reply: %v", err)
					}
				}
			}
		}
	}
}

func (c *Client) handleChannelOpen(conn io.Writer, data []byte) {
	if len(data) < 10 {
		return
	}

	offset := 1
	typeLen := binary.BigEndian.Uint32(data[offset:])
	offset += 4 + int(typeLen)

	senderChannel := binary.BigEndian.Uint32(data[offset:])
	offset += 4
	window := binary.BigEndian.Uint32(data[offset:])

	c.mu.Lock()
	localID := c.nextChannelID
	c.nextChannelID++
	c.channelMap[localID] = senderChannel
	c.mu.Unlock()

	reply := new(bytes.Buffer)
	reply.WriteByte(91) // CHANNEL_OPEN_CONFIRMATION
	if err := binary.Write(reply, binary.BigEndian, senderChannel); err != nil {
		logging.Error("Failed to write sender channel: %v", err)
		return
	}
	if err := binary.Write(reply, binary.BigEndian, localID); err != nil {
		logging.Error("Failed to write local ID: %v", err)
		return
	}
	if err := binary.Write(reply, binary.BigEndian, window); err != nil {
		logging.Error("Failed to write window: %v", err)
		return
	}
	if err := binary.Write(reply, binary.BigEndian, uint32(0)); err != nil {
		logging.Error("Failed to write max packet: %v", err)
		return
	}

	if _, err := conn.Write(reply.Bytes()); err != nil {
		logging.Error("Failed to write channel open confirmation: %v", err)
		return
	}
	logging.Debug("Sent Channel Open Confirmation. ServerChannel: %d -> LocalChannel: %d", senderChannel, localID)
}

func (c *Client) handleChannelData(conn io.Writer, data []byte) {
	logging.Debug("handleChannelData: Start")
	if len(data) < 9 {
		return
	}
	recipient := binary.BigEndian.Uint32(data[1:])
	length := binary.BigEndian.Uint32(data[5:])

	c.mu.RLock()
	serverID, exists := c.channelMap[recipient]
	c.mu.RUnlock()

	if !exists {
		logging.Error("Received CIRA data for unknown local channel: %d", recipient)
		return
	}

	// Send Window Adjust (Opcode 93) to replenish credits on server side
	// This is critical for flow control.
	logging.Debug("handleChannelData: Sending Window Adjust...")
	if err := sendWindowAdjust(conn, serverID, uint32(length)); err != nil {
		logging.Error("Failed to send Window Adjust: %v", err)
		return
	}
	logging.Debug("handleChannelData: Window Adjust Sent")

	// Check Auth State
	c.mu.RLock()
	isAuthenticated := c.authenticated
	c.mu.RUnlock()

	if !isAuthenticated {
		logging.Info("Connection not authenticated. Sending 401 Challenge on Channel %d.", serverID)

		httpResp := new(bytes.Buffer)
		httpResp.WriteString("HTTP/1.1 401 Unauthorized\r\n")
		httpResp.WriteString("Content-Type: text/xml; charset=utf-8\r\n")
		httpResp.WriteString("Www-Authenticate: Digest realm=\"Intel(r) AMT\", nonce=\"testnonce\", qop=\"auth\", opaque=\"testopaque\", algorithm=\"MD5-sess\"\r\n")
		httpResp.WriteString("Transfer-Encoding: chunked\r\n")
		httpResp.WriteString("\r\n")

		httpResp.WriteString("0\r\n\r\n")
		fullResp := httpResp.Bytes()

		reply := new(bytes.Buffer)
		reply.WriteByte(94)
		if err := binary.Write(reply, binary.BigEndian, serverID); err != nil {
			logging.Error("Failed to write server ID: %v", err)
			return
		}
		if err := binary.Write(reply, binary.BigEndian, uint32(len(fullResp))); err != nil {
			logging.Error("Failed to write response length: %v", err)
			return
		}
		reply.Write(fullResp)

		logging.Debug("handleChannelData: Writing 401 Response...")
		if _, err := conn.Write(reply.Bytes()); err != nil {
			logging.Error("Failed to write 401 response: %v", err)
			return
		}

		c.mu.Lock()
		c.authenticated = true
		c.mu.Unlock()
		return
	}

	payload := data[9 : 9+length]
	logging.Debug("Processing CIRA WSMAN from ServerChannel %d (Local %d)", serverID, recipient)

	resp, err := wsman.Process(payload)
	if err == nil && resp != nil {
		// MPS CIRAChannel.ts strictly expects "0\r\n\r\n" (Chunked end) or "</html>" to process data.
		// We must send the response as HTTP Chunked.

		httpResp := new(bytes.Buffer)
		httpResp.WriteString("HTTP/1.1 200 OK\r\n")
		httpResp.WriteString("Content-Type: application/soap+xml; charset=utf-8\r\n")
		httpResp.WriteString("Transfer-Encoding: chunked\r\n")
		httpResp.WriteString("\r\n") // End of Headers

		// Write Chunk
		// standard chunked encoding: HexSize\r\nData\r\n
		fmt.Fprintf(httpResp, "%x\r\n", len(resp))
		httpResp.Write(resp)
		httpResp.WriteString("\r\n")

		// Write Terminal Chunk
		httpResp.WriteString("0\r\n\r\n")

		fullResp := httpResp.Bytes()

		reply := new(bytes.Buffer)
		reply.WriteByte(94)
		if err := binary.Write(reply, binary.BigEndian, serverID); err != nil {
			logging.Error("Failed to write server ID: %v", err)
			return
		}
		if err := binary.Write(reply, binary.BigEndian, uint32(len(fullResp))); err != nil {
			logging.Error("Failed to write response length: %v", err)
			return
		}
		reply.Write(fullResp)

		logging.Debug("handleChannelData: Writing Response...")
		if _, err := conn.Write(reply.Bytes()); err != nil {
			logging.Error("Failed to write response: %v", err)
			return
		}
		logging.Debug("Sent CIRA WSMAN Response (%d bytes) to ServerChannel %d", len(resp), serverID)
		logging.Debug("handleChannelData: Response Written")
	} else if err != nil {
		logging.Error("WSMAN Error: %v", err)
	}
	logging.Debug("handleChannelData: End")
}

func sendWindowAdjust(w io.Writer, recipientChannel uint32, bytesToAdd uint32) error {
	buf := new(bytes.Buffer)
	buf.WriteByte(93) // APF_CHANNEL_WINDOW_ADJUST (Correct Opcode)
	if err := binary.Write(buf, binary.BigEndian, recipientChannel); err != nil {
		logging.Error("Failed to write recipient channel: %v", err)
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, bytesToAdd); err != nil {
		logging.Error("Failed to write bytes to add: %v", err)
		return err
	}
	_, err := w.Write(buf.Bytes())
	return err
}

func sendProtocolVersion(w io.Writer, uuidStr string) error {
	// ProtocolVersion = 192
	// MPS: Major(1), Minor(5), SystemId(UUID at 13)
	// Len >= 93

	// UUID parsing
	// Format: "1d8aa800-918d-11e4-b33f-123400000000"
	cleanUuid := ""
	for _, c := range uuidStr {
		if c != '-' {
			cleanUuid += string(c)
		}
	}

	uuidBytes := make([]byte, 16)
	if len(cleanUuid) == 32 {
		for i := 0; i < 16; i++ {
			var val byte
			if _, err := fmt.Sscanf(cleanUuid[i*2:i*2+2], "%x", &val); err != nil {
				logging.Error("Failed to parse UUID byte: %v", err)
				return err
			}
			uuidBytes[i] = val
		}

		// Apply byte swapping for CIRA/AMT wire format (Little Endian for first 3 parts)
		// Swap Part 1 (Indices 0-3)
		msgUuid := make([]byte, 16)
		copy(msgUuid, uuidBytes)

		msgUuid[0], msgUuid[1], msgUuid[2], msgUuid[3] = uuidBytes[3], uuidBytes[2], uuidBytes[1], uuidBytes[0]
		// Swap Part 2 (Indices 4-5)
		msgUuid[4], msgUuid[5] = uuidBytes[5], uuidBytes[4]
		// Swap Part 3 (Indices 6-7)
		msgUuid[6], msgUuid[7] = uuidBytes[7], uuidBytes[6]

		// Update uuidBytes with the swapped version
		uuidBytes = msgUuid
	}

	buf := new(bytes.Buffer)
	buf.WriteByte(192) // PROTOCOLVERSION
	if err := binary.Write(buf, binary.BigEndian, uint32(1)); err != nil {
		logging.Error("Failed to write major version: %v", err)
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, uint32(0)); err != nil {
		logging.Error("Failed to write minor version: %v", err)
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, uint32(0)); err != nil {
		logging.Error("Failed to write trigger: %v", err)
		return err
	}
	buf.Write(uuidBytes) // 16 bytes at offset 13 (1+4+4+4=13)

	// Padding to reach 93 bytes
	// Current size: 1 + 12 + 16 = 29
	// Need 64 more bytes
	padding := make([]byte, 64)
	buf.Write(padding)

	_, err := w.Write(buf.Bytes())
	return err
}

func writeString(buf *bytes.Buffer, s string) {
	if err := binary.Write(buf, binary.BigEndian, uint32(len(s))); err != nil {
		logging.Error("Failed to write string length: %v", err)
	}
	buf.WriteString(s)
}

func sendServiceRequest(w io.Writer, service string) error {
	buf := new(bytes.Buffer)
	buf.WriteByte(APF_SERVICE_REQUEST)
	writeString(buf, service)
	_, err := w.Write(buf.Bytes())
	return err
}

func expectServiceAccept(r io.Reader, expectedService string) error {
	header := make([]byte, 5) // Code(1) + Len(4)
	if _, err := io.ReadFull(r, header); err != nil {
		return err
	}

	code := header[0]
	if code != APF_SERVICE_ACCEPT {
		return fmt.Errorf("expected SERVICE_ACCEPT (6), got %d", code)
	}

	length := binary.BigEndian.Uint32(header[1:])
	data := make([]byte, length)
	if _, err := io.ReadFull(r, data); err != nil {
		return err
	}

	if string(data) != expectedService {
		logging.Error("Warning: Expected service %s, got %s", expectedService, string(data))
	}
	return nil
}

func sendUserAuthRequest(w io.Writer, user, pass string) error {
	buf := new(bytes.Buffer)
	buf.WriteByte(APF_USERAUTH_REQUEST)
	writeString(buf, user)
	writeString(buf, "pfwd@amt.intel.com") // Service
	writeString(buf, "password")           // Method

	buf.WriteByte(0) // <--- FALSE (Boolean) - The change of password flag?

	writeString(buf, pass)

	_, err := w.Write(buf.Bytes())
	return err
}

func expectAuthSuccess(r io.Reader) error {
	codeBuf := make([]byte, 1)
	if _, err := r.Read(codeBuf); err != nil {
		return err
	}

	if codeBuf[0] == APF_USERAUTH_SUCCESS {
		return nil
	}

	if codeBuf[0] == APF_USERAUTH_FAILURE {
		return fmt.Errorf("received USERAUTH_FAILURE")
	}

	return fmt.Errorf("expected USERAUTH_SUCCESS (52), got %d", codeBuf[0])
}
