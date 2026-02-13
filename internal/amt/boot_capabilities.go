package amt

import (
	"encoding/xml"
	"log"
)

type BootCapabilities struct {
	XMLName                 xml.Name `xml:"amt:AMT_BootCapabilities"`
	InstanceID              string   `xml:"amt:InstanceID"`
	ElementName             string   `xml:"amt:ElementName"`
	IDE                     bool     `xml:"amt:IDE"`
	HERD                    bool     `xml:"amt:HERD"`
	CD_DVD                  bool     `xml:"amt:CD_DVD"`
	PXE                     bool     `xml:"amt:PXE"`
	BIOSReflash             bool     `xml:"amt:BIOSReflash"`
	PowerUp                 bool     `xml:"amt:PowerUp"`
	PowerDown               bool     `xml:"amt:PowerDown"`
	PowerCycle              bool     `xml:"amt:PowerCycle"`
	PowerDwnRbt             bool     `xml:"amt:PowerDwnRbt"`
	ForceUEFIHTTPSBoot      bool     `xml:"amt:ForceUEFIHTTPSBoot"`
	ForceWinREBoot          bool     `xml:"amt:ForceWinREBoot"`
	ForceUEFIPBABoot        bool     `xml:"amt:ForceUEFIPBABoot"`
}

func HandleBootCapabilities(action string, content []byte) []byte {
	log.Printf("HandleBootCapabilities Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		resp := BootCapabilities{
			InstanceID:         "Intel(r) AMT: Boot Capabilities 0",
			ElementName:        "Intel(r) AMT: Boot Capabilities",
			IDE:                true,
			HERD:               true,
			CD_DVD:             true,
			PXE:                true,
			BIOSReflash:        true,
			PowerUp:            true,
			PowerDown:          true,
			PowerCycle:         true,
			PowerDwnRbt:        true,
			ForceUEFIHTTPSBoot: true,
			ForceWinREBoot:     true,
			ForceUEFIPBABoot:   true,
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled AMT_BootCapabilities Action: %s", action)
		return nil
	}
}
