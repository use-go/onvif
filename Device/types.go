package Device

import (
	"github.com/yakovlevdmv/goonvif/xsdTypes"
	"github.com/yakovlevdmv/goonvif/xsdTypes/onvif"
)

type Service struct {
	Namespace xsdTypes.AnyURI
	XAddr xsdTypes.AnyURI
	Capabilities
	Version onvif.OnvifVersion
}

type Capabilities struct {
	Any string
}

type DeviceServiceCapabilities struct {
	Network NetworkCapabilities
	Security SecurityCapabilities
	System SystemCapabilities
	Misc MiscCapabilities
}

type NetworkCapabilities struct {
	IPFilter 			xsdTypes.Boolean `xml:"IPFilter,attr"`
	ZeroConfiguration 	xsdTypes.Boolean `xml:"ZeroConfiguration,attr"`
	IPVersion6 			xsdTypes.Boolean `xml:"IPVersion6,attr"`
	DynDNS 				xsdTypes.Boolean `xml:"DynDNS,attr"`
	Dot11Configuration 	xsdTypes.Boolean `xml:"Dot11Configuration,attr"`
	Dot1XConfigurations int 	`xml:"Dot1XConfigurations,attr"`
	HostnameFromDHCP	xsdTypes.Boolean `xml:"HostnameFromDHCP,attr"`
	NTP 				int 	`xml:"NTP,attr"`
	DHCPv6 				xsdTypes.Boolean `xml:"DHCPv6,attr"`
}

type SecurityCapabilities struct {
	TLS1_0 					xsdTypes.Boolean 		`xml:"TLS1_0,attr"`
	TLS1_1 					xsdTypes.Boolean 		`xml:"TLS1_1,attr"`
	TLS1_2 					xsdTypes.Boolean 		`xml:"TLS1_2,attr"`
	OnboardKeyGeneration 	xsdTypes.Boolean 		`xml:"OnboardKeyGeneration,attr"`
	AccessPolicyConfig 		xsdTypes.Boolean 		`xml:"AccessPolicyConfig,attr"`
	DefaultAccessPolicy 	xsdTypes.Boolean 		`xml:"DefaultAccessPolicy,attr"`
	Dot1X 					xsdTypes.Boolean 		`xml:"Dot1X,attr"`
	RemoteUserHandling 		xsdTypes.Boolean 		`xml:"RemoteUserHandling,attr"`
	X_509Token 				xsdTypes.Boolean 		`xml:"X_509Token,attr"`
	SAMLToken 				xsdTypes.Boolean 		`xml:"SAMLToken,attr"`
	KerberosToken 			xsdTypes.Boolean 		`xml:"KerberosToken,attr"`
	UsernameToken 			xsdTypes.Boolean 		`xml:"UsernameToken,attr"`
	HttpDigest 				xsdTypes.Boolean 		`xml:"HttpDigest,attr"`
	RELToken 				xsdTypes.Boolean 		`xml:"RELToken,attr"`
	SupportedEAPMethods 	EAPMethodTypes 			`xml:"SupportedEAPMethods,attr"`
	MaxUsers 				int 					`xml:"MaxUsers,attr"`
	MaxUserNameLength 		int 					`xml:"MaxUserNameLength,attr"`
	MaxPasswordLength 		int 					`xml:"MaxPasswordLength,attr"`
}

type EAPMethodTypes struct {
	Types []int
}

type SystemCapabilities struct {
	DiscoveryResolve 			xsdTypes.Boolean 		`xml:"DiscoveryResolve,attr"`
	DiscoveryBye 				xsdTypes.Boolean 		`xml:"DiscoveryBye,attr"`
	RemoteDiscovery 			xsdTypes.Boolean 		`xml:"RemoteDiscovery,attr"`
	SystemBackup 				xsdTypes.Boolean 		`xml:"SystemBackup,attr"`
	SystemLogging 				xsdTypes.Boolean 		`xml:"SystemLogging,attr"`
	FirmwareUpgrade 			xsdTypes.Boolean 		`xml:"FirmwareUpgrade,attr"`
	HttpFirmwareUpgrade 		xsdTypes.Boolean 		`xml:"HttpFirmwareUpgrade,attr"`
	HttpSystemBackup 			xsdTypes.Boolean 		`xml:"HttpSystemBackup,attr"`
	HttpSystemLogging 			xsdTypes.Boolean 		`xml:"HttpSystemLogging,attr"`
	HttpSupportInformation 		xsdTypes.Boolean 		`xml:"HttpSupportInformation,attr"`
	StorageConfiguration 		xsdTypes.Boolean 		`xml:"StorageConfiguration,attr"`
	MaxStorageConfigurations 	int 					`xml:"MaxStorageConfigurations,attr"`
	GeoLocationEntries 			int 					`xml:"GeoLocationEntries,attr"`
	AutoGeo 					onvif.StringAttrList 			`xml:"AutoGeo,attr"`
}

