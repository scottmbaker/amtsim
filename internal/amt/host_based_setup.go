package amt

import (
	"amtsim/internal/amt/consts"
	"encoding/xml"
	"log"
)

type IPS_HostBasedSetupService struct {
	XMLName             xml.Name `xml:"IPS_HostBasedSetupService"`
	Xmlns               string   `xml:"xmlns,attr"`
	AllowedControlModes string   `xml:"AllowedControlModes"` // "1" for CCM, "2" for ACM? checking rps...
	ConfigurationNonce  string   `xml:"ConfigurationNonce"`  // Base64
	CurrentControlMode  int      `xml:"CurrentControlMode"`  // 0=Pre, 1=CCM, 2=ACM
	CertChainStatus     int      `xml:"CertChainStatus"`
}

type Setup_OUTPUT struct {
	XMLName     xml.Name `xml:"Setup_OUTPUT"`
	ReturnValue int      `xml:"ReturnValue"`
}

type Setup_INPUT struct {
	XMLName                    xml.Name `xml:"Setup_INPUT"`
	NetAdminPassEncryptionType int      `xml:"NetAdminPassEncryptionType"`
	NetworkAdminPassword       string   `xml:"NetworkAdminPassword"`
}

var (
	currentMode = 0 // 0: Pre-provisioning, 1: CCM, 2: ACM
)

func HandleHostBasedSetupService(action string, content []byte) []byte {
	switch action {
	case "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/Get":
		return getHostBasedSetupService()
	case "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/Setup":
		return setupHostBasedSetupService(content)
	default:
		return nil
	}
}

func getHostBasedSetupService() []byte {
	// Mock Data
	response := IPS_HostBasedSetupService{
		Xmlns:               "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService",
		AllowedControlModes: "1",                                    // Admin=2, Client=1. RPS checks validation for capabilities.
		ConfigurationNonce:  "QmFzZTY0Tm9uY2VNb2NrMTIzNDU2Nzg5MA==", // "Base64NonceMock1234567890" encoded
		CurrentControlMode:  currentMode,
		CertChainStatus:     0,
	}

	data, _ := xml.Marshal(response)
	return data
}

func setupHostBasedSetupService(content []byte) []byte {
	// Parse Input
	var input Setup_INPUT
	log.Printf("setupHostBasedSetupService Content: %s", string(content))
	if err := xml.Unmarshal(content, &input); err == nil {
		if input.NetworkAdminPassword != "" {
			log.Printf("Received NetworkAdminPassword (len: %d)", len(input.NetworkAdminPassword))
			// NOTE: This is not the CIRA password. Don't try to use it.
		}
	} else {
		log.Printf("Failed to unmarshal Setup_INPUT: %v", err)
	}

	// Simulate successful activation to CCM
	// Host Based Setup typically results in Client Control Mode (4)
	SetProvisioningMode(consts.ProvisioningModeCCM)
	SetProvisioningState(consts.ProvisioningStatePost)

	response := Setup_OUTPUT{
		ReturnValue: 0,
	}

	data, _ := xml.Marshal(response)
	return data
}
