package amt

import (
	"encoding/xml"
	"log"
)

type RemoteAccessPolicyAppliesToMPS struct {
	XMLName        xml.Name          `xml:"AMT_RemoteAccessPolicyAppliesToMPS"`
	PolicySet      EndpointReference `xml:"PolicySet"`
	ManagedElement EndpointReference `xml:"ManagedElement"`
	MpsType        int               `xml:"MpsType"`
	OrderOfAccess  int               `xml:"OrderOfAccess"`
}

type EndpointReference struct {
	Address             string              `xml:"a:Address"`
	ReferenceParameters ReferenceParameters `xml:"a:ReferenceParameters"`
}

type ReferenceParameters struct {
	ResourceURI string      `xml:"w:ResourceURI"`
	SelectorSet SelectorSet `xml:"w:SelectorSet"`
}

type SelectorSet struct {
	Selectors []Selector `xml:"w:Selector"`
}

type Selector struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:",chardata"`
}

func HandleRemoteAccessPolicyAppliesToMPS(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handleRemoteAccessPolicyAppliesToMPSEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handleRemoteAccessPolicyAppliesToMPSPull()
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put":
		return handleRemoteAccessPolicyAppliesToMPSPut(content)
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Create":
		return handleRemoteAccessPolicyAppliesToMPSCreate(content)
	case "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete":
		return handleRemoteAccessPolicyAppliesToMPSDelete(content)
	default:
		log.Printf("Unhandled RemoteAccessPolicyAppliesToMPS Action: %s", action)
		return nil
	}
}

func handleRemoteAccessPolicyAppliesToMPSEnumerate() []byte {
	// Standard Enumerate Response
	output := struct {
		XMLName            xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string   `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-rap-applies-mps",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleRemoteAccessPolicyAppliesToMPSPull() []byte {
	item1 := RemoteAccessPolicyAppliesToMPS{
		MpsType:       2, // Both
		OrderOfAccess: 0,
		PolicySet: EndpointReference{
			Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
			ReferenceParameters: ReferenceParameters{
				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
				SelectorSet: SelectorSet{
					Selectors: []Selector{
						{Name: "CreationClassName", Value: "AMT_RemoteAccessPolicyRule"},
						{Name: "PolicyRuleName", Value: "Periodic"},
						{Name: "SystemCreationClassName", Value: "CIM_ComputerSystem"},
						{Name: "SystemName", Value: "Intel(r) AMT"},
					},
				},
			},
		},
		ManagedElement: EndpointReference{
			Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
			ReferenceParameters: ReferenceParameters{
				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
				SelectorSet: SelectorSet{
					Selectors: []Selector{
						{Name: "CreationClassName", Value: "AMT_ManagementPresenceRemoteSAP"},
						{Name: "Name", Value: "Intel(r) AMT:Management Presence Server 0"},
						{Name: "SystemCreationClassName", Value: "CIM_ComputerSystem"},
						{Name: "SystemName", Value: "Intel(r) AMT"},
					},
				},
			},
		},
	}

	item2 := RemoteAccessPolicyAppliesToMPS{
		MpsType:       2, // Both
		OrderOfAccess: 0,
		PolicySet: EndpointReference{
			Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
			ReferenceParameters: ReferenceParameters{
				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule",
				SelectorSet: SelectorSet{
					Selectors: []Selector{
						{Name: "CreationClassName", Value: "AMT_RemoteAccessPolicyRule"},
						{Name: "PolicyRuleName", Value: "UserInitiated"},
						{Name: "SystemCreationClassName", Value: "CIM_ComputerSystem"},
						{Name: "SystemName", Value: "Intel(r) AMT"},
					},
				},
			},
		},
		ManagedElement: EndpointReference{
			Address: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
			ReferenceParameters: ReferenceParameters{
				ResourceURI: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP",
				SelectorSet: SelectorSet{
					Selectors: []Selector{
						{Name: "CreationClassName", Value: "AMT_ManagementPresenceRemoteSAP"},
						{Name: "Name", Value: "Intel(r) AMT:Management Presence Server 0"},
						{Name: "SystemCreationClassName", Value: "CIM_ComputerSystem"},
						{Name: "SystemName", Value: "Intel(r) AMT"},
					},
				},
			},
		},
	}

	output := struct {
		XMLName       xml.Name                         `xml:"PullResponse"`
		Items         []RemoteAccessPolicyAppliesToMPS `xml:"Items>AMT_RemoteAccessPolicyAppliesToMPS"`
		EndOfSequence string                           `xml:"EndOfSequence"`
	}{
		Items: []RemoteAccessPolicyAppliesToMPS{item1, item2},
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleRemoteAccessPolicyAppliesToMPSPut(content []byte) []byte {
	// Mock success
	return []byte(`<a:PutResponse xmlns:a="http://schemas.xmlsoap.org/ws/2004/09/transfer"/>`)
}

func handleRemoteAccessPolicyAppliesToMPSCreate(content []byte) []byte {
	// Mock success (Create returns the created object or a reference)
	return []byte(`<a:CreateResponse xmlns:a="http://schemas.xmlsoap.org/ws/2004/09/transfer"><ResourceCreated>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</ResourceCreated></a:CreateResponse>`)
}

func handleRemoteAccessPolicyAppliesToMPSDelete(content []byte) []byte {
	// Mock success
	return []byte(`<a:DeleteResponse xmlns:a="http://schemas.xmlsoap.org/ws/2004/09/transfer"/>`)
}
