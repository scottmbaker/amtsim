package amt

import (
	"encoding/xml"
	"log"
)

type IPS_PowerManagementService struct {
	XMLName                 xml.Name `xml:"ips:IPS_PowerManagementService"`
	ElementName             string   `xml:"ips:ElementName"`
	InstanceID              string   `xml:"ips:InstanceID"`
	ClassName               string   `xml:"ips:CreationClassName"`
	Name                    string   `xml:"ips:Name"`
	SystemCreationClassName string   `xml:"ips:SystemCreationClassName"`
	SystemName              string   `xml:"ips:SystemName"`
	OSPowerSavingState      int      `xml:"ips:OSPowerSavingState"`
}

func HandleIPSPowerManagementService(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return handleIPSPowerManagementServiceGet()
	default:
		log.Printf("Unhandled IPS_PowerManagementService Action: %s", action)
		return nil
	}
}

func handleIPSPowerManagementServiceGet() []byte {
	item := IPS_PowerManagementService{
		ElementName:             "Intel(r) AMT Power Management Service",
		InstanceID:              "Intel(r) AMT:IPS_PowerManagementService 1",
		ClassName:               "IPS_PowerManagementService",
		Name:                    "Intel(r) AMT Power Management Service",
		SystemCreationClassName: "CIM_ComputerSystem",
		SystemName:              "Intel(r) AMT",
		OSPowerSavingState:      1,
	}
	resp, _ := xml.Marshal(item)
	return resp
}
