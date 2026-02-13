package amt

import (
	"encoding/xml"
	"log"
)

type RedirectionService struct {
	XMLName            xml.Name `xml:"AMT_RedirectionService"`
	ElementName        string   `xml:"ElementName"`
	InstanceID         string   `xml:"InstanceID"`
	ListenerEnabled    bool     `xml:"ListenerEnabled"`
	AccessLogOutput    []string `xml:"AccessLogOutput"`
}

func HandleRedirectionService(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return handleRedirectionServiceGet()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put":
		return handleRedirectionServicePut(content)
	default:
		log.Printf("Unhandled RedirectionService Action: %s", action)
		return nil
	}
}

func handleRedirectionServiceGet() []byte {
	item := RedirectionService{
		ElementName:     "Intel(r) AMT Redirection Service",
		InstanceID:      "Intel(r) AMT:RedirectionService 0",
		ListenerEnabled: true,
	}
	resp, _ := xml.Marshal(item)
	return resp
}

func handleRedirectionServicePut(content []byte) []byte {
	// Mock success
	return handleRedirectionServiceGet()
}
