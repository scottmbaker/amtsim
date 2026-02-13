package amt

import (
	"encoding/xml"
	"log"
)

type Processor struct {
	XMLName           xml.Name `xml:"cim:CIM_Processor"`
	InstanceID        string   `xml:"cim:InstanceID"`
	ElementName       string   `xml:"cim:ElementName"`
	DeviceID          string   `xml:"cim:DeviceID"`
	MaxClockSpeed     int      `xml:"cim:MaxClockSpeed"`
	Names             string   `xml:"cim:Name"`
	Family            int      `xml:"cim:Family"` // 205 = Core i7
	NumberOfEnabledCores int    `xml:"cim:NumberOfEnabledCores"`
}

type ProcessorEnumerateResponse struct {
	XMLName            xml.Name `xml:"wsen:EnumerateResponse"`
	EnumerationContext string   `xml:"wsen:EnumerationContext"`
}

type ProcessorPullResponse struct {
	XMLName       xml.Name         `xml:"wsen:PullResponse"`
	Items         ProcessorItems   `xml:"wsen:Items"`
	EndOfSequence string           `xml:"wsen:EndOfSequence"`
}

type ProcessorItems struct {
	Processor []Processor
}

func HandleProcessor(action string, content []byte) []byte {
	log.Printf("HandleProcessor Action: %s", action)
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		resp := ProcessorEnumerateResponse{
			EnumerationContext: "uuid:00000000-0000-0000-0000-000000000003",
		}
		data, _ := xml.Marshal(resp)
		return data
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Pull", "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		items := []Processor{
			{
				InstanceID:           "Intel(r) AMT: Processor 0",
				ElementName:          "Intel(r) Core(TM) i7-8665U CPU @ 1.90GHz",
				DeviceID:             "CPU0",
				MaxClockSpeed:        1900,
				Names:                "Intel(r) Core(TM) i7-8665U CPU @ 1.90GHz",
				Family:               205,
				NumberOfEnabledCores: 4,
			},
		}
		resp := ProcessorPullResponse{
			Items: ProcessorItems{
				Processor: items,
			},
		}
		data, _ := xml.Marshal(resp)
		return data
	default:
		log.Printf("Unhandled CIM_Processor Action: %s", action)
		return nil
	}
}
