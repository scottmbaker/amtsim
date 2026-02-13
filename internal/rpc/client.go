package rpc

import (
	"amtsim/internal/logging"
	"amtsim/internal/wsman"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Client struct {
	URL        string
	Profile    string
	Password   string
	NoVerify   bool
	DeviceInfo MessagePayload
}

func NewClient(url, profile, password string, noVerify bool) *Client {
	c := &Client{
		URL:      url,
		Profile:  profile,
		Password: password,
		NoVerify: noVerify,
	}

	// "don't change this" comments because antigravity was revering the mock data

	// Mock Device Info (Ideally query amtsim, but hardcoding for sim speed)
	c.DeviceInfo = MessagePayload{
		Version:           "16.0.0",
		Build:             "3000",                                 // don't change this
		SKU:               "8",                                    // don't change this
		UUID:              "1d8aa800-918d-11e4-b33f-123400000000", // don't change this
		Username:          "$$OsAdmin",                            // don't change this
		Password:          c.Password,
		CurrentMode:       0, // Pre-provisioning
		Hostname:          "amtsim",
		FQDN:              "amtsim.local",
		Client:            "RPC",
		CertificateHashes: []string{},
		IPConfiguration: IPConfiguration{
			DHCP:       false, // don't change this
			Static:     false, // don't change this
			IPSync:     false, // don't change this
			IpAddress:  "",    // don't change this
			Netmask:    "",    // don't change this
			Gateway:    "",    // don't change this
			PrimaryDns: "",    // don't change this
		},
		HostnameInfo: HostnameInfo{
			DnsSuffixOS: "", // don't change this
			Hostname:    "", // don't change this
		},
		Profile: c.Profile,
	}
	return c
}

type Message struct {
	Method          string `json:"method"`
	APIKey          string `json:"apiKey"`
	AppVersion      string `json:"appVersion"`
	ProtocolVersion string `json:"protocolVersion"`
	Status          string `json:"status"`
	Message         string `json:"message"`
	Payload         string `json:"payload"` // Base64
}

type IPConfiguration struct {
	DHCP       bool   `json:"dhcp"`
	Static     bool   `json:"static"`
	IPSync     bool   `json:"ipsync"`
	IpAddress  string `json:"ipAddress"`
	Netmask    string `json:"netmask"`
	Gateway    string `json:"gateway"`
	PrimaryDns string `json:"primaryDns"`
}

type HostnameInfo struct {
	DnsSuffixOS string `json:"dnsSuffixOS"`
	Hostname    string `json:"hostname"`
}

type MessagePayload struct {
	Version           string          `json:"ver"`
	Build             string          `json:"build"`
	SKU               string          `json:"sku"`
	Features          string          `json:"features"`
	UUID              string          `json:"uuid"`
	Username          string          `json:"username"`
	Password          string          `json:"password"`
	CurrentMode       int             `json:"currentMode"`
	Hostname          string          `json:"hostname"`
	FQDN              string          `json:"fqdn"`
	Client            string          `json:"client"`
	CertificateHashes []string        `json:"certHashes"`
	IPConfiguration   IPConfiguration `json:"ipConfiguration"`
	HostnameInfo      HostnameInfo    `json:"hostnameInfo"`
	Profile           string          `json:"profile"`
}

type AMT_GeneralSettings struct {
	XMLName          xml.Name `xml:"AMT_GeneralSettings"`
	ProvisioningMode int      `xml:"ProvisioningMode"`
}

func (c *Client) Execute(command string) error {
	// 1. Connect
	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.NoVerify},
	}

	log.Printf("Connecting to %s...", c.URL)
	conn, _, err := dialer.Dial(c.URL, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close connection: %v", err)
		}
	}()

	// 2. Mock Device Info (Ideally query amtsim, but hardcoding for sim speed)
	// Now stored in c.DeviceInfo

	fullCommand := command
	if strings.HasPrefix(command, "activate") && c.Profile != "" {
		fullCommand += " --profile " + c.Profile
	} else if strings.HasPrefix(command, "deactivate") {
		// If deactivating, we must simulate being in a provisioned state (ACM=1)
		// otherwise RPS rejects the request with "in pre-provisioning mode".
		c.DeviceInfo.CurrentMode = 1
	}

	payloadBytes, _ := json.Marshal(c.DeviceInfo)
	logging.Debug("Outgoing Payload: %s", string(payloadBytes))

	msg := Message{
		Method:          fullCommand,
		APIKey:          "key",
		AppVersion:      "2.0.0",
		ProtocolVersion: "4.0.0",
		Status:          "ok",
		Message:         "ok",
		Payload:         base64.StdEncoding.EncodeToString(payloadBytes),
	}

	if err := conn.WriteJSON(msg); err != nil {
		return err
	}

	// 3. Loop
	for {
		var resp Message
		if err := conn.ReadJSON(&resp); err != nil {
			return err
		}

		func() {
			// define a temporary struct to capture everything for debugging
			var raw map[string]interface{}
			_ = json.Unmarshal([]byte(resp.Message), &raw)
		}()
		log.Printf("Received Method: %s Status: %s", resp.Method, resp.Status)

		if resp.Method == "error" {
			log.Printf("RPS Error Message: %s", resp.Message)
			// Decode payload if present
			if resp.Payload != "" {
				decoded, _ := base64.StdEncoding.DecodeString(resp.Payload)
				log.Printf("RPS Error Payload: %s", string(decoded))
			}
		}

		if resp.Method == "success" {
			log.Println("RPS indicates Success!")
			return nil
		}

		if resp.Method == "heartbeat_request" {
			log.Println("Received Heartbeat Request. Sending Response...")
			resp.Method = "heartbeat_response"
			resp.Status = "success"
			if err := conn.WriteJSON(resp); err != nil {
				return err
			}
			continue
		}

		if resp.Payload != "" {
			data, _ := base64.StdEncoding.DecodeString(resp.Payload)
			// Remove legacy string check if method is the standard way

			logging.Debug("Executing Payload: %s", string(data))

			var env wsman.Envelope
			if err := xml.Unmarshal(data, &env); err == nil {
				logging.Info("Processing WSMAN Action: %s", env.Header.Action)
				logging.Info("ResourceURI: %s", env.Header.ResourceURI)
			}

			// Proxy to AMT (WSMAN)
			// Assuming payload is raw WSMAN XML

			// We can use http to post to localhost:16992/wsman
			// Assumes amtsim is running

			var wsmanResp []byte
			wsmanResp, err = c.PostToAmtSim(data)
			if err != nil {
				log.Printf("AMT Simulation Error: %v", err)
				// Send empty response?
				wsmanResp = []byte("")
			}

			// Format WSMAN response as Chunked Transfer Encoding for RPS
			// Header
			header := "HTTP/1.1 200 OK\r\nContent-Type: application/soap+xml; charset=utf-8\r\nTransfer-Encoding: chunked\r\n\r\n"

			// Chunked Body: <Length in Hex>\r\n<Data>\r\n0\r\n\r\n
			chunkLen := fmt.Sprintf("%x", len(wsmanResp))
			chunkedBody := fmt.Sprintf("%s\r\n%s\r\n0\r\n\r\n", chunkLen, string(wsmanResp))

			fullResponse := header + chunkedBody

			// Send response back to RPS
			reply := Message{
				Method:          "response",
				APIKey:          "key",
				AppVersion:      "2.0.0",
				ProtocolVersion: "4.0.0",
				Status:          "ok",
				Message:         "ok",
				Payload:         base64.StdEncoding.EncodeToString([]byte(fullResponse)),
			}

			logging.Debug("Sending reply %+v\n", reply)

			if err := conn.WriteJSON(reply); err != nil {
				return err
			}
		}
	}
}

