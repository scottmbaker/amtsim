package amt

import (
	"encoding/xml"
	"log"
)

type IEEE8021xSettings struct {
	XMLName      xml.Name `xml:"IPS_IEEE8021xSettings"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
	Enabled      bool     `xml:"Enabled"`
}

func HandleIEEE8021xSettings(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return handleIEEE8021xSettingsGet()
	default:
		log.Printf("Unhandled IEEE8021xSettings Action: %s", action)
		return nil
	}
}

func handleIEEE8021xSettingsGet() []byte {
	// Return disabled 802.1x settings
	item := IEEE8021xSettings{
		ElementName:  "Intel(r) AMT 802.1x Settings",
		InstanceID:   "Intel(r) AMT:IEEE8021xSettings 0",
		Enabled:      false,
	}

	resp, _ := xml.Marshal(item)
	return resp
}
