package rpc

import (
	"amtsim/internal/amt/consts"
	"fmt"
)

func (c *Client) AmtInfo() error {
	d := c.DeviceInfo

	// Fetch General Settings from AMT (simulated via localhost WSMAN)
	// We need to set the URL for the client if it wasn't set (amtinfo case)
	if c.URL == "" {
		c.URL = "http://localhost:16992"
	}

	generalSettings, err := c.GetGeneralSettings()
	var controlMode string
	if err == nil {
		switch generalSettings.ProvisioningMode {
		case consts.ProvisioningModePre:
			controlMode = "pre-provisioning"
		case consts.ProvisioningModeACM:
			controlMode = "activated in admin control mode"
		case consts.ProvisioningModeCCM:
			controlMode = "activated in client control mode"
		default:
			controlMode = fmt.Sprintf("unknown (%d)", generalSettings.ProvisioningMode)
		}
	} else {
		controlMode = fmt.Sprintf("error fetching: %v", err)
	}

	// Fetch UUID from CIM_ComputerSystem
	var uuidStatus string
	cs, err := c.GetComputerSystem()
	if err == nil {
		uuidStatus = cs.Name
	} else {
		uuidStatus = fmt.Sprintf("error: %v", err)
	}

	var rasStatus string
	mpsHostname := c.URL // Fallback
	var rasTrigger string

	ras, err := c.GetRemoteAccessConnectionStatus()
	if err == nil {
		switch ras.RemoteStatus {
		case consts.RemoteAccessStatusNotConnected:
			rasStatus = "not connected"
		case consts.RemoteAccessStatusConnecting:
			rasStatus = "connecting"
		case consts.RemoteAccessStatusConnected:
			rasStatus = "connected"
		default:
			rasStatus = fmt.Sprintf("unknown (%d)", ras.RemoteStatus)
		}

		if ras.RemoteTrigger == 0 {
			rasTrigger = "user initiated"
		} else {
			rasTrigger = fmt.Sprintf("unknown (%d)", ras.RemoteTrigger)
		}

		if ras.MPSHostname != "" {
			mpsHostname = ras.MPSHostname
		}
	} else {
		rasStatus = fmt.Sprintf("error: %v", err)
	}

	fmt.Printf("Version                 : %s\n", d.Version)
	fmt.Printf("Build Number            : %s\n", d.Build)
	fmt.Printf("SKU                     : %s\n", d.SKU)
	fmt.Printf("Features                : AMT Pro\n") // Hardcoded based on SKU=8 usually
	fmt.Printf("UUID                    : %s\n", uuidStatus)
	fmt.Printf("Control Mode            : %s\n", controlMode)
	fmt.Printf("DNS Suffix              : %s\n", "example.com") // Mock
	fmt.Printf("DNS Suffix (OS)         : %s\n", d.HostnameInfo.DnsSuffixOS)
	fmt.Printf("Hostname (OS)           : %s\n", d.HostnameInfo.Hostname)
	fmt.Printf("RAS Network             : direct\n")
	fmt.Printf("RAS Remote Status       : %s\n", rasStatus)
	fmt.Printf("RAS Trigger             : %s\n", rasTrigger)
	fmt.Printf("RAS MPS Hostname        : %s\n", mpsHostname)
	fmt.Println("---Wired Adapter---")
	fmt.Printf("DHCP Enabled            : %t\n", !d.IPConfiguration.Static)
	fmt.Printf("DHCP Mode               : %s\n", "passive")
	fmt.Printf("Link Status             : up\n")
	fmt.Printf("AMT IP Address          : %s\n", "192.168.0.100") // Simulator IP
	fmt.Printf("OS IP Address           : %s\n", "192.168.0.100")
	fmt.Printf("MAC Address             : %s\n", "00:00:00:00:00:00") // Mock
	fmt.Println("---Wireless Adapter---")
	fmt.Printf("DHCP Enabled            : false\n")
	fmt.Printf("DHCP Mode               : passive\n")
	fmt.Printf("Link Status             : down\n")
	fmt.Printf("AMT IP Address          : 0.0.0.0\n")
	fmt.Printf("OS IP Address           : 0.0.0.0\n")
	fmt.Printf("MAC Address             : 00:00:00:00:00:00\n")

	return nil
}
