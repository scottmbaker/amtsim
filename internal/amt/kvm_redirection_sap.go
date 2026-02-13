package amt

import (
	"encoding/xml"
	"log"
)

type KVMRedirectionSAP struct {
	XMLName      xml.Name `xml:"CIM_KVMRedirectionSAP"`
	Xmlns        string   `xml:"xmlns,attr"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
	EnabledState uint16   `xml:"EnabledState"` // 2=Enabled, 3=Disabled
}

func HandleKVMRedirectionSAP(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return handleKVMRedirectionSAPGet()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put":
		return handleKVMRedirectionSAPPut(content)
	default:
		log.Printf("Unhandled KVMRedirectionSAP Action: %s", action)
		return nil
	}
}

func handleKVMRedirectionSAPGet() []byte {
	item := KVMRedirectionSAP{
		Xmlns:        "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP",
		ElementName:  "Intel(r) AMT KVM Redirection SAP",
		InstanceID:   "Intel(r) AMT:KVMRedirectionSAP 0",
		EnabledState: 2, // Enabled
	}
	resp, _ := xml.Marshal(item)
	return resp
}

func handleKVMRedirectionSAPPut(content []byte) []byte {
	// Mock success
	return handleKVMRedirectionSAPGet()
}