type MiscCapabilities struct {
	AuxiliaryCommands onvif.StringAttrList `xml:"AuxiliaryCommands,attr"`
}

//Device main types

type GetServices struct {
	IncludeCapability xsdTypes.Boolean

}


type GetServicesResponse struct {
	Service Service

}


type GetServiceCapabilities struct {

}


type GetServiceCapabilitiesResponse struct {
	Capabilities DeviceServiceCapabilities

}


type GetDeviceInformation struct {

}


type GetDeviceInformationResponse struct {
	Manufacturer string
	Model string
	FirmwareVersion string
	SerialNumber string
	HardwareId string

}


type SetSystemDateAndTime struct {
	DateTimeType onvif.SetDateTimeType
	DaylightSavings xsdTypes.Boolean
	TimeZone onvif.TimeZone
	UTCDateTime xsdTypes.DateTime
}


type SetSystemDateAndTimeResponse struct {

}


type GetSystemDateAndTime struct {

}


type GetSystemDateAndTimeResponse struct {
	SystemDateAndTime onvif.SystemDateTime

}


type SetSystemFactoryDefault struct {
	FactoryDefault onvif.FactoryDefaultType

}


type SetSystemFactoryDefaultResponse struct {

}


type UpgradeSystemFirmware struct {
	Firmware onvif.AttachmentData

}


type UpgradeSystemFirmwareResponse struct {
	Message string

}


type SystemReboot struct {

}


type SystemRebootResponse struct {
	Message string

}


type RestoreSystem struct {
	BackupFiles onvif.BackupFile

}


type RestoreSystemResponse struct {

}


type GetSystemBackup struct {

}


type GetSystemBackupResponse struct {
	BackupFiles onvif.BackupFile

}


type GetSystemLog struct {
	LogType onvif.SystemLogType

}


type GetSystemLogResponse struct {
	SystemLog onvif.SystemLog

}


type GetSystemSupportInformation struct {

}


type GetSystemSupportInformationResponse struct {
	SupportInformation onvif.SupportInformation

}


type GetScopes struct {

}


type GetScopesResponse struct {
	Scopes onvif.Scope

}


type SetScopes struct {
	Scopes xsdTypes.AnyURI

}


type SetScopesResponse struct {

}


type AddScopes struct {
	ScopeItem xsdTypes.AnyURI

}


type AddScopesResponse struct {

}


type RemoveScopes struct {
	ScopeItem xsdTypes.AnyURI

}


type RemoveScopesResponse struct {
	ScopeItem xsdTypes.AnyURI

}


type GetDiscoveryMode struct {

}


type GetDiscoveryModeResponse struct {
	DiscoveryMode onvif.DiscoveryMode

}


type SetDiscoveryMode struct {
	DiscoveryMode onvif.DiscoveryMode

}


type SetDiscoveryModeResponse struct {

}


type GetRemoteDiscoveryMode struct {

}


type GetRemoteDiscoveryModeResponse struct {
	RemoteDiscoveryMode onvif.DiscoveryMode

}


type SetRemoteDiscoveryMode struct {
	RemoteDiscoveryMode onvif.DiscoveryMode

}


type SetRemoteDiscoveryModeResponse struct {

}


type GetDPAddresses struct {

}


type GetDPAddressesResponse struct {
	DPAddress onvif.NetworkHost

}


type SetDPAddresses struct {
	DPAddress onvif.NetworkHost

}


type SetDPAddressesResponse struct {

}


type GetEndpointReference struct {

}


type GetEndpointReferenceResponse struct {
	GUID string

}


type GetRemoteUser struct {

}


type GetRemoteUserResponse struct {
	RemoteUser onvif.RemoteUser

}


type SetRemoteUser struct {
	RemoteUser onvif.RemoteUser

}


type SetRemoteUserResponse struct {

}


type GetUsers struct {

}


type GetUsersResponse struct {
	User onvif.User

}


type CreateUsers struct {
	User onvif.User

}


type CreateUsersResponse struct {

}


type DeleteUsers struct {
	Username string

}


type DeleteUsersResponse struct {

}


type SetUser struct {
	User onvif.User

}


type SetUserResponse struct {

}


type GetWsdlUrl struct {

}


