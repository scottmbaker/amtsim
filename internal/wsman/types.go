package wsman

import (
	"encoding/xml"
)

// Namespaces
const (
	NS_SOAP_1_2 = "http://www.w3.org/2003/05/soap-envelope"
	NS_ADDRESSING = "http://schemas.xmlsoap.org/ws/2004/08/addressing"
	NS_WSMAN = "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
	NS_CIM_BINDING = "http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd"
	NS_CIM_SCHEMA = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header  Header   `xml:"http://www.w3.org/2003/05/soap-envelope Header"`
	Body    Body     `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
}

type Header struct {
	Action      string   `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Action"`
	MessageID   string   `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing MessageID"`
	To          string   `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing To"`
	ReplyTo     *ReplyTo `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing ReplyTo,omitempty"`
	RelatesTo   string   `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing RelatesTo,omitempty"`
	ResourceURI string   `xml:"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd ResourceURI"`
}

type ReplyTo struct {
	Address string `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Address"`
}

type Body struct {
	Content []byte `xml:",innerxml"`
}

// Fault Generation
func CreateFault(messageID string, code string, subCode string, reason string) *Envelope {
	return &Envelope{
		Header: Header{
			Action:    "http://schemas.dmtf.org/wbem/wsman/1/wsman/fault",
			MessageID: "uuid:generated-fault-" + messageID, // simplified
			To:        "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
			RelatesTo: messageID,
		},
		Body: Body{
			Content: []byte("<s:Fault>...</s:Fault>"), // TODO: Implement proper fault struct
		},
	}
}