func (c *Client) PostToAmtSim(data []byte) ([]byte, error) {
	// First attempt (will likely fail with 401)
	// Use c.URL if set and http/s, otherwise default to localhost:16992
	targetUrl := "http://localhost:16992/wsman"
	if strings.HasPrefix(c.URL, "http") {
		// If c.URL ends with /wsman, use it, else append
		if strings.HasSuffix(c.URL, "/wsman") {
			targetUrl = c.URL
		} else {
			targetUrl = c.URL + "/wsman"
		}
	}

	req, _ := http.NewRequest("POST", targetUrl, strings.NewReader(string(data)))
	req.Header.Set("Content-Type", "application/soap+xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()

	if resp.StatusCode == http.StatusUnauthorized {
		// Handle Digest Auth
		authHeader := resp.Header.Get("WWW-Authenticate")
		if authHeader == "" {
			return nil, fmt.Errorf("401 response missing WWW-Authenticate header")
		}

		digestParts := parseDigestHeader(authHeader)

		realm := digestParts["realm"]
		nonce := digestParts["nonce"]
		qop := digestParts["qop"]
		opaque := digestParts["opaque"]
		algorithm := digestParts["algorithm"]
		if algorithm == "" {
			algorithm = "MD5"
		}

		// Calculate response
		// HA1 = MD5(username:realm:password)
		ha1 := md5Hex("admin:" + realm + ":password")
		// HA2 = MD5(method:digestURI)
		ha2 := md5Hex("POST:/wsman")
		// Response = MD5(HA1:nonce:nc:cnonce:qop:HA2)

		nc := "00000001"
		cnonce := "abcdef0123456789" // random-ish

		response := md5Hex(ha1 + ":" + nonce + ":" + nc + ":" + cnonce + ":" + qop + ":" + ha2)

		authVal := fmt.Sprintf(`Digest username="admin", realm="%s", nonce="%s", uri="/wsman", qop=%s, nc=%s, cnonce="%s", response="%s", opaque="%s", algorithm="%s"`,
			realm, nonce, qop, nc, cnonce, response, opaque, algorithm)

		// Retry with auth
		req2, _ := http.NewRequest("POST", targetUrl, strings.NewReader(string(data)))
		req2.Header.Set("Content-Type", "application/soap+xml")
		req2.Header.Set("Authorization", authVal)

		resp2, err := client.Do(req2)
		if err != nil {
			return nil, err
		}

		defer func() {
			if err := resp2.Body.Close(); err != nil {
				fmt.Printf("error closing response body: %v\n", err)
			}
		}()

		if resp2.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp2.Body)
			return nil, fmt.Errorf("AMTSim returned %d after auth: %s", resp2.StatusCode, string(body))
		}

		return io.ReadAll(resp2.Body)
	}

	// Check status code for initial request (if it passed without auth?)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("AMTSim returned %d: %s", resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}