type GetWsdlUrlResponse struct {
	WsdlUrl xsdTypes.AnyURI

}


type GetCapabilities struct {
	Category onvif.CapabilityCategory

}


type GetCapabilitiesResponse struct {
	Capabilities onvif.Capabilities

}


type GetHostname struct {

}


type GetHostnameResponse struct {
	HostnameInformation onvif.HostnameInformation

}


type SetHostname struct {
	Name xsdTypes.token

}


type SetHostnameResponse struct {

}


type SetHostnameFromDHCP struct {
	FromDHCP xsdTypes.Boolean

}


type SetHostnameFromDHCPResponse struct {
	RebootNeeded xsdTypes.Boolean

}


type GetDNS struct {

}


type GetDNSResponse struct {
	DNSInformation onvif.DNSInformation

}


type SetDNS struct {
	FromDHCP 		xsdTypes.Boolean
	SearchDomain 	xsdTypes.token
	DNSManual 		onvif.IPAddress

}


type SetDNSResponse struct {

}


type GetNTP struct {

}


type GetNTPResponse struct {
	NTPInformation NTPInformation

}


type SetNTP struct {
	FromDHCP xsdTypes.Boolean
	NTPManual NetworkHost

}


type SetNTPResponse struct {

}


type GetDynamicDNS struct {

}


type GetDynamicDNSResponse struct {
	DynamicDNSInformation DynamicDNSInformation

}


type SetDynamicDNS struct {
	Type 	DynamicDNSType
	Name 	onvif.DNSName
	TTL 	xsdTypes.Duration

}


type SetDynamicDNSResponse struct {

}


type GetNetworkInterfaces struct {

}


type GetNetworkInterfacesResponse struct {
	NetworkInterfaces NetworkInterface

}


type SetNetworkInterfaces struct {
	InterfaceToken onvif.ReferenceToken
	NetworkInterface NetworkInterfaceSetConfiguration

}


type SetNetworkInterfacesResponse struct {
	RebootNeeded xsdTypes.Boolean

}


type GetNetworkProtocols struct {

}


type GetNetworkProtocolsResponse struct {
	NetworkProtocols NetworkProtocol

}


type SetNetworkProtocols struct {
	NetworkProtocols NetworkProtocol

}


type SetNetworkProtocolsResponse struct {

}


type GetNetworkDefaultGateway struct {

}


type GetNetworkDefaultGatewayResponse struct {
	NetworkGateway NetworkGateway

}


type SetNetworkDefaultGateway struct {
	IPv4Address onvif.IPv4Address
	IPv6Address onvif.IPv6Address

}


type SetNetworkDefaultGatewayResponse struct {

}


type GetZeroConfiguration struct {

}


type GetZeroConfigurationResponse struct {
	ZeroConfiguration NetworkZeroConfiguration

}


type SetZeroConfiguration struct {
	InterfaceToken 	onvif.ReferenceToken
	Enabled 		xsdTypes.Boolean

}


type SetZeroConfigurationResponse struct {

}


type GetIPAddressFilter struct {

}


type GetIPAddressFilterResponse struct {
	IPAddressFilter onvif.IPAddressFilter

}


type SetIPAddressFilter struct {
	IPAddressFilter onvif.IPAddressFilter

}


type SetIPAddressFilterResponse struct {

}


type AddIPAddressFilter struct {
	IPAddressFilter onvif.IPAddressFilter

}


type AddIPAddressFilterResponse struct {

}


type RemoveIPAddressFilter struct {
	IPAddressFilter onvif.IPAddressFilter

}


type RemoveIPAddressFilterResponse struct {

}


type GetAccessPolicy struct {

}


type GetAccessPolicyResponse struct {
	PolicyFile BinaryData

}


type SetAccessPolicy struct {
	PolicyFile BinaryData

}


type SetAccessPolicyResponse struct {

}


type CreateCertificate struct {
	CertificateID 	xsdTypes.token
	Subject 		string
	ValidNotBefore 	xsdTypes.DateTime
	ValidNotAfter 	xsdTypes.DateTime

}


type CreateCertificateResponse struct {
	NvtCertificate Certificate

}


type GetCertificates struct {

}


type GetCertificatesResponse struct {
	NvtCertificate Certificate

}


type GetCertificatesStatus struct {

}


type GetCertificatesStatusResponse struct {
	CertificateStatus CertificateStatus

}


type SetCertificatesStatus struct {
	CertificateStatus CertificateStatus

}


type SetCertificatesStatusResponse struct {

}


type DeleteCertificates struct {
	CertificateID token

}


type DeleteCertificatesResponse struct {

}


