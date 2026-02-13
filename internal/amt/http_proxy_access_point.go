package amt

import (
	"encoding/xml"
	"log"
)

type HTTPProxyAccessPoint struct {
	XMLName      xml.Name `xml:"IPS_HTTPProxyAccessPoint"`
	ElementName  string   `xml:"ElementName"`
	InstanceID   string   `xml:"InstanceID"`
	AccessInfo   string   `xml:"AccessInfo"`
	ProxyPort    int      `xml:"ProxyPort"`
}

type PullResponseHTTPProxyAccessPoint struct {
	XMLName xml.Name `xml:"PullResponse"`
	Items   []HTTPProxyAccessPoint `xml:"Items>IPS_HTTPProxyAccessPoint"`
}

func HandleHTTPProxyAccessPoint(action string, content []byte) []byte {
	switch action {
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate":
		return handleHTTPProxyAccessPointEnumerate()
	case "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull":
		return handleHTTPProxyAccessPointPull()
	default:
		log.Printf("Unhandled HTTPProxyAccessPoint Action: %s", action)
		return nil
	}
}

func handleHTTPProxyAccessPointEnumerate() []byte {
	output := struct {
		XMLName xml.Name `xml:"EnumerateResponse"`
		EnumerationContext string `xml:"EnumerationContext"`
	}{
		EnumerationContext: "ctx-http-proxy",
	}
	resp, _ := xml.Marshal(output)
	return resp
}

func handleHTTPProxyAccessPointPull() []byte {
	// Return empty list (no proxies)
	output := PullResponseHTTPProxyAccessPoint{
		Items: []HTTPProxyAccessPoint{},
	}

	resp, _ := xml.Marshal(output)
	return resp
}
