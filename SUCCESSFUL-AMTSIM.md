This is what output for a successful activation looks like in amtsim:

```

2026/02/12 15:35:25 [INFO]  Loaded saved CIRA credentials for user: admin
2026/02/12 15:35:25 [INFO]  Starting AMT Simulator on :16992...
2026/02/12 15:35:25 [INFO]  Initializing CIRA client for 198.0.0.61:4433
2026/02/12 15:35:25 [INFO]  Main: Wait for terminate...
2026/02/12 15:35:25 Listening on :16992
2026/02/12 15:35:25 [INFO]  Starting CIRA client connecting to 198.0.0.61:4433
2026/02/12 15:35:25 [INFO]  Connecting to MPS at 198.0.0.61:4433...
2026/02/12 15:35:25 [INFO]  Attempting CIRA Auth with User: admin, Pass: 0Y7cFNbnbL%Bz6@d
2026/02/12 15:35:25 [INFO]  CIRA Authenticated Successfully.
2026/02/12 15:35:41 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

2CE
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings</w:ResourceURI><a:MessageID>0</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body></Body></Envelope>
0

2026/02/12 15:35:41 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:35:41 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings
2026/02/12 15:35:41 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</a:Action>
    <a:RelatesTo>0</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>eec764e2-e25a-433c-b2e7-b8c350a4f36b</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings</w:ResourceURI>
  </a:Header>
  <s:Body><AMT_GeneralSettings xmlns="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings"><AMTNetworkEnabled>1</AMTNetworkEnabled><DDNSPeriodicUpdateInterval>1440</DDNSPeriodicUpdateInterval><DDNSTTL>900</DDNSTTL><DDNSUpdateByDHCPServerEnabled>true</DDNSUpdateByDHCPServerEnabled><DDNSUpdateEnabled>false</DDNSUpdateEnabled><DHCPv6ConfigurationTimeout>0</DHCPv6ConfigurationTimeout><DigestRealm>Digest:A3829B3827DE4D33D4449B366831FD01</DigestRealm><DomainName></DomainName><ElementName>Intel(r) AMT: General Settings</ElementName><HostName></HostName><HostOSFQDN>AMTSIM-HOST</HostOSFQDN><IdleWakeTimeout>1</IdleWakeTimeout><InstanceID>Intel(r) AMT: General Settings</InstanceID><NetworkInterfaceEnabled>true</NetworkInterfaceEnabled><PingResponseEnabled>true</PingResponseEnabled><PowerSource>0</PowerSource><PreferredAddressFamily>0</PreferredAddressFamily><PresenceNotificationInterval>0</PresenceNotificationInterval><PrivacyLevel>0</PrivacyLevel><RmcpPingResponseEnabled>true</RmcpPingResponseEnabled><SharedFQDN>true</SharedFQDN><ThunderboltDockEnabled>0</ThunderboltDockEnabled><WsmanOnlyMode>false</WsmanOnlyMode></AMT_GeneralSettings></s:Body>
</s:Envelope>
2026/02/12 15:35:41 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

3E9
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/Setup</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService</w:ResourceURI><a:MessageID>0</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:Setup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>2</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>74345087234a9ac5ba001c2fe6059e2d</h:NetworkAdminPassword></h:Setup_INPUT></Body></Envelope>
0

2026/02/12 15:35:41 [INFO]  Received Action: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/Setup
2026/02/12 15:35:41 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService
2026/02/12 15:35:41 setupHostBasedSetupService Content: <h:Setup_INPUT xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService"><h:NetAdminPassEncryptionType>2</h:NetAdminPassEncryptionType><h:NetworkAdminPassword>74345087234a9ac5ba001c2fe6059e2d</h:NetworkAdminPassword></h:Setup_INPUT>
2026/02/12 15:35:41 Received NetworkAdminPassword (len: 32)
2026/02/12 15:35:41 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/SetupResponse</a:Action>
    <a:RelatesTo>0</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>b187781a-bd34-4d7c-982c-a3142c0eae89</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService</w:ResourceURI>
  </a:Header>
  <s:Body><Setup_OUTPUT><ReturnValue>0</ReturnValue></Setup_OUTPUT></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

323
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI><a:MessageID>1</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</a:Action>
    <a:RelatesTo>1</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>5841d9c1-09ba-4681-a2e4-dbb9fa941c59</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI>
  </a:Header>
  <s:Body><EnumerateResponse><EnumerationContext>ctx-ethernet-settings</EnumerationContext></EnumerateResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

39E
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI><a:MessageID>2</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>ctx-ethernet-settings</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</a:Action>
    <a:RelatesTo>2</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>8c149cd1-51be-4782-8239-c492aac146c8</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI>
  </a:Header>
  <s:Body><PullResponse><Items><AMT_EthernetPortSettings><ElementName>Intel(r) AMT Ethernet Port Settings</ElementName><InstanceID>Intel(r) AMT:EthernetPortSettings 0</InstanceID><MACAddress>00-00-00-00-00-00</MACAddress><LinkIsUp>true</LinkIsUp><DHCPEnabled>true</DHCPEnabled></AMT_EthernetPortSettings></Items></PullResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

2D0
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings</w:ResourceURI><a:MessageID>1</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</a:Action>
    <a:RelatesTo>1</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>d1243bd3-6e25-4cca-9e01-9d2b98f5a3e9</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings</w:ResourceURI>
  </a:Header>
  <s:Body><IPS_IEEE8021xSettings><ElementName>Intel(r) AMT 802.1x Settings</ElementName><InstanceID>Intel(r) AMT:IEEE8021xSettings 0</InstanceID><Enabled>false</Enabled></IPS_IEEE8021xSettings></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

323
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint</w:ResourceURI><a:MessageID>2</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</a:Action>
    <a:RelatesTo>2</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>92949c18-cfea-49a9-9663-ddf22349735d</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint</w:ResourceURI>
  </a:Header>
  <s:Body><EnumerateResponse><EnumerationContext>ctx-http-proxy</EnumerationContext></EnumerateResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

397
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint</w:ResourceURI><a:MessageID>3</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>ctx-http-proxy</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</a:Action>
    <a:RelatesTo>3</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>81d305eb-6c90-4069-8676-68412c07ed1c</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint</w:ResourceURI>
  </a:Header>
  <s:Body><PullResponse><Items></Items></PullResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

334
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI><a:MessageID>3</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name="PolicyRuleName">User Initiated</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/DeleteResponse</a:Action>
    <a:RelatesTo>3</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>82d97a41-69b6-4997-b446-c70f4f7ce010</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI>
  </a:Header>
  <s:Body><DeleteResponse></DeleteResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

32B
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI><a:MessageID>4</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name="PolicyRuleName">Alert</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/DeleteResponse</a:Action>
    <a:RelatesTo>4</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>f812390d-1a3b-4696-8681-7546adff6fbb</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI>
  </a:Header>
  <s:Body><DeleteResponse></DeleteResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

32E
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI><a:MessageID>5</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name="PolicyRuleName">Periodic</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/DeleteResponse</a:Action>
    <a:RelatesTo>5</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>b8ddc1e7-35a3-424e-91c4-cf80574f0867</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI>
  </a:Header>
  <s:Body><DeleteResponse></DeleteResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

32A
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><a:MessageID>6</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</a:Action>
    <a:RelatesTo>6</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>43d77bc8-845c-4854-aa33-11fada6866f9</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI>
  </a:Header>
  <s:Body><EnumerateResponse><EnumerationContext>ctx-remote-sap</EnumerationContext></EnumerateResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

39E
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><a:MessageID>7</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>ctx-remote-sap</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</a:Action>
    <a:RelatesTo>7</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>61fce626-a3aa-407c-98e4-3f190baa138a</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI>
  </a:Header>
  <s:Body><PullResponse><Items><AMT_ManagementPresenceRemoteSAP><ElementName>Intel(r) AMT:Management Presence Server 0</ElementName><InstanceID>Intel(r) AMT:Management Presence Server 0</InstanceID><Name>Intel(r) AMT:Management Presence Server 0</Name><AccessInfo>192.168.0.1</AccessInfo></AMT_ManagementPresenceRemoteSAP></Items></PullResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

34A
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><a:MessageID>8</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name="Name">Intel(r) AMT:Management Presence Server 0</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/DeleteResponse</a:Action>
    <a:RelatesTo>8</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>bb4da9a0-0399-4686-92ea-7d8f78748235</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI>
  </a:Header>
  <s:Body><DeleteResponse></DeleteResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

31D
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData</w:ResourceURI><a:MessageID>9</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</a:Action>
    <a:RelatesTo>9</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>88b43a16-9aa7-4ec6-95c2-8317e608fc02</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData</w:ResourceURI>
  </a:Header>
  <s:Body><EnumerateResponse><EnumerationContext>ctx-tls-settings</EnumerationContext></EnumerateResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

394
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData</w:ResourceURI><a:MessageID>10</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>ctx-tls-settings</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</a:Action>
    <a:RelatesTo>10</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>e466e9a1-c091-4843-9257-f1d7d7a799f2</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData</w:ResourceURI>
  </a:Header>
  <s:Body><PullResponse><Items><AMT_TLSSettingData><ElementName>Intel(r) AMT LMS TLS Settings</ElementName><InstanceID>Intel(r) AMT:TLSSettingData 0</InstanceID><Enabled>false</Enabled><MutualAuthentication>false</MutualAuthentication></AMT_TLSSettingData><AMT_TLSSettingData><ElementName>Intel(r) AMT 802.3 TLS Settings</ElementName><InstanceID>Intel(r) AMT:TLSSettingData 1</InstanceID><Enabled>false</Enabled><MutualAuthentication>false</MutualAuthentication></AMT_TLSSettingData></Items></PullResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

324
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><a:MessageID>11</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</a:Action>
    <a:RelatesTo>11</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>b2a89b42-1e14-4673-91aa-e287ce7699b5</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI>
  </a:Header>
  <s:Body><EnumerateResponse><EnumerationContext>ctx-public-key-cert</EnumerationContext></EnumerateResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

39D
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI><a:MessageID>12</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>ctx-public-key-cert</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</a:Action>
    <a:RelatesTo>12</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>0c6c578e-5ba6-4b60-b125-34d74542e5be</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate</w:ResourceURI>
  </a:Header>
  <s:Body><PullResponse><Items></Items></PullResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

2DF
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData</w:ResourceURI><a:MessageID>13</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</a:Action>
    <a:RelatesTo>13</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>a09ba3ef-1ff2-49e0-97cb-49d6fc742b70</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData</w:ResourceURI>
  </a:Header>
  <s:Body><AMT_EnvironmentDetectionSettingData><ElementName>Intel(r) AMT Environment Detection Settings</ElementName><InstanceID>Intel(r) AMT:EnvironmentDetectionSettingData 0</InstanceID></AMT_EnvironmentDetectionSettingData></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

324
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI><a:MessageID>14</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</a:Action>
    <a:RelatesTo>14</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>062e4c92-ce3e-4c93-a750-aa8e985a42c6</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI>
  </a:Header>
  <s:Body><EnumerateResponse><EnumerationContext>ctx-ethernet-settings</EnumerationContext></EnumerateResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

39F
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI><a:MessageID>15</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>ctx-ethernet-settings</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</a:Action>
    <a:RelatesTo>15</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>b6e400c6-0650-4d5b-9acb-bbdaceed78d8</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI>
  </a:Header>
  <s:Body><PullResponse><Items><AMT_EthernetPortSettings><ElementName>Intel(r) AMT Ethernet Port Settings</ElementName><InstanceID>Intel(r) AMT:EthernetPortSettings 0</InstanceID><MACAddress>00-00-00-00-00-00</MACAddress><LinkIsUp>true</LinkIsUp><DHCPEnabled>true</DHCPEnabled></AMT_EthernetPortSettings></Items></PullResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

509
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI><a:MessageID>16</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name="InstanceID">Intel(r) AMT:EthernetPortSettings 0</w:Selector></w:SelectorSet></Header><Body><h:AMT_EthernetPortSettings xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings"><h:ElementName>Intel(r) AMT Ethernet Port Settings</h:ElementName><h:InstanceID>Intel(r) AMT:EthernetPortSettings 0</h:InstanceID><h:MACAddress>00-00-00-00-00-00</h:MACAddress><h:LinkIsUp>true</h:LinkIsUp><h:DHCPEnabled>true</h:DHCPEnabled><h:IpSyncEnabled>true</h:IpSyncEnabled><h:SharedStaticIp>false</h:SharedStaticIp></h:AMT_EthernetPortSettings></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/PutResponse</a:Action>
    <a:RelatesTo>16</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>fb7dc5de-1b19-44f8-b187-27db9d0947f3</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings</w:ResourceURI>
  </a:Header>
  <s:Body><AMT_EthernetPortSettings><ElementName>Intel(r) AMT Ethernet Port Settings</ElementName><InstanceID>Intel(r) AMT:EthernetPortSettings 0</InstanceID><MACAddress>00-00-00-00-00-00</MACAddress><LinkIsUp>true</LinkIsUp><DHCPEnabled>true</DHCPEnabled></AMT_EthernetPortSettings></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

2D2
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService</w:ResourceURI><a:MessageID>17</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</a:Action>
    <a:RelatesTo>17</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>529e1eba-fe62-4753-8b3b-56248cab5138</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService</w:ResourceURI>
  </a:Header>
  <s:Body><AMT_RedirectionService><ElementName>Intel(r) AMT Redirection Service</ElementName><InstanceID>Intel(r) AMT:RedirectionService 0</InstanceID><ListenerEnabled>true</ListenerEnabled></AMT_RedirectionService></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

2CB
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService</w:ResourceURI><a:MessageID>4</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</a:Action>
    <a:RelatesTo>4</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>ba92e3d7-8499-4b11-b125-1a8c455d9ed2</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService</w:ResourceURI>
  </a:Header>
  <s:Body><IPS_OptInService><ElementName>Intel(r) AMT OptIn Service</ElementName><InstanceID>Intel(r) AMT:OptInService 0</InstanceID><OptInRequired>0</OptInRequired><OptInState>0</OptInState><CanModifyOptInPolicy>true</CanModifyOptInPolicy></IPS_OptInService></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

2D7
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP</w:ResourceURI><a:MessageID>0</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</a:Action>
    <a:RelatesTo>0</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>a680370b-0db9-44da-8455-c767fb9cf763</a:MessageID>
    <w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP</w:ResourceURI>
  </a:Header>
  <s:Body><CIM_KVMRedirectionSAP xmlns="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP"><ElementName>Intel(r) AMT KVM Redirection SAP</ElementName><InstanceID>Intel(r) AMT:KVMRedirectionSAP 0</InstanceID><EnabledState>2</EnabledState></CIM_KVMRedirectionSAP></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

3A3
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService/RequestStateChange</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService</w:ResourceURI><a:MessageID>18</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:RequestStateChange_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService/RequestStateChange
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService
2026/02/12 15:35:53 Unhandled RedirectionService Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService/RequestStateChange
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService/RequestStateChangeResponse</a:Action>
    <a:RelatesTo>18</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>167bfe5f-58cf-4afd-8507-2a7ec7dade9c</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService</w:ResourceURI>
  </a:Header>
  <s:Body></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

3B0
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP/RequestStateChange</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP</w:ResourceURI><a:MessageID>1</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:RequestStateChange_INPUT xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP"><h:RequestedState>2</h:RequestedState></h:RequestStateChange_INPUT></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP/RequestStateChange
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP
2026/02/12 15:35:53 Unhandled KVMRedirectionSAP Action: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP/RequestStateChange
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP/RequestStateChangeResponse</a:Action>
    <a:RelatesTo>1</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>4d1f93cb-13bb-4a91-aaf4-ef6a1b343458</a:MessageID>
    <w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP</w:ResourceURI>
  </a:Header>
  <s:Body></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

421
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService</w:ResourceURI><a:MessageID>19</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:AMT_RedirectionService xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService"><h:ElementName>Intel(r) AMT Redirection Service</h:ElementName><h:InstanceID>Intel(r) AMT:RedirectionService 0</h:InstanceID><h:ListenerEnabled>true</h:ListenerEnabled><h:EnabledState>32771</h:EnabledState></h:AMT_RedirectionService></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/PutResponse</a:Action>
    <a:RelatesTo>19</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>848980f4-1fd4-44c5-952f-26f700d73192</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService</w:ResourceURI>
  </a:Header>
  <s:Body><AMT_RedirectionService><ElementName>Intel(r) AMT Redirection Service</ElementName><InstanceID>Intel(r) AMT:RedirectionService 0</InstanceID><ListenerEnabled>true</ListenerEnabled></AMT_RedirectionService></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

42B
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService</w:ResourceURI><a:MessageID>5</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:IPS_OptInService xmlns:h="http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService"><h:ElementName>Intel(r) AMT OptIn Service</h:ElementName><h:InstanceID>Intel(r) AMT:OptInService 0</h:InstanceID><h:OptInRequired>4294967295</h:OptInRequired><h:OptInState>0</h:OptInState><h:CanModifyOptInPolicy>true</h:CanModifyOptInPolicy></h:IPS_OptInService></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/PutResponse</a:Action>
    <a:RelatesTo>5</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>b977a0b1-4d22-4711-b95f-4db092083c63</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService</w:ResourceURI>
  </a:Header>
  <s:Body><IPS_OptInService><ElementName>Intel(r) AMT OptIn Service</ElementName><InstanceID>Intel(r) AMT:OptInService 0</InstanceID><OptInRequired>0</OptInRequired><OptInState>0</OptInState><CanModifyOptInPolicy>true</CanModifyOptInPolicy></IPS_OptInService></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

979
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService</w:ResourceURI><a:MessageID>20</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:AddTrustedRootCertificate_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService"><h:CertificateBlob>MIIEOzCCAqOgAwIBAgIDCTZYMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtN2VmZjRkMRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTI1MDIwNzAwMzk0NFoYDzIwNTYwMjA3MDAzOTQ0WjA9MRcwFQYDVQQDEw5NUFNSb290LTdlZmY0ZDEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBAKc9vA3M3N5QtT/81INdeJd2g2zwbKBvRFV1ZMNqssEObwJIFV56vaZEHCaCIvmR60QtX71YIiDASu1/3YY7HemriSl38rcnpKDZHDl6DDUYCj9G5bz15zNyBDbWaen9F0u96UAGiwBd2aLI1U2ducL2z938NCrDrppJlbH4nvFoWtiK4UIESBpsMqxk329z+2MX5hnKr3j6gd6gUKeSb35ZL/8SARQHmmQhX3d6KWlwhbxnIomzxG9ZeguEeNvXcV1WcLh4cVGSol58htkJe1XaNMSFcMMKV5f7TXk+QW0gK8n/3ZvUiYS0EkaYrSqitSLT5N8FcDXxpUPyWbK2SNc5K2Sios0Vj26Yk+yXOI5vCP4w/2h7QrA3+8xTNZv00zLUDfeUHKDFkVGTKjPDpEBjrmu6zxZZpTePru3kLQxMMpYi5ewJO/0gG0wFDDU6CUOfIsU2+XKCatqFadVC5B4hLPCE+5psPIqL2qJRltYfqCAulN26LX7UjdKBw1b4CwIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUfv9NXQe52+dcfNwBVlkFshYQBUwwDQYJKoZIhvcNAQEMBQADggGBAIfzsoqkZBRTM/03NCYroyQDfUvbs9WKExnj1FEzZ1JYCNuzEc5wNLvzpVdX4Sy+8y80rC6hw1+xL68KzodigkIF5FZendJPBXpKUnuRuRupmmTpOwQp/csg07K5l/ui9WPnfTWGRvCi+pnzJF01r0PIz6a8zY29ey+NBmbw5NJnuGdupF1w0mJRmd9yYTvay5EwNEb50tFozO+qaMBg6RJG9cU+pVWeuzewngQ32x8qxAyS/5aNg1IeBPZdWy7lsq5q9tHzp4RPi5FsYWQUdkfqYY3Vs2Ha80CXzNF8zJ/eX/pMUcG5AhASCKgyNi2PbDJnVH82dvgeND1QBatqk/IEr3h+qiqxafo7+TF1skKJu73RaHhvHCOHFH3hyzCLNX5KeM6re2NYFxvYoqXw7my4IR8ca6YIsPC0Ry0nGO8LgGwjtlSFhatj91mj14cP4jMLcWFLgVQ6cGwrm+FgB4gDpS53SEEkZIBeMhiinzilVwelP52sy/rkji1kd9Kc/w==</h:CertificateBlob></h:AddTrustedRootCertificate_INPUT></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService
2026/02/12 15:35:53 [ERROR] Unknown ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificateResponse</a:Action>
    <a:RelatesTo>20</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>15cf8c26-61fd-42d0-b150-edff04929375</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService</w:ResourceURI>
  </a:Header>
  <s:Body></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

43D
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddMpServer</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService</w:ResourceURI><a:MessageID>21</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:AddMpServer_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService"><h:AccessInfo>198.0.0.61</h:AccessInfo><h:InfoFormat>3</h:InfoFormat><h:Port>4433</h:Port><h:AuthMethod>2</h:AuthMethod><h:Username>admin</h:Username><h:Password>%BqYF!B#QccJo0OM</h:Password><h:CN>198.0.0.61</h:CN></h:AddMpServer_INPUT></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddMpServer
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService
2026/02/12 15:35:53 [INFO]  CIRA Client credentials updated. User: admin
2026/02/12 15:35:53 [INFO]  Credentials saved to cira_creds.json
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddMpServerResponse</a:Action>
    <a:RelatesTo>21</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>7458aad4-7445-4d3f-8fc4-83af146290c8</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService</w:ResourceURI>
  </a:Header>
  <s:Body><AddMpServer_OUTPUT><ReturnValue>0</ReturnValue></AddMpServer_OUTPUT></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

32B
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><a:MessageID>22</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</a:Action>
    <a:RelatesTo>22</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>3f3ff0df-6314-43ae-9119-4a3178db1101</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI>
  </a:Header>
  <s:Body><EnumerateResponse><EnumerationContext>ctx-remote-sap</EnumerationContext></EnumerateResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

39F
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><a:MessageID>23</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>ctx-remote-sap</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</a:Action>
    <a:RelatesTo>23</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>bfb9ad2d-22f1-4d7f-94d2-aaf5c1e22147</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI>
  </a:Header>
  <s:Body><PullResponse><Items><AMT_ManagementPresenceRemoteSAP><ElementName>Intel(r) AMT:Management Presence Server 0</ElementName><InstanceID>Intel(r) AMT:Management Presence Server 0</InstanceID><Name>Intel(r) AMT:Management Presence Server 0</Name><AccessInfo>192.168.0.1</AccessInfo></AMT_ManagementPresenceRemoteSAP></Items></PullResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

63C
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRule</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService</w:ResourceURI><a:MessageID>24</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:AddRemoteAccessPolicyRule_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService"><h:Trigger>2</h:Trigger><h:TunnelLifeTime>0</h:TunnelLifeTime><h:ExtendedData>AAAAAAAAABk=</h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="Name">Intel(r) AMT:Management Presence Server 0</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRule
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRuleResponse</a:Action>
    <a:RelatesTo>24</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>c1d875a3-daae-45e3-b5c1-4a3cba856c2f</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService</w:ResourceURI>
  </a:Header>
  <s:Body><AddRemoteAccessPolicyRule_OUTPUT><ReturnValue>0</ReturnValue></AddRemoteAccessPolicyRule_OUTPUT></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

632
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRule</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService</w:ResourceURI><a:MessageID>25</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:AddRemoteAccessPolicyRule_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService"><h:Trigger>0</h:Trigger><h:TunnelLifeTime>300</h:TunnelLifeTime><h:ExtendedData></h:ExtendedData><h:MpServer><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="Name">Intel(r) AMT:Management Presence Server 0</Selector></SelectorSet></ReferenceParameters></h:MpServer></h:AddRemoteAccessPolicyRule_INPUT></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRule
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRuleResponse</a:Action>
    <a:RelatesTo>25</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>a3cf8e01-27ee-41e8-b01e-70e2331b5894</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService</w:ResourceURI>
  </a:Header>
  <s:Body><AddRemoteAccessPolicyRule_OUTPUT><ReturnValue>0</ReturnValue></AddRemoteAccessPolicyRule_OUTPUT></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

32E
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</w:ResourceURI><a:MessageID>26</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/EnumerateResponse</a:Action>
    <a:RelatesTo>26</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>2d4cb68b-be0f-4797-b0bb-a815e69f8211</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</w:ResourceURI>
  </a:Header>
  <s:Body><EnumerateResponse><EnumerationContext>ctx-rap-applies-mps</EnumerationContext></EnumerateResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

3A7
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</w:ResourceURI><a:MessageID>27</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>ctx-rap-applies-mps</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/PullResponse</a:Action>
    <a:RelatesTo>27</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>3e1ba725-dc7e-45b2-a27e-ccb2f7c33d96</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</w:ResourceURI>
  </a:Header>
  <s:Body><PullResponse><Items><AMT_RemoteAccessPolicyAppliesToMPS><PolicySet><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</w:Selector><w:Selector Name="PolicyRuleName">Periodic</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet></a:ReferenceParameters></PolicySet><ManagedElement><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</w:Selector><w:Selector Name="Name">Intel(r) AMT:Management Presence Server 0</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet></a:ReferenceParameters></ManagedElement><MpsType>2</MpsType><OrderOfAccess>0</OrderOfAccess></AMT_RemoteAccessPolicyAppliesToMPS><AMT_RemoteAccessPolicyAppliesToMPS><PolicySet><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</w:Selector><w:Selector Name="PolicyRuleName">UserInitiated</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet></a:ReferenceParameters></PolicySet><ManagedElement><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</w:Selector><w:Selector Name="Name">Intel(r) AMT:Management Presence Server 0</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet></a:ReferenceParameters></ManagedElement><MpsType>2</MpsType><OrderOfAccess>0</OrderOfAccess></AMT_RemoteAccessPolicyAppliesToMPS></Items><EndOfSequence></EndOfSequence></PullResponse></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

F7E
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</w:ResourceURI><a:MessageID>28</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><wsman:SelectorSet xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><wsman:Selector Name="ManagedElement"><EndpointReference xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</Selector><Selector Name="Name">Intel(r) AMT:Management Presence Server 0</Selector><Selector Name="SystemCreationClassName">CIM_ComputerSystem</Selector><Selector Name="SystemName">Intel(r) AMT</Selector></SelectorSet></ReferenceParameters></EndpointReference></wsman:Selector><wsman:Selector Name="PolicySet"><EndpointReference xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</Selector><Selector Name="PolicyRuleName">Periodic</Selector><Selector Name="SystemCreationClassName">CIM_ComputerSystem</Selector><Selector Name="SystemName">Intel(r) AMT</Selector></SelectorSet></ReferenceParameters></EndpointReference></wsman:Selector></wsman:SelectorSet></Header><Body><h:AMT_RemoteAccessPolicyAppliesToMPS xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><h:PolicySet><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</w:Selector><w:Selector Name="PolicyRuleName">Periodic</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet></a:ReferenceParameters></h:PolicySet><h:ManagedElement><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</w:Selector><w:Selector Name="Name">Intel(r) AMT:Management Presence Server 0</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ManagedElement><h:MpsType>2</h:MpsType><h:OrderOfAccess>0</h:OrderOfAccess></h:AMT_RemoteAccessPolicyAppliesToMPS></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/PutResponse</a:Action>
    <a:RelatesTo>28</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>5f694e41-c4f8-43fd-a01a-dbd4dcc1a78d</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</w:ResourceURI>
  </a:Header>
  <s:Body><a:PutResponse xmlns:a="http://schemas.xmlsoap.org/ws/2004/09/transfer"/></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

F88
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</w:ResourceURI><a:MessageID>29</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><wsman:SelectorSet xmlns:wsman="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><wsman:Selector Name="ManagedElement"><EndpointReference xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</Selector><Selector Name="Name">Intel(r) AMT:Management Presence Server 0</Selector><Selector Name="SystemCreationClassName">CIM_ComputerSystem</Selector><Selector Name="SystemName">Intel(r) AMT</Selector></SelectorSet></ReferenceParameters></EndpointReference></wsman:Selector><wsman:Selector Name="PolicySet"><EndpointReference xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><Address xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</Address><ReferenceParameters xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing"><ResourceURI xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd">http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</ResourceURI><SelectorSet xmlns="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</Selector><Selector Name="PolicyRuleName">UserInitiated</Selector><Selector Name="SystemCreationClassName">CIM_ComputerSystem</Selector><Selector Name="SystemName">Intel(r) AMT</Selector></SelectorSet></ReferenceParameters></EndpointReference></wsman:Selector></wsman:SelectorSet></Header><Body><h:AMT_RemoteAccessPolicyAppliesToMPS xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS"><h:PolicySet><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">AMT_RemoteAccessPolicyRule</w:Selector><w:Selector Name="PolicyRuleName">UserInitiated</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet></a:ReferenceParameters></h:PolicySet><h:ManagedElement><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address><a:ReferenceParameters><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP</w:ResourceURI><w:SelectorSet><w:Selector Name="CreationClassName">AMT_ManagementPresenceRemoteSAP</w:Selector><w:Selector Name="Name">Intel(r) AMT:Management Presence Server 0</w:Selector><w:Selector Name="SystemCreationClassName">CIM_ComputerSystem</w:Selector><w:Selector Name="SystemName">Intel(r) AMT</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ManagedElement><h:MpsType>2</h:MpsType><h:OrderOfAccess>0</h:OrderOfAccess></h:AMT_RemoteAccessPolicyAppliesToMPS></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/PutResponse</a:Action>
    <a:RelatesTo>29</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>4eba5fe7-1d9d-4778-8355-88f55b2a100c</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS</w:ResourceURI>
  </a:Header>
  <s:Body><a:PutResponse xmlns:a="http://schemas.xmlsoap.org/ws/2004/09/transfer"/></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

3C7
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService/RequestStateChange</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService</w:ResourceURI><a:MessageID>30</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:RequestStateChange_INPUT xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService"><h:RequestedState>32771</h:RequestedState></h:RequestStateChange_INPUT></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService/RequestStateChange
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService/RequestStateChangeResponse</a:Action>
    <a:RelatesTo>30</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>29c5c1e0-a9e7-40e4-92cb-e76f3f15f76f</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService</w:ResourceURI>
  </a:Header>
  <s:Body><RequestStateChange_OUTPUT><ReturnValue>0</ReturnValue></RequestStateChange_OUTPUT></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

2DF
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData</w:ResourceURI><a:MessageID>31</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/GetResponse</a:Action>
    <a:RelatesTo>31</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>f34d0204-989a-4a5d-9755-561affcc5ed9</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData</w:ResourceURI>
  </a:Header>
  <s:Body><AMT_EnvironmentDetectionSettingData><ElementName>Intel(r) AMT Environment Detection Settings</ElementName><InstanceID>Intel(r) AMT:EnvironmentDetectionSettingData 0</InstanceID></AMT_EnvironmentDetectionSettingData></s:Body>
</s:Envelope>
2026/02/12 15:35:53 [DEBUG] Raw WSMAN Request: POST /wsman HTTP/1.1
Host: 1d8aa800-918d-11e4-b33f-123400000000:16992
Transfer-Encoding: chunked

4E5
<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData</w:ResourceURI><a:MessageID>32</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name="InstanceID">Intel(r) AMT:EnvironmentDetectionSettingData 0</w:Selector></w:SelectorSet></Header><Body><h:AMT_EnvironmentDetectionSettingData xmlns:h="http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData"><h:ElementName>Intel(r) AMT Environment Detection Settings</h:ElementName><h:InstanceID>Intel(r) AMT:EnvironmentDetectionSettingData 0</h:InstanceID><h:DetectionStrings>e1b2d0f7-e443-4eae-b8a0-6d03dd2a6282.com</h:DetectionStrings></h:AMT_EnvironmentDetectionSettingData></Body></Envelope>
0

2026/02/12 15:35:53 [INFO]  Received Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:35:53 [DEBUG] Received ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData
2026/02/12 15:35:53 [DEBUG] WSMAN Response XML: <?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" 
            xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" 
            xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" 
            xmlns:wsen="http://schemas.xmlsoap.org/ws/2004/09/enumeration"
            xmlns:amt="http://intel.com/wbem/wscim/1/amt-schema/1"
            xmlns:ips="http://intel.com/wbem/wscim/1/ips-schema/1"
            xmlns:cim="http://schemas.dmtf.org/wbem/wscim/1/common"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <a:Header>
    <a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/PutResponse</a:Action>
    <a:RelatesTo>32</a:RelatesTo>
    <a:To>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:To>
    <a:MessageID>f5749b1e-a3f9-43b7-b6e6-53430c416269</a:MessageID>
    <w:ResourceURI>http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData</w:ResourceURI>
  </a:Header>
  <s:Body><AMT_EnvironmentDetectionSettingData><ElementName>Intel(r) AMT Environment Detection Settings</ElementName><InstanceID>Intel(r) AMT:EnvironmentDetectionSettingData 0</InstanceID></AMT_EnvironmentDetectionSettingData></s:Body>
</s:Envelope>
```
