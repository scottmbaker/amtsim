package amt

import (
	"encoding/xml"
	"log"
)

type ComputerSystemPackage struct {
	XMLName             xml.Name `xml:"cim:CIM_ComputerSystemPackage"`
	InstanceID          string   `xml:"cim:InstanceID"`
	ElementName         string   `xml:"cim:ElementName"`
	PlatformGUID        string   `xml:"cim:PlatformGUID"`
}

func HandleComputerSystemPackage(action string, content []byte) []byte {
	log.Printf("HandleComputerSystemPackage Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		resp := ComputerSystemPackage{
			InstanceID:   "Intel(r) AMT: System Package 0",
			ElementName:  "Intel(r) AMT: System Package",
			PlatformGUID: "00000000-0000-0000-0000-000000000000", // Would match UUID usually
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_ComputerSystemPackage Action: %s", action)
		return nil
	}
}
