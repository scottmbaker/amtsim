package wsman

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"amtsim/internal/amt"
	"amtsim/internal/logging"

	"github.com/google/uuid"
)

func Handle(r *http.Request) ([]byte, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			logging.Error("Failed to close request body: %v", err)
		}
	}()

	return Process(bodyBytes)
}

// Process handles the raw WSMAN envelope bytes and returns a response envelope bytes
func Process(bodyBytes []byte) ([]byte, error) {
	// Parse Envelope
	logging.Debug("Raw WSMAN Request: %s", string(bodyBytes))

	var reqEnvelope Envelope
	if err := xml.Unmarshal(bodyBytes, &reqEnvelope); err != nil {
		logging.Error("Failed to unmarshal body: %v", err)
		return nil, err
	}

	logging.Info("Received Action: %s", reqEnvelope.Header.Action)
	logging.Debug("Received ResourceURI: %s", reqEnvelope.Header.ResourceURI)

	// Route
	var respBody []byte
	switch reqEnvelope.Header.ResourceURI {
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystem":
		respBody = amt.HandleComputerSystem(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService":
		respBody = amt.HandlePowerService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_AssociatedPowerManagementService":
		respBody = amt.HandleAssociatedPowerManagementService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ServiceAvailableToElement":
		respBody = amt.HandleAssociatedPowerManagementService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings":
		respBody = amt.HandleGeneralSettings(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService":
		respBody = amt.HandleHostBasedSetupService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService":
		respBody = amt.HandleRemoteAccessService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings":
		respBody = amt.HandleEthernetPortSettings(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings":
		respBody = amt.HandleIEEE8021xSettings(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint":
		respBody = amt.HandleHTTPProxyAccessPoint(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP":
		respBody = amt.HandleManagementPresenceRemoteSAP(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData":
		respBody = amt.HandleTLSSettingData(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule":
		respBody = amt.HandleRemoteAccessPolicyRule(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate":
		respBody = amt.HandlePublicKeyCertificate(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicPrivateKeyPair":
		respBody = amt.HandlePublicPrivateKeyPair(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSCredentialContext":
		respBody = amt.HandleTLSCredentialContext(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData":
		respBody = amt.HandleEnvironmentDetectionSettingData(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService":
		respBody = amt.HandleOptInService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP":
		respBody = amt.HandleKVMRedirectionSAP(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService":
		respBody = amt.HandleRedirectionService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_PowerManagementService":
		respBody = amt.HandleIPSPowerManagementService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS":
		respBody = amt.HandleRemoteAccessPolicyAppliesToMPS(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootSourceSetting":
		respBody = amt.HandleBootSourceSetting(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService":
		respBody = amt.HandleBootService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService":
		respBody = amt.HandleSetupAndConfigurationService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData":
		respBody = amt.HandleBootSettingData(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities":
		respBody = amt.HandleBootCapabilities(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Processor":
		respBody = amt.HandleProcessor(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BIOSElement":
		respBody = amt.HandleBIOSElement(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Card":
		respBody = amt.HandleCard(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_Chassis":
		respBody = amt.HandleChassis(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ComputerSystemPackage":
		respBody = amt.HandleComputerSystemPackage(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_SoftwareIdentity":
		respBody = amt.HandleSoftwareIdentity(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	case "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService":
		respBody = amt.HandleUserInitiatedConnectionService(reqEnvelope.Header.Action, reqEnvelope.Body.Content)
	default:
		// Unknown resource, maybe return Fault
		logging.Error("Unknown ResourceURI: %s", reqEnvelope.Header.ResourceURI)
	}

	// Create Response using standard prefixes
	// MPS/RPS often require specific prefixes (s, a, w) to parse correctly.
	// Go's xml.Marshal is not good at enforcing prefixes, so we use a template.

	respHeaders := fmt.Sprintf(`<a:Header>
    <a:Action>%sResponse</a:Action>
    <a:RelatesTo>%s</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>%s</a:MessageID>
    <w:ResourceURI>%s</w:ResourceURI>
  </a:Header>`,
		reqEnvelope.Header.Action,
		reqEnvelope.Header.MessageID,
		uuid.New().String(),
		reqEnvelope.Header.ResourceURI,
	)

	respXml := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  %s
  <s:Body>%s</s:Body>
</s:Envelope>`, respHeaders, string(respBody))

	logging.Debug("WSMAN Response XML: %s", respXml)

	return []byte(respXml), nil
}
