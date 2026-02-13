package amt

import (
	"amtsim/internal/amt/consts"
	"encoding/xml"
	"log"
)

type SetupAndConfigurationService struct {
	XMLName                 xml.Name `xml:"AMT_SetupAndConfigurationService"`
	Xmlns                   string   `xml:"xmlns,attr"`
	CreationClassName       string   `xml:"CreationClassName"`
	ElementName             string   `xml:"ElementName"`
	Name                    string   `xml:"Name"`
	SystemCreationClassName string   `xml:"SystemCreationClassName"`
	SystemName              string   `xml:"SystemName"`
	InstanceID              string   `xml:"InstanceID"`
	ProvisioningMode        int      `xml:"ProvisioningMode"` // 1 = Admin Control Mode
	ProvisioningState       int      `xml:"ProvisioningState"` // 2 = Post Provisioning
	DhcpDNSSuffix           string   `xml:"DhcpDNSSuffix"`
	TrustedDNSSuffix        string   `xml:"TrustedDNSSuffix"`
}

var OnUnprovision func()

func HandleSetupAndConfigurationService(action string, content []byte) []byte {
	log.Printf("HandleSetupAndConfigurationService Action: %s", action)
	switch action {
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService/Unprovision":
		SetProvisioningMode(consts.ProvisioningModePre)
		SetProvisioningState(consts.ProvisioningStatePre)
		if OnUnprovision != nil {
			OnUnprovision()
		}
		return []byte(`<Unprovision_OUTPUT xmlns="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><ReturnValue>0</ReturnValue></Unprovision_OUTPUT>`)
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		resp := SetupAndConfigurationService{
			Xmlns:                   "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService",
			CreationClassName:       "AMT_SetupAndConfigurationService",
			ElementName:             "Intel(r) AMT Setup and Configuration Service",
			Name:                    "Intel(r) AMT Setup and Configuration Service",
			SystemCreationClassName: "CIM_ComputerSystem",
			SystemName:              "Intel(r) AMT",
			InstanceID:              "Intel(r) AMT Setup and Configuration Service",
			ProvisioningMode:        GetProvisioningMode(),
			ProvisioningState:       GetProvisioningState(),
			DhcpDNSSuffix:           "simulated.local",
			TrustedDNSSuffix:        "simulated.local",
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled AMT_SetupAndConfigurationService Action: %s", action)
		return nil
	}
}