func parseDigestHeader(header string) map[string]string {
	// Remove "Digest " prefix
	header = strings.TrimPrefix(header, "Digest ")
	parts := strings.Split(header, ",")
	result := make(map[string]string)
	for _, part := range parts {
		kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(kv) == 2 {
			key := kv[0]
			val := strings.Trim(kv[1], "\"")
			result[key] = val
		}
	}
	return result
}

func md5Hex(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (c *Client) GetGeneralSettings() (AMT_GeneralSettings, error) {
	// Construct WSMAN Get Request
	resourceURI := "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings"
	action := "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"

	header := fmt.Sprintf(`<s:Header>
    <a:Action>%s</a:Action>
    <a:To>%s/wsman</a:To>
    <w:ResourceURI>%s</w:ResourceURI>
    <a:MessageID>%s</a:MessageID>
    <a:ReplyTo>
      <a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address>
    </a:ReplyTo>
  </s:Header>`, action, c.URL, resourceURI, "uuid:"+uuidStr()) // simplify uuid generation or import

	body := `<s:Body/>`

	envelope := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">
  %s
  %s
</s:Envelope>`, header, body)

	respBytes, err := c.PostToAmtSim([]byte(envelope))
	if err != nil {
		return AMT_GeneralSettings{}, err
	}

	// Parse Response
	// We need to unwrap the Envelope/Body/AMT_GeneralSettings
	type ResponseEnvelope struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			GeneralSettings AMT_GeneralSettings `xml:"AMT_GeneralSettings"`
		} `xml:"Body"`
	}

	var respEnvelope ResponseEnvelope
	if err := xml.Unmarshal(respBytes, &respEnvelope); err != nil {
		return AMT_GeneralSettings{}, err
	}

	return respEnvelope.Body.GeneralSettings, nil
}

type AMT_RemoteAccessService struct {
	XMLName       xml.Name `xml:"AMT_RemoteAccessService"`
	RemoteStatus  int      `xml:"RemoteStatus"`
	RemoteTrigger int      `xml:"RemoteTrigger"`
	MPSHostname   string   `xml:"MPSHostname"`
}

func (c *Client) GetRemoteAccessConnectionStatus() (AMT_RemoteAccessService, error) {
	// Construct WSMAN Get Request
	resourceURI := "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService"
	action := "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"

	header := fmt.Sprintf(`<s:Header>
    <a:Action>%s</a:Action>
    <a:To>%s/wsman</a:To>
    <w:ResourceURI>%s</w:ResourceURI>
    <a:MessageID>%s</a:MessageID>
    <a:ReplyTo>
      <a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address>
    </a:ReplyTo>
  </s:Header>`, action, c.URL, resourceURI, "uuid:"+uuidStr())

	body := `<s:Body/>`

	envelope := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">
  %s
  %s
</s:Envelope>`, header, body)

	respBytes, err := c.PostToAmtSim([]byte(envelope))
	if err != nil {
		return AMT_RemoteAccessService{}, err
	}

	type ResponseEnvelope struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			RemoteAccessService AMT_RemoteAccessService `xml:"AMT_RemoteAccessService"`
		} `xml:"Body"`
	}

	var respEnvelope ResponseEnvelope
	if err := xml.Unmarshal(respBytes, &respEnvelope); err != nil {
		return AMT_RemoteAccessService{}, err
	}

	return respEnvelope.Body.RemoteAccessService, nil
}

