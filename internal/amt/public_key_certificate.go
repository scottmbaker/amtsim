package amt

import (
	"encoding/xml"
	"log"
)

type PublicKeyCertificate struct {
	XMLName      xml.Name `xml:"AMT_PublicKeyCertificate"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
	X509Certificate string `xml:"X509Certificate"`
}

type PullResponsePublicKeyCertificate struct {
	XMLName xml.Name `xml:"PullResponse"`
	Items   []PublicKeyCertificate `xml:"Items>AMT_PublicKeyCertificate"`
}

func HandlePublicKeyCertificate(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handlePublicKeyCertificateEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handlePublicKeyCertificatePull()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete":
		return handlePublicKeyCertificateDelete()
	default:
		log.Printf("Unhandled PublicKeyCertificate Action: %s", action)
		return nil
	}
}

func handlePublicKeyCertificateEnumerate() []byte {
	output := struct {
		XMLName xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-public-key-cert",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handlePublicKeyCertificatePull() []byte {
	// Return empty list
	output := PullResponsePublicKeyCertificate{
		Items: []PublicKeyCertificate{},
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handlePublicKeyCertificateDelete() []byte {
	output := struct {
		XMLName xml.Name `xml:"DeleteResponse"`
	}{}
	resp, _ := xml.Marshal(output)
	return resp
}
