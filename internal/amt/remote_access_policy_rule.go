package amt

import (
	"encoding/xml"
	"log"
)

type RemoteAccessPolicyRule struct {
	XMLName      xml.Name `xml:"AMT_RemoteAccessPolicyRule"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
	PolicyRuleName string `xml:"PolicyRuleName"`
	Trigger      int      `xml:"Trigger"`
}

type PullResponseRemoteAccessPolicyRule struct {
	XMLName xml.Name `xml:"PullResponse"`
	Items   []RemoteAccessPolicyRule `xml:"Items>AMT_RemoteAccessPolicyRule"`
}

func HandleRemoteAccessPolicyRule(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handleRemoteAccessPolicyRuleEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handleRemoteAccessPolicyRulePull()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete":
		return handleRemoteAccessPolicyRuleDelete()
	default:
		log.Printf("Unhandled RemoteAccessPolicyRule Action: %s", action)
		return nil
	}
}

func handleRemoteAccessPolicyRuleEnumerate() []byte {
	output := struct {
		XMLName xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-remote-policy-rule",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleRemoteAccessPolicyRulePull() []byte {
	// Return a default rule for enumeration so RPS sees something to delete if needed, 
	// or return empty if we want to simulate a clean state.
	// Since RPS is sending Delete, it likely found something (or is blindly deleting).
	// Let's return one dummy item.
	item := RemoteAccessPolicyRule{
		ElementName:    "Intel(r) AMT Remote Access Policy Rule",
		InstanceID:     "Intel(r) AMT:RemoteAccessPolicyRule 0",
		PolicyRuleName: "Periodic",
		Trigger:        2,
	}

	output := PullResponseRemoteAccessPolicyRule{
		Items: []RemoteAccessPolicyRule{item},
	}

	resp, _ := xml.Marshal(output)
	return resp
}

func handleRemoteAccessPolicyRuleDelete() []byte {
	// Simple DeleteResponse
	output := struct {
		XMLName xml.Name `xml:"DeleteResponse"`
	}{}
	resp, _ := xml.Marshal(output)
	return resp
}
