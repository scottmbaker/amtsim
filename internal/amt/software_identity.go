package amt

import (
	"encoding/xml"
	"log"
)

type SoftwareIdentity struct {
	XMLName           xml.Name `xml:"CIM_SoftwareIdentity"`
	Xmlns             string   `xml:"xmlns,attr"`
	InstanceID        string   `xml:"InstanceID"`
	ElementName       string   `xml:"ElementName"`
	VersionString     string   `xml:"VersionString"`
	IdentityInfoType  string   `xml:"IdentityInfoType"` // "Intel(r) AMT:MEBx Version" etc
	IdentityInfoValue string   `xml:"IdentityInfoValue"`
}

type SoftwareIdentityEnumerateResponse struct {
	XMLName            xml.Name `xml:"wsen:EnumerateResponse"`
	EnumerationContext string   `xml:"wsen:EnumerationContext"`
}

type SoftwareIdentityPullResponse struct {
	XMLName       xml.Name               `xml:"wsen:PullResponse"`
	Items         SoftwareIdentityItems  `xml:"wsen:Items"`
	EndOfSequence string                 `xml:"wsen:EndOfSequence"`
}

type SoftwareIdentityItems struct {
	SoftwareIdentity []SoftwareIdentity
}

func HandleSoftwareIdentity(action string, content []byte) []byte {
	log.Printf("HandleSoftwareIdentity Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		resp := SoftwareIdentityEnumerateResponse{
			EnumerationContext: "uuid:00000000-0000-0000-0000-000000000004",
		}
		data, _ := xml.Marshal(resp)
		return data
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Pull", "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		// MPS expects specific indices: 4 (SKU), 6 (Build/MEBx), 10 (Version/Firmware)
		// We create 12 items to cover index 10.
		items := make([]SoftwareIdentity, 12)
		commonXmlns := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity"

		for i := range items {
			items[i] = SoftwareIdentity{
				Xmlns:         commonXmlns,
				InstanceID:    "Intel(r) AMT: Software Identity " + string(rune('A'+i)), // A, B, C...
				ElementName:   "Intel(r) AMT: Software Identity",
				VersionString: "0.0.0",
                // Ensure IdentityInfoType is present to avoid empty tags if omitempty is not used (it's not)
                IdentityInfoType: "Unknown",
                IdentityInfoValue: "Unknown",
			}
		}

		// Index 4: SKU
		items[4].InstanceID = "Intel(r) AMT: PCH SKU"
		items[4].ElementName = "Intel(r) AMT: PCH SKU"
		items[4].VersionString = "16392"
		items[4].IdentityInfoType = "PCHSKU"
        items[4].IdentityInfoValue = "16392"

		// Index 6: MEBx Version
		items[6].InstanceID = "Intel(r) AMT: MEBx Version"
		items[6].ElementName = "Intel(r) AMT: MEBx Version"
		items[6].VersionString = "14.0.0.0001"
        items[6].IdentityInfoType = "MEBx"
        items[6].IdentityInfoValue = "14.0.0.0001"

		// Index 10: Firmware Version
		items[10].InstanceID = "Intel(r) AMT: Firmware Version"
		items[10].ElementName = "Intel(r) AMT: Firmware Version"
		items[10].VersionString = "14.0.45"
		items[10].IdentityInfoType = "VendorID"
		items[10].IdentityInfoValue = "Intel"

		resp := SoftwareIdentityPullResponse{
			Items: SoftwareIdentityItems{
				SoftwareIdentity: items,
			},
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_SoftwareIdentity Action: %s", action)
		return nil
	}
}
