package amt

import (
	"encoding/xml"
	"log"
)

type Chassis struct {
	XMLName             xml.Name `xml:"cim:CIM_Chassis"`
	InstanceID          string   `xml:"cim:InstanceID"`
	ElementName         string   `xml:"cim:ElementName"`
	Manufacturer        string   `xml:"cim:Manufacturer"`
	Model               string   `xml:"cim:Model"`
	SerializeNumber     string   `xml:"cim:SerialNumber"` // Note: CIM often uses SerialNumber, checking field name
	LockPresent         bool     `xml:"cim:LockPresent"`
	SecurityStatus      int      `xml:"cim:SecurityStatus"` // 3 = None/Unknown
	ChassisPackageType  int      `xml:"cim:ChassisPackageType"` // 9 = Laptop
}

func HandleChassis(action string, content []byte) []byte {
	log.Printf("HandleChassis Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		resp := Chassis{
			InstanceID:         "Intel(r) AMT: Chassis 0",
			ElementName:        "Intel(r) AMT: Chassis",
			Manufacturer:       "LENOVO",
			Model:              "Wait",
			SerializeNumber:    "PF1XY1Z2",
			LockPresent:        false,
			SecurityStatus:     3,
			ChassisPackageType: 9, // Laptop
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_Chassis Action: %s", action)
		return nil
	}
}
