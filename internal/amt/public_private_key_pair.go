package amt

import (
	"encoding/xml"
	"log"
)

type PublicPrivateKeyPair struct {
	XMLName      xml.Name `xml:"AMT_PublicPrivateKeyPair"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
}

type PullResponsePublicPrivateKeyPair struct {
	XMLName xml.Name `xml:"PullResponse"`
	Items   []PublicPrivateKeyPair `xml:"Items>AMT_PublicPrivateKeyPair"`
}

func HandlePublicPrivateKeyPair(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handlePublicPrivateKeyPairEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handlePublicPrivateKeyPairPull()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete":
		return handlePublicPrivateKeyPairDelete()
	default:
		log.Printf("Unhandled PublicPrivateKeyPair Action: %s", action)
		return nil
	}
}

func handlePublicPrivateKeyPairEnumerate() []byte {
	output := struct {
		XMLName xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-key-pair",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handlePublicPrivateKeyPairPull() []byte {
	// Return empty list
	output := PullResponsePublicPrivateKeyPair{
		Items: []PublicPrivateKeyPair{},
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handlePublicPrivateKeyPairDelete() []byte {
	output := struct {
		XMLName xml.Name `xml:"DeleteResponse"`
	}{}
	resp, _ := xml.Marshal(output)
	return resp
}
