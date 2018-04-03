package Device

import (
	"github.com/yakovlevdmv/goonvif/xsd"
	"github.com/yakovlevdmv/goonvif/xsd/onvif"
)

type Service struct {
	Namespace xsd.AnyURI
	XAddr     xsd.AnyURI
	Capabilities
	Version   onvif.OnvifVersion
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
	IPFilter            xsd.Boolean `xml:"IPFilter,attr"`
	ZeroConfiguration   xsd.Boolean `xml:"ZeroConfiguration,attr"`
	IPVersion6          xsd.Boolean `xml:"IPVersion6,attr"`
	DynDNS              xsd.Boolean `xml:"DynDNS,attr"`
	Dot11Configuration  xsd.Boolean `xml:"Dot11Configuration,attr"`
	Dot1XConfigurations int         `xml:"Dot1XConfigurations,attr"`
	HostnameFromDHCP    xsd.Boolean `xml:"HostnameFromDHCP,attr"`
	NTP                 int         `xml:"NTP,attr"`
	DHCPv6              xsd.Boolean `xml:"DHCPv6,attr"`
}

type SecurityCapabilities struct {
	TLS1_0               xsd.Boolean    `xml:"TLS1_0,attr"`
	TLS1_1               xsd.Boolean    `xml:"TLS1_1,attr"`
	TLS1_2               xsd.Boolean    `xml:"TLS1_2,attr"`
	OnboardKeyGeneration xsd.Boolean    `xml:"OnboardKeyGeneration,attr"`
	AccessPolicyConfig   xsd.Boolean    `xml:"AccessPolicyConfig,attr"`
	DefaultAccessPolicy  xsd.Boolean    `xml:"DefaultAccessPolicy,attr"`
	Dot1X                xsd.Boolean    `xml:"Dot1X,attr"`
	RemoteUserHandling   xsd.Boolean    `xml:"RemoteUserHandling,attr"`
	X_509Token           xsd.Boolean    `xml:"X_509Token,attr"`
	SAMLToken            xsd.Boolean    `xml:"SAMLToken,attr"`
	KerberosToken        xsd.Boolean    `xml:"KerberosToken,attr"`
	UsernameToken        xsd.Boolean    `xml:"UsernameToken,attr"`
	HttpDigest           xsd.Boolean    `xml:"HttpDigest,attr"`
	RELToken             xsd.Boolean    `xml:"RELToken,attr"`
	SupportedEAPMethods  EAPMethodTypes `xml:"SupportedEAPMethods,attr"`
	MaxUsers             int            `xml:"MaxUsers,attr"`
	MaxUserNameLength    int            `xml:"MaxUserNameLength,attr"`
	MaxPasswordLength    int            `xml:"MaxPasswordLength,attr"`
}

type EAPMethodTypes struct {
	Types []int
}

type SystemCapabilities struct {
	DiscoveryResolve         xsd.Boolean          `xml:"DiscoveryResolve,attr"`
	DiscoveryBye             xsd.Boolean          `xml:"DiscoveryBye,attr"`
	RemoteDiscovery          xsd.Boolean          `xml:"RemoteDiscovery,attr"`
	SystemBackup             xsd.Boolean          `xml:"SystemBackup,attr"`
	SystemLogging            xsd.Boolean          `xml:"SystemLogging,attr"`
	FirmwareUpgrade          xsd.Boolean          `xml:"FirmwareUpgrade,attr"`
	HttpFirmwareUpgrade      xsd.Boolean          `xml:"HttpFirmwareUpgrade,attr"`
	HttpSystemBackup         xsd.Boolean          `xml:"HttpSystemBackup,attr"`
	HttpSystemLogging        xsd.Boolean          `xml:"HttpSystemLogging,attr"`
	HttpSupportInformation   xsd.Boolean          `xml:"HttpSupportInformation,attr"`
	StorageConfiguration     xsd.Boolean          `xml:"StorageConfiguration,attr"`
	MaxStorageConfigurations int                  `xml:"MaxStorageConfigurations,attr"`
	GeoLocationEntries       int                  `xml:"GeoLocationEntries,attr"`
	AutoGeo                  onvif.StringAttrList `xml:"AutoGeo,attr"`
}

type MiscCapabilities struct {
	AuxiliaryCommands onvif.StringAttrList `xml:"AuxiliaryCommands,attr"`
}

type StorageConfiguration struct {
	onvif.DeviceEntity
	Data StorageConfigurationData
}

type StorageConfigurationData struct {
	Type xsd.String `xml:"type,attr"`
	LocalPath xsd.AnyURI
	StorageUri xsd.AnyURI
	User UserCredential
	Extension xsd.AnyURI
}

type UserCredential struct {
	UserName xsd.String
	Password xsd.String
	Extension xsd.AnyType
}

//Device main types

