This is what a successful RPS activation looked like:

```
$ ./rpcsim activate -u wss://198.0.0.61:443/activate -n --profile scottccm
Enter AMT Password: 
2026/02/12 15:31:27 [INFO]  Activating against wss://198.0.0.61:443/activate with profile scottccm
2026/02/12 15:31:27 Connecting to wss://198.0.0.61:443/activate...
2026/02/12 15:31:27 Received Method: wsman Status: ok
2026/02/12 15:31:27 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:31:27 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_GeneralSettings
2026/02/12 15:31:27 Received Method: wsman Status: ok
2026/02/12 15:31:27 [INFO]  Processing WSMAN Action: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/Setup
2026/02/12 15:31:27 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyAccessPoint
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyRule
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TLSSettingData
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyCertificate
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EthernetPortSettings
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:31:39 [INFO]  ResourceURI: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService/RequestStateChange
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP/RequestStateChange
2026/02/12 15:31:39 [INFO]  ResourceURI: http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_KVMRedirectionSAP
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RedirectionService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService/AddTrustedRootCertificate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_PublicKeyManagementService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddMpServer
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_ManagementPresenceRemoteSAP
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRule
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService/AddRemoteAccessPolicyRule
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_RemoteAccessPolicyAppliesToMPS
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService/RequestStateChange
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_UserInitiatedConnectionService
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Get
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData
2026/02/12 15:31:39 Received Method: wsman Status: ok
2026/02/12 15:31:39 [INFO]  Processing WSMAN Action: http://schemas.xmlsoap.org/ws/2004/09/transfer/Put
2026/02/12 15:31:39 [INFO]  ResourceURI: http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData
2026/02/12 15:31:39 Received Method: success Status: success
2026/02/12 15:31:39 RPS indicates Success!
2026/02/12 15:31:39 [INFO]  Activation Sequence Completed Successfully.
```
