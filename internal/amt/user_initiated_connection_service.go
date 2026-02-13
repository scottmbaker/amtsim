package amt

import (
	"encoding/xml"
	"log"
)

type UserInitiatedConnectionService struct {
	XMLName xml.Name `xml:"AMT_UserInitiatedConnectionService"`
}

func HandleUserInitiatedConnectionService(action string, content []byte) []byte {
	switch action {
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService/RequestStateChange":
		return handleUserInitiatedConnectionServiceRequestStateChange(content)
	default:
		log.Printf("Unhandled UserInitiatedConnectionService Action: %s", action)
		return nil
	}
}

func handleUserInitiatedConnectionServiceRequestStateChange(content []byte) []byte {
	// RequestStateChange Output
	// ReturnValue 0 = Success
	output := struct {
		XMLName     xml.Name `xml:"RequestStateChange_OUTPUT"`
		ReturnValue int      `xml:"ReturnValue"`
	}{
		ReturnValue: 0,
	}
	resp, _ := xml.Marshal(output)
	return resp
}