type GetServices struct {
	IncludeCapability xsd.Boolean

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
	XMLName string `xml:"wsdl:GetDeviceInformation"`
	nmsp string `xml:"wsdl:"`
}


type GetDeviceInformationResponse struct {
	Manufacturer string
	Model string
	FirmwareVersion string
	SerialNumber string
	HardwareId string

}


type SetSystemDateAndTime struct {
	DateTimeType    onvif.SetDateTimeType
	DaylightSavings xsd.Boolean
	TimeZone        onvif.TimeZone
	UTCDateTime     xsd.DateTime
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
	Scopes xsd.AnyURI

}


type SetScopesResponse struct {

}


type AddScopes struct {
	ScopeItem xsd.AnyURI

}


type AddScopesResponse struct {

}


type RemoveScopes struct {
	ScopeItem xsd.AnyURI

}


type RemoveScopesResponse struct {
	ScopeItem xsd.AnyURI

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
	WsdlUrl xsd.AnyURI

}


type GetCapabilities struct {
	Category onvif.CapabilityCategory

}


type GetCapabilitiesResponse struct {
	Capabilities onvif.Capabilities

}


type GetHostname struct {
	XMLName string `xml:"wsdl:GetHostname"`
}


type GetHostnameResponse struct {
	HostnameInformation onvif.HostnameInformation

}


type SetHostname struct {
	Name xsd.Token

}


type SetHostnameResponse struct {

}


type SetHostnameFromDHCP struct {
	FromDHCP xsd.Boolean

}


type SetHostnameFromDHCPResponse struct {
	RebootNeeded xsd.Boolean

}


type GetDNS struct {

}


type GetDNSResponse struct {
	DNSInformation onvif.DNSInformation

}


type SetDNS struct {
	FromDHCP     xsd.Boolean
	SearchDomain xsd.Token
	DNSManual    onvif.IPAddress

}


type SetDNSResponse struct {

}


type GetNTP struct {

}


type GetNTPResponse struct {
	NTPInformation onvif.NTPInformation

}


type SetNTP struct {
	FromDHCP  xsd.Boolean
	NTPManual onvif.NetworkHost

}


type SetNTPResponse struct {

}


type GetDynamicDNS struct {

}


type GetDynamicDNSResponse struct {
	DynamicDNSInformation onvif.DynamicDNSInformation

}


type SetDynamicDNS struct {
	Type onvif.DynamicDNSType
	Name onvif.DNSName
	TTL  xsd.Duration

}


type SetDynamicDNSResponse struct {

}


type GetNetworkInterfaces struct {

}


type GetNetworkInterfacesResponse struct {
	NetworkInterfaces onvif.NetworkInterface

}


type SetNetworkInterfaces struct {
	InterfaceToken onvif.ReferenceToken
	NetworkInterface onvif.NetworkInterfaceSetConfiguration

}


type SetNetworkInterfacesResponse struct {
	RebootNeeded xsd.Boolean

}


type GetNetworkProtocols struct {

}


type GetNetworkProtocolsResponse struct {
	NetworkProtocols onvif.NetworkProtocol

}


type SetNetworkProtocols struct {
	NetworkProtocols onvif.NetworkProtocol

}


type SetNetworkProtocolsResponse struct {

}


type GetNetworkDefaultGateway struct {

}


type GetNetworkDefaultGatewayResponse struct {
	NetworkGateway onvif.NetworkGateway

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
	ZeroConfiguration onvif.NetworkZeroConfiguration

}


type SetZeroConfiguration struct {
	InterfaceToken onvif.ReferenceToken
	Enabled        xsd.Boolean

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
	PolicyFile onvif.BinaryData

}


type SetAccessPolicy struct {
	PolicyFile onvif.BinaryData

}


type SetAccessPolicyResponse struct {

}


type CreateCertificate struct {
	CertificateID  xsd.Token
	Subject        string
	ValidNotBefore xsd.DateTime
	ValidNotAfter  xsd.DateTime

}


type CreateCertificateResponse struct {
	NvtCertificate onvif.Certificate

}


type GetCertificates struct {

}


type GetCertificatesResponse struct {
	NvtCertificate onvif.Certificate

}


type GetCertificatesStatus struct {

}


type GetCertificatesStatusResponse struct {
	CertificateStatus onvif.CertificateStatus

}


type SetCertificatesStatus struct {
	CertificateStatus onvif.CertificateStatus

}


type SetCertificatesStatusResponse struct {

}


type DeleteCertificates struct {
	CertificateID xsd.Token

}


type DeleteCertificatesResponse struct {

}


type GetPkcs10Request struct {
	CertificateID xsd.Token
	Subject xsd.String
	Attributes onvif.BinaryData

}


type GetPkcs10RequestResponse struct {
	Pkcs10Request onvif.BinaryData

}


type LoadCertificates struct {
	NVTCertificate onvif.Certificate

}


type LoadCertificatesResponse struct {

}


