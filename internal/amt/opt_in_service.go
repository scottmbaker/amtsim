package amt

import (
	"encoding/xml"
	"log"
)

type OptInService struct {
	XMLName         xml.Name `xml:"IPS_OptInService"`
	ElementName     string   `xml:"ElementName"`
	InstanceID      string   `xml:"InstanceID"`
	OptInRequired   uint32   `xml:"OptInRequired"`
	OptInState      uint32   `xml:"OptInState"`
	CanModifyOptInPolicy bool `xml:"CanModifyOptInPolicy"`
}

func HandleOptInService(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return handleOptInServiceGet()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put":
		return handleOptInServicePut(content)
	default:
		log.Printf("Unhandled OptInService Action: %s", action)
		return nil
	}
}

func handleOptInServiceGet() []byte {
	item := OptInService{
		ElementName:          "Intel(r) AMT OptIn Service",
		InstanceID:           "Intel(r) AMT:OptInService 0",
		OptInRequired:        0, // None
		OptInState:           0, // Uninitialized
		CanModifyOptInPolicy: true,
	}
	resp, _ := xml.Marshal(item)
	return resp
}

func handleOptInServicePut(content []byte) []byte {
	// Mock success, return what we would return for Get usually
	return handleOptInServiceGet()
}
