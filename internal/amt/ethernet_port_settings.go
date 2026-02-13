package amt

import (
	"encoding/xml"
	"log"
)

type EthernetPortSettings struct {
	XMLName     xml.Name `xml:"AMT_EthernetPortSettings"`
	ElementName string   `xml:"ElementName"`
	InstanceID  string   `xml:"InstanceID"`
	MACAddress  string   `xml:"MACAddress"`
	LinkIsUp    bool     `xml:"LinkIsUp"`
	DHCPEnabled bool     `xml:"DHCPEnabled"`
}

type PullResponseEthernetPortSettings struct {
	XMLName xml.Name               `xml:"PullResponse"`
	Items   []EthernetPortSettings `xml:"Items>AMT_EthernetPortSettings"`
}

func HandleEthernetPortSettings(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handleEthernetPortSettingsEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handleEthernetPortSettingsPull()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put":
		return handleEthernetPortSettingsPut(content)
	default:
		log.Printf("Unhandled EthernetPortSettings Action: %s", action)
		return nil
	}
}

func handleEthernetPortSettingsPut(content []byte) []byte {
	// Simple success response via GetResponse (Put returns the object)
	// Mocking success
	item := EthernetPortSettings{
		ElementName: "Intel(r) AMT Ethernet Port Settings",
		InstanceID:  "Intel(r) AMT:EthernetPortSettings 0",
		MACAddress:  "00-00-00-00-00-00",
		LinkIsUp:    true,
		DHCPEnabled: true,
	}
	resp, _ := xml.Marshal(item)
	return resp
}

func handleEthernetPortSettingsEnumerate() []byte {
	output := struct {
		XMLName            xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string   `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-ethernet-settings",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleEthernetPortSettingsPull() []byte {
	// Return one dummy ethernet port
	item := EthernetPortSettings{
		ElementName: "Intel(r) AMT Ethernet Port Settings",
		InstanceID:  "Intel(r) AMT:EthernetPortSettings 0",
		MACAddress:  "00-00-00-00-00-00",
		LinkIsUp:    true,
		DHCPEnabled: true,
	}

	output := PullResponseEthernetPortSettings{
		Items: []EthernetPortSettings{item},
	}

	resp, _ := xml.Marshal(output)
	return resp
}