type GetClientCertificateMode struct {

}


type GetClientCertificateModeResponse struct {
	Enabled xsd.Boolean

}


type SetClientCertificateMode struct {
	Enabled xsd.Boolean

}


type SetClientCertificateModeResponse struct {

}


type GetRelayOutputs struct {

}


type GetRelayOutputsResponse struct {
	RelayOutputs onvif.RelayOutput

}


type SetRelayOutputSettings struct {
	RelayOutputToken onvif.ReferenceToken
	Properties onvif.RelayOutputSettings

}


type SetRelayOutputSettingsResponse struct {

}


type SetRelayOutputState struct {
	RelayOutputToken onvif.ReferenceToken
	LogicalState onvif.RelayLogicalState

}


type SetRelayOutputStateResponse struct {

}


type SendAuxiliaryCommand struct {
	AuxiliaryCommand onvif.AuxiliaryData

}


type SendAuxiliaryCommandResponse struct {
	AuxiliaryCommandResponse onvif.AuxiliaryData

}


type GetCACertificates struct {

}


type GetCACertificatesResponse struct {
	CACertificate onvif.Certificate

}


type LoadCertificateWithPrivateKey struct {
	CertificateWithPrivateKey onvif.CertificateWithPrivateKey

}


type LoadCertificateWithPrivateKeyResponse struct {

}


type GetCertificateInformation struct {
	CertificateID xsd.Token

}


type GetCertificateInformationResponse struct {
	CertificateInformation onvif.CertificateInformation

}


type LoadCACertificates struct {
	CACertificate onvif.Certificate

}


type LoadCACertificatesResponse struct {

}


type CreateDot1XConfiguration struct {
	Dot1XConfiguration onvif.Dot1XConfiguration

}


type CreateDot1XConfigurationResponse struct {

}


type SetDot1XConfiguration struct {
	Dot1XConfiguration onvif.Dot1XConfiguration

}


type SetDot1XConfigurationResponse struct {

}


type GetDot1XConfiguration struct {
	Dot1XConfigurationToken onvif.ReferenceToken

}


type GetDot1XConfigurationResponse struct {
	Dot1XConfiguration onvif.Dot1XConfiguration

}


type GetDot1XConfigurations struct {

}


type GetDot1XConfigurationsResponse struct {
	Dot1XConfiguration onvif.Dot1XConfiguration

}


type DeleteDot1XConfiguration struct {
	Dot1XConfigurationToken onvif.ReferenceToken

}


type DeleteDot1XConfigurationResponse struct {

}


type GetDot11Capabilities struct {

}


type GetDot11CapabilitiesResponse struct {
	Capabilities onvif.Dot11Capabilities

}


type GetDot11Status struct {
	InterfaceToken onvif.ReferenceToken

}


type GetDot11StatusResponse struct {
	Status onvif.Dot11Status

}


type ScanAvailableDot11Networks struct {
	InterfaceToken onvif.ReferenceToken

}


type ScanAvailableDot11NetworksResponse struct {
	Networks onvif.Dot11AvailableNetworks

}


type GetSystemUris struct {

}


type GetSystemUrisResponse struct {
	SystemLogUris onvif.SystemLogUriList
	SupportInfoUri xsd.AnyURI
	SystemBackupUri xsd.AnyURI
	Extension xsd.AnyType
}


type StartFirmwareUpgrade struct {

}


type StartFirmwareUpgradeResponse struct {
	UploadUri xsd.AnyURI
	UploadDelay xsd.Duration
	ExpectedDownTime xsd.Duration

}


type StartSystemRestore struct {

}


type StartSystemRestoreResponse struct {
	UploadUri xsd.AnyURI
	ExpectedDownTime xsd.Duration

}


type GetStorageConfigurations struct {

}


type GetStorageConfigurationsResponse struct {
	StorageConfigurations StorageConfiguration

}


type CreateStorageConfiguration struct {
	StorageConfiguration StorageConfigurationData

}


type CreateStorageConfigurationResponse struct {
	Token onvif.ReferenceToken

}


type GetStorageConfiguration struct {
	Token onvif.ReferenceToken

}


type GetStorageConfigurationResponse struct {
	StorageConfiguration StorageConfiguration

}


type SetStorageConfiguration struct {
	StorageConfiguration StorageConfiguration

}


type SetStorageConfigurationResponse struct {

}


type DeleteStorageConfiguration struct {
	Token onvif.ReferenceToken

}


type DeleteStorageConfigurationResponse struct {

}


type GetGeoLocation struct {

}


type GetGeoLocationResponse struct {
	Location onvif.LocationEntity

}


type SetGeoLocation struct {
	Location onvif.LocationEntity

}


type SetGeoLocationResponse struct {

}


type DeleteGeoLocation struct {
	Location onvif.LocationEntity

}


type DeleteGeoLocationResponse struct {

}