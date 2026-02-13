package amt

import (
	"encoding/xml"
)

type AMT_GeneralSettings struct {
	XMLName              xml.Name `xml:"AMT_GeneralSettings"`
	Xmlns                string   `xml:"xmlns,attr"`
	AMTNetworkEnabled    int      `xml:"AMTNetworkEnabled"`
	DDNSPeriodicUpdateInterval int `xml:"DDNSPeriodicUpdateInterval"`
	DDNSTTL              int      `xml:"DDNSTTL"`
	DDNSUpdateByDHCPServerEnabled bool `xml:"DDNSUpdateByDHCPServerEnabled"`
	DDNSUpdateEnabled    bool     `xml:"DDNSUpdateEnabled"`
	DHCPv6ConfigurationTimeout int `xml:"DHCPv6ConfigurationTimeout"`
	DigestRealm          string   `xml:"DigestRealm"`
	DomainName           string   `xml:"DomainName"`
	ElementName          string   `xml:"ElementName"`
	HostName             string   `xml:"HostName"`
	HostOSFQDN           string   `xml:"HostOSFQDN"`
	IdleWakeTimeout      int      `xml:"IdleWakeTimeout"`
	InstanceID           string   `xml:"InstanceID"`
	NetworkInterfaceEnabled bool  `xml:"NetworkInterfaceEnabled"`
	PingResponseEnabled  bool     `xml:"PingResponseEnabled"`
	PowerSource          int      `xml:"PowerSource"`
	PreferredAddressFamily int    `xml:"PreferredAddressFamily"`
	PresenceNotificationInterval int `xml:"PresenceNotificationInterval"`
	PrivacyLevel         int      `xml:"PrivacyLevel"`
	RmcpPingResponseEnabled bool  `xml:"RmcpPingResponseEnabled"`
	SharedFQDN           bool     `xml:"SharedFQDN"`
	ThunderboltDockEnabled int    `xml:"ThunderboltDockEnabled"`
	WsmanOnlyMode        bool     `xml:"WsmanOnlyMode"`
	ProvisioningMode     int      `xml:"ProvisioningMode"`
}

func HandleGeneralSettings(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return getGeneralSettings()
	default:
		return nil
	}
}

func getGeneralSettings() []byte {
	// Mock Data based on valid response
	response := AMT_GeneralSettings{
		Xmlns:                         "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings",
		AMTNetworkEnabled:             1,
		DDNSPeriodicUpdateInterval:    1440,
		DDNSTTL:                       900,
		DDNSUpdateByDHCPServerEnabled: true,
		DDNSUpdateEnabled:             false,
		DHCPv6ConfigurationTimeout:    0,
		DigestRealm:                   "Digest:A3829B3827DE4D33D4449B366831FD01",
		DomainName:                    "",
		ElementName:                   "Intel(r) AMT: General Settings",
		HostName:                      "",
		HostOSFQDN:                    "AMTSIM-HOST",
		IdleWakeTimeout:               1,
		InstanceID:                    "Intel(r) AMT: General Settings",
		NetworkInterfaceEnabled:       true,
		PingResponseEnabled:           true,
		PowerSource:                   0,
		PreferredAddressFamily:        0,
		PresenceNotificationInterval:  0,
		PrivacyLevel:                  0,
		RmcpPingResponseEnabled:       true,
		SharedFQDN:                    true,
		ThunderboltDockEnabled:        0,
		WsmanOnlyMode:                 false,
		ProvisioningMode:              GetProvisioningMode(),
	}

	data, _ := xml.Marshal(response)
	return data
}
