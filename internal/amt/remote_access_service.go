package amt

import (
	"encoding/xml"
	"log"
)

// Global callback for CIRA Credential changes
var OnCiraCredentialsChange func(user, pass string)

type AddMpServer_INPUT struct {
	XMLName    xml.Name `xml:"AddMpServer_INPUT"`
	AccessInfo string   `xml:"AccessInfo"`
	InfoFormat uint8    `xml:"InfoFormat"`
	Port       int      `xml:"Port"`
	AuthMethod uint8    `xml:"AuthMethod"`
	Username   string   `xml:"Username"`
	Password   string   `xml:"Password"`
	CommonName string   `xml:"CN"`
}

type AddMpServer_OUTPUT struct {
	XMLName     xml.Name `xml:"AddMpServer_OUTPUT"`
	ReturnValue int      `xml:"ReturnValue"`
}

func HandleRemoteAccessService(action string, content []byte) []byte {
	switch action {
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddMpServer":
		return handleAddMpServer(content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRule":
		return handleAddRemoteAccessPolicyRule(content)
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return handleRemoteAccessServiceGet()
	default:
		log.Printf("Unhandled RemoteAccessService Action: %s", action)
		return nil
	}
}

type AMT_RemoteAccessService struct {
	XMLName       xml.Name `xml:"AMT_RemoteAccessService"`
	Xmlns         string   `xml:"xmlns,attr"`
	RemoteStatus  int      `xml:"RemoteStatus"`  // 0=Not Connected, 1=Connecting, 2=Connected
	RemoteTrigger int      `xml:"RemoteTrigger"` // 0=User Initiated
	MPSHostname   string   `xml:"MPSHostname"`
}

func handleRemoteAccessServiceGet() []byte {
	status := GetRemoteAccessStatus()
	hostname := "amtsim.local" // Mock or retrieve if AddMpServer stored it

	response := AMT_RemoteAccessService{
		Xmlns:         "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService",
		RemoteStatus:  status,
		RemoteTrigger: 0, // User Initiated default
		MPSHostname:   hostname,
	}

	data, _ := xml.Marshal(response)
	return data
}

func handleAddRemoteAccessPolicyRule(content []byte) []byte {
	// We just return success (0)
	type AddRemoteAccessPolicyRule_OUTPUT struct {
		XMLName     xml.Name `xml:"AddRemoteAccessPolicyRule_OUTPUT"`
		ReturnValue int
	}

	output := AddRemoteAccessPolicyRule_OUTPUT{
		ReturnValue: 0,
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleAddMpServer(content []byte) []byte {
	var input AddMpServer_INPUT
	if err := xml.Unmarshal(content, &input); err != nil {
		log.Printf("Error unmarshalling AddMpServer input: %v", err)
		return nil
	}

	// Notify listener
	if OnCiraCredentialsChange != nil {
		OnCiraCredentialsChange(input.Username, input.Password)
	}

	// Update Remote Status to Connected (2)
	SetRemoteAccessStatus(2)

	output := AddMpServer_OUTPUT{
		ReturnValue: 0, // Success
	}

	resp, _ := xml.Marshal(output)
	return resp
}
