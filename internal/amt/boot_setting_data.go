package amt

import (
	"encoding/xml"
	"log"
)

type BootSettingData struct {
	XMLName                 xml.Name `xml:"amt:AMT_BootSettingData"`
	InstanceID              string   `xml:"amt:InstanceID"`
	ElementName             string   `xml:"amt:ElementName"`
	OWBEEnabled             bool     `xml:"amt:OWBEEnabled"`
	LinkIsUp                bool     `xml:"amt:LinkIsUp"`
	UEFIModifiable          bool     `xml:"amt:UEFIModifiable"`
	SecureBootControlEnabled bool     `xml:"amt:SecureBootControlEnabled"`
	UEFIHTTPSBootEnabled     bool     `xml:"amt:UEFIHTTPSBootEnabled"`
	WinREBootEnabled         bool     `xml:"amt:WinREBootEnabled"`
	UEFILocalPBABootEnabled  bool     `xml:"amt:UEFILocalPBABootEnabled"`
}

func HandleBootSettingData(action string, content []byte) []byte {
	log.Printf("HandleBootSettingData Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		resp := BootSettingData{
			InstanceID:               "Intel(r) AMT: Boot Setting Data 0",
			ElementName:              "Intel(r) AMT: Boot Setting Data",
			OWBEEnabled:              true,
			LinkIsUp:                 true,
			UEFIModifiable:           true,
			SecureBootControlEnabled: false,
			UEFIHTTPSBootEnabled:     true,
			WinREBootEnabled:         true,
			UEFILocalPBABootEnabled:  true,
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled AMT_BootSettingData Action: %s", action)
		return nil
	}
}