type CIM_ComputerSystem struct {
	XMLName xml.Name `xml:"CIM_ComputerSystem"`
	Name    string   `xml:"Name"` // UUID
}

func (c *Client) GetComputerSystem() (CIM_ComputerSystem, error) {
	resourceURI := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem"
	action := "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"

	header := fmt.Sprintf(`<s:Header>
    <a:Action>%s</a:Action>
    <a:To>%s/wsman</a:To>
    <w:ResourceURI>%s</w:ResourceURI>
    <a:MessageID>%s</a:MessageID>
    <a:ReplyTo>
      <a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address>
    </a:ReplyTo>
  </s:Header>`, action, c.URL, resourceURI, "uuid:"+uuidStr())

	body := `<s:Body/>`

	envelope := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">
  %s
  %s
</s:Envelope>`, header, body)

	respBytes, err := c.PostToAmtSim([]byte(envelope))
	if err != nil {
		return CIM_ComputerSystem{}, err
	}

	type ResponseEnvelope struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			ComputerSystem CIM_ComputerSystem `xml:"CIM_ComputerSystem"`
		} `xml:"Body"`
	}

	var respEnvelope ResponseEnvelope
	if err := xml.Unmarshal(respBytes, &respEnvelope); err != nil {
		return CIM_ComputerSystem{}, err
	}

	return respEnvelope.Body.ComputerSystem, nil
}

// Helper for uuid
func uuidStr() string {
	// simple mock uuid for now or import google/uuid if available
	return "12345678-1234-5678-1234-567812345678"
}
