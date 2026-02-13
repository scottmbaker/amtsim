package amt

import (
	"encoding/xml"
	"log"
)

const (
	ActionEnumerate = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	ActionPull      = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
)

type AssociatedPowerManagementService struct {
	XMLName    xml.Name `xml:"CIM_AssociatedPowerManagementService"`
	PowerState int      `xml:"PowerState"` // 0=Unknown, 2=On, 3-9=Sleep/Off...
}

type EnumerateResponse struct {
	XMLName            xml.Name `xml:"wsen:EnumerateResponse"`
	EnumerationContext string   `xml:"wsen:EnumerationContext"`
}

type PullResponse struct {
	XMLName       xml.Name `xml:"wsen:PullResponse"`
	Items         Items    `xml:"wsen:Items"`
	EndOfSequence string   `xml:"wsen:EndOfSequence"`
}

type Items struct {
	AssociatedPowerManagementService AssociatedPowerManagementService
}

func HandleAssociatedPowerManagementService(action string, body []byte) []byte {
	log.Printf("HandleAssociatedPowerManagementService Action: %s", action)
	switch action {
	case ActionGet:
		// Not typically used for this class by MPS, but good to have
		return nil
	case ActionEnumerate:
		resp := EnumerateResponse{
			EnumerationContext: "uuid:00000000-0000-0000-0000-000000000001",
		}
		data, _ := xml.Marshal(resp)
		return data
	case ActionPull:
		// Return current tracked state
		currentState := GetPowerState()
		log.Printf("[AMT] CIM_AssociatedPowerManagementService: PowerState=%d", currentState)
		item := AssociatedPowerManagementService{
			PowerState: currentState,
		}
		resp := PullResponse{
			Items: Items{
				AssociatedPowerManagementService: item,
			},
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_AssociatedPowerManagementService Action: %s", action)
		return nil
	}
}
