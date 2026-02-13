package amt

import (
	"encoding/xml"
	"log"
)

type TLSSettingData struct {
	XMLName      xml.Name `xml:"AMT_TLSSettingData"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
	Enabled      bool     `xml:"Enabled"`
	MutualAuthentication bool `xml:"MutualAuthentication"`
}

type PullResponseTLSSettingData struct {
	XMLName xml.Name `xml:"PullResponse"`
	Items   []TLSSettingData `xml:"Items>AMT_TLSSettingData"`
}

func HandleTLSSettingData(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handleTLSSettingDataEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handleTLSSettingDataPull()
	default:
		log.Printf("Unhandled TLSSettingData Action: %s", action)
		return nil
	}
}

func handleTLSSettingDataEnumerate() []byte {
	output := struct {
		XMLName xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-tls-settings",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleTLSSettingDataPull() []byte {
	// Return default TLS settings (disabled for sim)
	// We MUST return 2 items to ensure RPS treats it as an array!
	// RPS accesses [0] and [1] blindly.
	item1 := TLSSettingData{
		ElementName:  "Intel(r) AMT LMS TLS Settings",
		InstanceID:   "Intel(r) AMT:TLSSettingData 0",
		Enabled:      false,
		MutualAuthentication: false,
	}

	item2 := TLSSettingData{
		ElementName:  "Intel(r) AMT 802.3 TLS Settings",
		InstanceID:   "Intel(r) AMT:TLSSettingData 1",
		Enabled:      false,
		MutualAuthentication: false,
	}

	output := PullResponseTLSSettingData{
		Items: []TLSSettingData{item1, item2},
	}

	resp, _ := xml.Marshal(output)
	return resp
}
