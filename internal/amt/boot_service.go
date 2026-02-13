package amt

import (
	"encoding/xml"
	"log"
)

type BootService struct {
	XMLName     xml.Name `xml:"cim:CIM_BootService"`
	InstanceID  string   `xml:"cim:InstanceID"`
	ElementName string   `xml:"cim:ElementName"`
}

type BootServiceGetResponse struct {
	XMLName     xml.Name    `xml:"cim:CIM_BootService"`
	InstanceID  string      `xml:"cim:InstanceID"`
	ElementName string      `xml:"cim:ElementName"`
	Name        string      `xml:"cim:Name"`
}

func HandleBootService(action string, content []byte) []byte {
	log.Printf("HandleBootService Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		resp := BootServiceGetResponse{
			InstanceID:  "Intel(r) AMT: Boot Service 0",
			ElementName: "Intel(r) AMT: Boot Service",
			Name:        "Intel(r) AMT: Boot Service",
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_BootService Action: %s", action)
		return nil
	}
}
