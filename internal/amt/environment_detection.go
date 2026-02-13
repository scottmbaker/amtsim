package amt

import (
	"encoding/xml"
	"log"
)

type EnvironmentDetectionSettingData struct {
	XMLName          xml.Name `xml:"AMT_EnvironmentDetectionSettingData"`
	ElementName      string   `xml:"ElementName"`
	InstanceID       string   `xml:"InstanceID"`
	DetectionStrings []string `xml:"DetectionStrings"`
}

type GetResponseEnvironmentDetectionSettingData struct {
	XMLName xml.Name `xml:"AMT_EnvironmentDetectionSettingData"`
	// Embed fields directly as Get returns the object
	ElementName      string   `xml:"ElementName"`
	InstanceID       string   `xml:"InstanceID"`
	DetectionStrings []string `xml:"DetectionStrings"`
}

func HandleEnvironmentDetectionSettingData(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get":
		return handleEnvironmentDetectionSettingDataGet()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put":
		return handleEnvironmentDetectionSettingDataPut(content)
	default:
		log.Printf("Unhandled EnvironmentDetectionSettingData Action: %s", action)
		return nil
	}
}

func handleEnvironmentDetectionSettingDataGet() []byte {
	// Return empty detection strings
	output := GetResponseEnvironmentDetectionSettingData{
		ElementName:      "Intel(r) AMT Environment Detection Settings",
		InstanceID:       "Intel(r) AMT:EnvironmentDetectionSettingData 0",
		DetectionStrings: []string{},
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleEnvironmentDetectionSettingDataPut(content []byte) []byte {
	// Just return properties (mock success)
	// In a real implementation we would parse content and update state
	output := GetResponseEnvironmentDetectionSettingData{
		ElementName:      "Intel(r) AMT Environment Detection Settings",
		InstanceID:       "Intel(r) AMT:EnvironmentDetectionSettingData 0",
		DetectionStrings: []string{},
	}
	resp, _ := xml.Marshal(output)
	return resp
}
