package amt

import (
	"encoding/xml"
	"log"
)

const (
	ActionRequestPowerStateChange = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService/RequestPowerStateChange"
)

type RequestPowerStateChangeInput struct {
	XMLName xml.Name `xml:"RequestPowerStateChange_INPUT"`
	PowerState int `xml:"PowerState"`
}

type RequestPowerStateChangeOutput struct {
	XMLName xml.Name `xml:"RequestPowerStateChange_OUTPUT"`
	Xmlns   string   `xml:"xmlns,attr"`
	ReturnValue int `xml:"ReturnValue"`
}

func HandlePowerService(action string, body []byte) []byte {
	switch action {
	case ActionRequestPowerStateChange:
		var input RequestPowerStateChangeInput
		if err := xml.Unmarshal(body, &input); err != nil {
			log.Printf("Error unmarshalling power request: %v", err)
			return nil
		}
		log.Printf("[AMT] RequestPowerStateChange: %d", input.PowerState)
		
		// Map requested state to tracked state
		targetState := input.PowerState
		
		// 5 = Power Cycle (Off-Soft), 10 = Master Bus Reset -> Result is ON (2)
		if targetState == 5 || targetState == 10 {
			targetState = 2 
		}

		SetPowerState(targetState)
		log.Printf("[AMT] Power State Updated to: %d", GetPowerState())

		output := RequestPowerStateChangeOutput{
			Xmlns: "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService",
			ReturnValue: 0,
		}
		data, _ := xml.Marshal(output)
		return data
	case ActionGet:
		// Return service info?
		return nil
	default:
		return nil
	}
}

