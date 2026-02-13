package amt

import (
	"encoding/xml"
	"log"
)

type ManagementPresenceRemoteSAP struct {
	XMLName      xml.Name `xml:"AMT_ManagementPresenceRemoteSAP"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
	Name         string   `xml:"Name"`
	AccessInfo   string   `xml:"AccessInfo"`
}

type PullResponseManagementPresenceRemoteSAP struct {
	XMLName xml.Name `xml:"PullResponse"`
	Items   []ManagementPresenceRemoteSAP `xml:"Items>AMT_ManagementPresenceRemoteSAP"`
}

func HandleManagementPresenceRemoteSAP(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handleManagementPresenceRemoteSAPEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handleManagementPresenceRemoteSAPPull()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete":
		return handleManagementPresenceRemoteSAPDelete()
	default:
		log.Printf("Unhandled ManagementPresenceRemoteSAP Action: %s", action)
		return nil
	}
}

func handleManagementPresenceRemoteSAPDelete() []byte {
	// Simple DeleteResponse
	output := struct {
		XMLName xml.Name `xml:"DeleteResponse"`
	}{}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleManagementPresenceRemoteSAPEnumerate() []byte {
	output := struct {
		XMLName xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-remote-sap",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleManagementPresenceRemoteSAPPull() []byte {
	// Return a default Remote SAP
	// Name is critical as RPS uses it for selector in Policy Rule
	item := ManagementPresenceRemoteSAP{
		ElementName: "Intel(r) AMT:Management Presence Server 0",
		InstanceID:  "Intel(r) AMT:Management Presence Server 0",
		Name:        "Intel(r) AMT:Management Presence Server 0",
		AccessInfo:  "192.168.0.1",
	}

	output := PullResponseManagementPresenceRemoteSAP{
		Items: []ManagementPresenceRemoteSAP{item},
	}

	resp, _ := xml.Marshal(output)
	return resp
}