type GetPkcs10Request struct {
	CertificateID token
	Subject string
	Attributes BinaryData

}


type GetPkcs10RequestResponse struct {
	Pkcs10Request BinaryData

}


type LoadCertificates struct {
	NVTCertificate Certificate

}


type LoadCertificatesResponse struct {

}


type GetClientCertificateMode struct {

}


type GetClientCertificateModeResponse struct {
	Enabled boolean

}


type SetClientCertificateMode struct {
	Enabled boolean

}


type SetClientCertificateModeResponse struct {

}


type GetRelayOutputs struct {

}


type GetRelayOutputsResponse struct {
	RelayOutputs RelayOutput

}


type SetRelayOutputSettings struct {
	RelayOutputToken ReferenceToken
	Properties RelayOutputSettings

}


type SetRelayOutputSettingsResponse struct {

}


type SetRelayOutputState struct {
	RelayOutputToken ReferenceToken
	LogicalState RelayLogicalState

}


type SetRelayOutputStateResponse struct {

}


type SendAuxiliaryCommand struct {
	AuxiliaryCommand AuxiliaryData

}


type SendAuxiliaryCommandResponse struct {
	AuxiliaryCommandResponse AuxiliaryData

}


type GetCACertificates struct {

}


type GetCACertificatesResponse struct {
	CACertificate Certificate

}


type LoadCertificateWithPrivateKey struct {
	CertificateWithPrivateKey CertificateWithPrivateKey

}


type LoadCertificateWithPrivateKeyResponse struct {

}


type GetCertificateInformation struct {
	CertificateID token

}


type GetCertificateInformationResponse struct {
	CertificateInformation CertificateInformation

}


type LoadCACertificates struct {
	CACertificate Certificate

}


type LoadCACertificatesResponse struct {

}


type CreateDot1XConfiguration struct {
	Dot1XConfiguration Dot1XConfiguration

}


type CreateDot1XConfigurationResponse struct {

}


type SetDot1XConfiguration struct {
	Dot1XConfiguration Dot1XConfiguration

}


type SetDot1XConfigurationResponse struct {

}


type GetDot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken

}


type GetDot1XConfigurationResponse struct {
	Dot1XConfiguration Dot1XConfiguration

}


type GetDot1XConfigurations struct {

}


type GetDot1XConfigurationsResponse struct {
	Dot1XConfiguration Dot1XConfiguration

}


type DeleteDot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken

}


type DeleteDot1XConfigurationResponse struct {

}


type GetDot11Capabilities struct {

}


type GetDot11CapabilitiesResponse struct {
	Capabilities Dot11Capabilities

}


type GetDot11Status struct {
	InterfaceToken ReferenceToken

}


type GetDot11StatusResponse struct {
	Status Dot11Status

}


type ScanAvailableDot11Networks struct {
	InterfaceToken ReferenceToken

}


type ScanAvailableDot11NetworksResponse struct {
	Networks Dot11AvailableNetworks

}


type GetSystemUris struct {

}


type GetSystemUrisResponse struct {
	SystemLogUris SystemLogUriList
	SupportInfoUri anyURI
	SystemBackupUri anyURI
	Extension

}


type StartFirmwareUpgrade struct {

}


type StartFirmwareUpgradeResponse struct {
	UploadUri anyURI
	UploadDelay duration
	ExpectedDownTime duration

}


type StartSystemRestore struct {

}


type StartSystemRestoreResponse struct {
	UploadUri anyURI
	ExpectedDownTime duration

}


type GetStorageConfigurations struct {

}


type GetStorageConfigurationsResponse struct {
	StorageConfigurations tds:StorageConfiguration

}


type CreateStorageConfiguration struct {
	StorageConfiguration tds:StorageConfigurationData

}


type CreateStorageConfigurationResponse struct {
	Token ReferenceToken

}


type GetStorageConfiguration struct {
	Token ReferenceToken

}


type GetStorageConfigurationResponse struct {
	StorageConfiguration tds:StorageConfiguration

}


type SetStorageConfiguration struct {
	StorageConfiguration tds:StorageConfiguration

}


type SetStorageConfigurationResponse struct {

}


type DeleteStorageConfiguration struct {
	Token ReferenceToken

}


type DeleteStorageConfigurationResponse struct {

}


type GetGeoLocation struct {

}


type GetGeoLocationResponse struct {
	Location LocationEntity

}


type SetGeoLocation struct {
	Location LocationEntity

}


type SetGeoLocationResponse struct {

}


type DeleteGeoLocation struct {
	Location LocationEntity

}


type DeleteGeoLocationResponse struct {

}