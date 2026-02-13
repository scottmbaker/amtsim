package amt

import (
	"encoding/xml"
	"log"
)

type BootSourceSetting struct {
	XMLName        xml.Name `xml:"CIM_BootSourceSetting"`
	Xmlns          string   `xml:"xmlns,attr"`
	InstanceID     string   `xml:"InstanceID"`
	ElementName    string   `xml:"ElementName"`
	BIOSBootString string   `xml:"BIOSBootString"`
}

type BootSourceSettingEnumerateResponse struct {
	XMLName            xml.Name `xml:"wsen:EnumerateResponse"`
	EnumerationContext string   `xml:"wsen:EnumerationContext"`
}

type BootSourceSettingPullResponse struct {
	XMLName       xml.Name               `xml:"wsen:PullResponse"`
	Items         BootSourceSettingItems `xml:"wsen:Items"`
	EndOfSequence string                 `xml:"wsen:EndOfSequence"`
}

type BootSourceSettingItems struct {
	BootSourceSetting []BootSourceSetting
}

func HandleBootSourceSetting(action string, content []byte) []byte {
	log.Printf("HandleBootSourceSetting Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		resp := BootSourceSettingEnumerateResponse{
			EnumerationContext: "uuid:00000000-0000-0000-0000-000000000002",
		}
		data, _ := xml.Marshal(resp)
		return data
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Pull", "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		commonXmlns := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting"
		items := []BootSourceSetting{
			{
				Xmlns:          commonXmlns,
				InstanceID:     "Intel(r) AMT: Force OCR UEFI HTTPS Boot",
				ElementName:    "Intel(r) AMT: Boot Configuration",
				BIOSBootString: "HTTPS Boot",
			},
			{
				Xmlns:          commonXmlns,
				InstanceID:     "Intel(r) AMT: Force OCR UEFI Boot Option 1",
				ElementName:    "Intel(r) AMT: Boot Configuration",
				BIOSBootString: "WinRe",
			},
			{
				Xmlns:          commonXmlns,
				InstanceID:     "Intel(r) AMT: Force OCR UEFI Boot Option 2",
				ElementName:    "Intel(r) AMT: Boot Configuration",
				BIOSBootString: "PBA",
			},
		}
		resp := BootSourceSettingPullResponse{
			Items: BootSourceSettingItems{
				BootSourceSetting: items,
			},
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_BootSourceSetting Action: %s", action)
		return nil
	}
}
