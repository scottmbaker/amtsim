package amt

import (
	"encoding/xml"
	"log"
)

type TLSCredentialContext struct {
	XMLName      xml.Name `xml:"AMT_TLSCredentialContext"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
}

type PullResponseTLSCredentialContext struct {
	XMLName xml.Name `xml:"PullResponse"`
	Items   []TLSCredentialContext `xml:"Items>AMT_TLSCredentialContext"`
}

func HandleTLSCredentialContext(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handleTLSCredentialContextEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handleTLSCredentialContextPull()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete":
		return handleTLSCredentialContextDelete()
	default:
		log.Printf("Unhandled TLSCredentialContext Action: %s", action)
		return nil
	}
}

func handleTLSCredentialContextEnumerate() []byte {
	output := struct {
		XMLName xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-tls-cred-context",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleTLSCredentialContextPull() []byte {
	// Return empty list
	output := PullResponseTLSCredentialContext{
		Items: []TLSCredentialContext{},
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleTLSCredentialContextDelete() []byte {
	output := struct {
		XMLName xml.Name `xml:"DeleteResponse"`
	}{}
	resp, _ := xml.Marshal(output)
	return resp
}
