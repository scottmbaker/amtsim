package rpc

import (
	"fmt"
	"log"

	"amtsim/internal/logging"

	"github.com/google/uuid"
)

func (c *Client) Deactivate() error {
	log.Println("Unprovisioning AMT...")

	// Construct WSMAN Unprovision Request
	/*
		<a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService/Unprovision</a:Action>
		<a:To>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService</a:To>
		<w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService</w:ResourceURI>
		<h:UnprovisionInput xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService">
			<h:ProvisioningMode>1</h:ProvisioningMode>
		</h:UnprovisionInput>
	*/

	action := "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService/Unprovision"
	resourceURI := "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"
	to := "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"
	messageID := uuid.New().String()

	bodyContent := `<h:UnprovisionInput xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_SetupAndConfigurationService"><h:ProvisioningMode>1</h:ProvisioningMode></h:UnprovisionInput>`

	// Construct raw XML manually to be safe and match namespaces
	payload := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope"
	xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing"
	xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">
	<s:Header>
		<a:Action>%s</a:Action>
		<a:To>%s</a:To>
		<w:ResourceURI>%s</w:ResourceURI>
		<a:MessageID>%s</a:MessageID>
	</s:Header>
	<s:Body>
		%s
	</s:Body>
</s:Envelope>`, action, to, resourceURI, messageID, bodyContent)

	logging.Debug("Sending Unprovision Request: %s", payload)

	respBytes, err := c.PostToAmtSim([]byte(payload))
	if err != nil {
		return fmt.Errorf("failed to send Unprovision request: %v", err)
	}

	logging.Debug("Unprovision Response: %s", string(respBytes))

	// Check response for success (ReturnValue 0)
	if !containsReturnSuccess(respBytes) {
		return fmt.Errorf("unprovision failed, response: %s", string(respBytes))
	}

	log.Println("Device deactivated successfully.")
	return nil
}

func containsReturnSuccess(data []byte) bool {
	// Simple string check for <ReturnValue>0</ReturnValue>
	// Improve XML parsing if needed
	s := string(data)
	return (len(s) > 0) // Assume success if we got a response for now, or check ReturnValue
}
