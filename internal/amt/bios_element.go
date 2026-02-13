package amt

import (
	"encoding/xml"
	"log"
)

type BIOSElement struct {
	XMLName           xml.Name `xml:"cim:CIM_BIOSElement"`
	InstanceID        string   `xml:"cim:InstanceID"`
	ElementName       string   `xml:"cim:ElementName"`
	Name              string   `xml:"cim:Name"`
	Version           string   `xml:"cim:Version"`
	ReleaseDate       string   `xml:"cim:ReleaseDate"` // DateTime format
	Manufacturer      string   `xml:"cim:Manufacturer"`
	PrimaryBIOS       bool     `xml:"cim:PrimaryBIOS"`
}

func HandleBIOSElement(action string, content []byte) []byte {
	log.Printf("HandleBIOSElement Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		resp := BIOSElement{
			InstanceID:   "Intel(r) AMT: BIOS 0",
			ElementName:  "Intel(r) AMT: BIOS",
			Name:         "BIOS",
			Version:      "N32ET75W (1.51 )",
			ReleaseDate:  "20200825000000.000000+000",
			Manufacturer: "LENOVO",
			PrimaryBIOS:  true,
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_BIOSElement Action: %s", action)
		return nil
	}
}
