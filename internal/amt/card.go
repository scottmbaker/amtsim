package amt

import (
	"encoding/xml"
	"log"
)

type Card struct {
	XMLName      xml.Name `xml:"cim:CIM_Card"`
	InstanceID   string   `xml:"cim:InstanceID"`
	ElementName  string   `xml:"cim:ElementName"`
	Manufacturer string   `xml:"cim:Manufacturer"`
	Model        string   `xml:"cim:Model"`
	SerialNumber string   `xml:"cim:SerialNumber"`
	Version      string   `xml:"cim:Version"`
}

func HandleCard(action string, content []byte) []byte {
	log.Printf("HandleCard Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		resp := Card{
			InstanceID:   "Intel(r) AMT: Motherboard 0",
			ElementName:  "Intel(r) AMT: Motherboard",
			Manufacturer: "LENOVO",
			Model:        "20N2004GE",
			SerialNumber: "L1HF99E0123",
			Version:      "SDK0T76240 WIN",
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_Card Action: %s", action)
		return nil
	}
}
