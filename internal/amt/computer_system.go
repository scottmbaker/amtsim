package amt

import (
	"encoding/xml"
)

const (
	ActionGet = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	ActionPut = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put"
)

type ComputerSystem struct {
	XMLName             xml.Name `xml:"CIM_ComputerSystem"`
	Xmlns               string   `xml:"xmlns,attr"`
	ElementName         string   `xml:"ElementName"`
	Name                string   `xml:"Name"`
	CreationClassName   string   `xml:"CreationClassName"`
	NameFormat          string   `xml:"NameFormat"`
	Dedicated           int      `xml:"Dedicated"`
}

func HandleComputerSystem(action string, body []byte) []byte {
	switch action {
	case ActionGet:
		cs := ComputerSystem{
			Xmlns:             "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem",
			ElementName:       "Managed Client",
			Name:              GetUUID(),
			CreationClassName: "CIM_ComputerSystem",
			NameFormat:        "IP",
			Dedicated:         0,
		}
		data, _ := xml.Marshal(cs)
		return data
	default:
		return nil
	}
}
