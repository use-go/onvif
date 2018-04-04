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
	Data StorageConfigurationData `xml:"wsdl:Data"`
}

type StorageConfigurationData struct {
	Type xsd.String `xml:"type,attr"`
	LocalPath xsd.AnyURI `xml:"wsdl:LocalPath"`
	StorageUri xsd.AnyURI `xml:"wsdl:StorageUri"`
	User UserCredential `xml:"wsdl:User"`
	Extension xsd.AnyURI `xml:"wsdl:Extension"`
}

type UserCredential struct {
	UserName xsd.String `xml:"wsdl:UserName"`
	Password xsd.String `xml:"wsdl:Password"`
	Extension xsd.AnyType `xml:"wsdl:Extension"`
}

//Device main types

type GetServices struct {
	XMLName string `xml:"wsdl:GetServices"`
	IncludeCapability xsd.Boolean `xml:"wsdl:IncludeCapability"`

}


type GetServicesResponse struct {
	Service Service

}


type GetServiceCapabilities struct {
	XMLName string `xml:"wsdl:GetServiceCapabilities"`

}


type GetServiceCapabilitiesResponse struct {
	Capabilities DeviceServiceCapabilities

}

type GetDeviceInformation struct {
	XMLName string `xml:"wsdl:GetDeviceInformation"`
}


type GetDeviceInformationResponse struct {
	Manufacturer string
	Model string
	FirmwareVersion string
	SerialNumber string
	HardwareId string

}


type SetSystemDateAndTime struct {
	XMLName string `xml:"wsdl:SetSystemDateAndTime"`
	DateTimeType    onvif.SetDateTimeType `xml:"wsdl:DateTimeType"`
	DaylightSavings xsd.Boolean `xml:"wsdl:DaylightSavings"`
	TimeZone        onvif.TimeZone `xml:"wsdl:TimeZone"`
	UTCDateTime     onvif.DateTime `xml:"wsdl:UTCDateTime"`
}


type SetSystemDateAndTimeResponse struct {

}


type GetSystemDateAndTime struct {
	XMLName string `xml:"wsdl:GetSystemDateAndTime"`

}


type GetSystemDateAndTimeResponse struct {
	SystemDateAndTime onvif.SystemDateTime

}


type SetSystemFactoryDefault struct {
	XMLName string `xml:"wsdl:SetSystemFactoryDefault"`
	FactoryDefault onvif.FactoryDefaultType `xml:"wsdl:FactoryDefault"`

}


type SetSystemFactoryDefaultResponse struct {

}


type UpgradeSystemFirmware struct {
	XMLName string `xml:"wsdl:UpgradeSystemFirmware"`
	Firmware onvif.AttachmentData `xml:"wsdl:Firmware"`

}


type UpgradeSystemFirmwareResponse struct {
	Message string

}


type SystemReboot struct {
	XMLName string `xml:"wsdl:SystemReboot"`

}


type SystemRebootResponse struct {
	Message string

}

//TODO: one or more repetitions
type RestoreSystem struct {
	XMLName string `xml:"wsdl:RestoreSystem"`
	BackupFiles onvif.BackupFile `xml:"wsdl:BackupFiles"`

}


type RestoreSystemResponse struct {

}


type GetSystemBackup struct {
	XMLName string `xml:"wsdl:GetSystemBackup"`

}


type GetSystemBackupResponse struct {
	BackupFiles onvif.BackupFile

}


type GetSystemLog struct {
	XMLName string `xml:"wsdl:GetSystemLog"`
	LogType onvif.SystemLogType `xml:"wsdl:LogType"`

}


type GetSystemLogResponse struct {
	SystemLog onvif.SystemLog

}


type GetSystemSupportInformation struct {
	XMLName string `xml:"wsdl:GetSystemSupportInformation"`

}


type GetSystemSupportInformationResponse struct {
	SupportInformation onvif.SupportInformation

}


type GetScopes struct {
	XMLName string `xml:"wsdl:GetScopes"`

}


type GetScopesResponse struct {
	Scopes onvif.Scope

}

//TODO: one or more scopes
type SetScopes struct {
	XMLName string `xml:"wsdl:SetScopes"`
	Scopes xsd.AnyURI `xml:"wsdl:Scopes"`

}


type SetScopesResponse struct {

}

//TODO: list of scopes
type AddScopes struct {
	XMLName string `xml:"wsdl:AddScopes"`
	ScopeItem xsd.AnyURI`xml:"wsdl:ScopeItem"`

}


type AddScopesResponse struct {

}

//TODO: One or more repetitions
type RemoveScopes struct {
	XMLName string `xml:"wsdl:RemoveScopes"`
	ScopeItem xsd.AnyURI `xml:"onvif:ScopeItem"`

}


type RemoveScopesResponse struct {
	ScopeItem xsd.AnyURI

}


type GetDiscoveryMode struct {
	XMLName string `xml:"wsdl:GetDiscoveryMode"`

}


type GetDiscoveryModeResponse struct {
	DiscoveryMode onvif.DiscoveryMode

}


type SetDiscoveryMode struct {
	XMLName string `xml:"wsdl:SetDiscoveryMode"`
	DiscoveryMode onvif.DiscoveryMode `xml:"wsdl:DiscoveryMode"`

}


type SetDiscoveryModeResponse struct {

}


type GetRemoteDiscoveryMode struct {
	XMLName string `xml:"wsdl:GetRemoteDiscoveryMode"`

}


type GetRemoteDiscoveryModeResponse struct {
	RemoteDiscoveryMode onvif.DiscoveryMode

}


type SetRemoteDiscoveryMode struct {
	XMLName string `xml:"wsdl:SetRemoteDiscoveryMode"`
	RemoteDiscoveryMode onvif.DiscoveryMode `xml:"wsdl:RemoteDiscoveryMode"`

}


type SetRemoteDiscoveryModeResponse struct {

}


type GetDPAddresses struct {
	XMLName string `xml:"wsdl:GetDPAddresses"`

}


type GetDPAddressesResponse struct {
	DPAddress onvif.NetworkHost

}


type SetDPAddresses struct {
	XMLName string `xml:"wsdl:SetDPAddresses"`
	DPAddress onvif.NetworkHost `xml:"wsdl:DPAddress"`

}


type SetDPAddressesResponse struct {

}


type GetEndpointReference struct {
	XMLName string `xml:"wsdl:GetEndpointReference"`

}


type GetEndpointReferenceResponse struct {
	GUID string

}


type GetRemoteUser struct {
	XMLName string `xml:"wsdl:GetRemoteUser"`

}


type GetRemoteUserResponse struct {
	RemoteUser onvif.RemoteUser

}


type SetRemoteUser struct {
	XMLName string `xml:"wsdl:SetRemoteUser"`
	RemoteUser onvif.RemoteUser `xml:"wsdl:RemoteUser"`

}


type SetRemoteUserResponse struct {

}


type GetUsers struct {
	XMLName string `xml:"wsdl:GetUsers"`

}


type GetUsersResponse struct {
	User onvif.User

}

//TODO: List of users
type CreateUsers struct {
	XMLName string `xml:"wsdl:CreateUsers"`
	User onvif.User `xml:"wsdl:User,omitempty"`

}


type CreateUsersResponse struct {

}

//TODO: one or more Username
type DeleteUsers struct {
	XMLName xsd.String `xml:"wsdl:DeleteUsers"`
	Username xsd.String `xml:"wsdl:Username"`

}


type DeleteUsersResponse struct {

}


type SetUser struct {
	XMLName string `xml:"wsdl:SetUser"`
	User onvif.User `xml:"wsdl:User"`

}


type SetUserResponse struct {

}


type GetWsdlUrl struct {
	XMLName string `xml:"wsdl:GetWsdlUrl"`

}


type GetWsdlUrlResponse struct {
	WsdlUrl xsd.AnyURI

}


type GetCapabilities struct {
	XMLName string `xml:"wsdl:GetCapabilities"`
	Category onvif.CapabilityCategory `xml:"wsdl:Category"`

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
	XMLName string `xml:"wsdl:SetHostname"`
	Name xsd.Token `xml:"wsdl:Name"`

}


type SetHostnameResponse struct {

}


type SetHostnameFromDHCP struct {
	XMLName string `xml:"wsdl:SetHostnameFromDHCP"`
	FromDHCP xsd.Boolean `xml:"wsdl:FromDHCP"`

}


type SetHostnameFromDHCPResponse struct {
	RebootNeeded xsd.Boolean

}


type GetDNS struct {
	XMLName string `xml:"wsdl:GetDNS"`

}


type GetDNSResponse struct {
	DNSInformation onvif.DNSInformation

}


type SetDNS struct {
	XMLName string `xml:"wsdl:SetDNS"`
	FromDHCP     xsd.Boolean `xml:"wsdl:FromDHCP"`
	SearchDomain xsd.Token `xml:"wsdl:SearchDomain"`
	DNSManual    onvif.IPAddress `xml:"wsdl:DNSManual"`

}


type SetDNSResponse struct {

}


type GetNTP struct {
	XMLName string `xml:"wsdl:GetNTP"`

}


type GetNTPResponse struct {
	NTPInformation onvif.NTPInformation

}


type SetNTP struct {
	XMLName string `xml:"wsdl:SetNTP"`
	FromDHCP  xsd.Boolean `xml:"wsdl:FromDHCP"`
	NTPManual onvif.NetworkHost `xml:"wsdl:NTPManual"`

}


type SetNTPResponse struct {

}


type GetDynamicDNS struct {
	XMLName string `xml:"wsdl:GetDynamicDNS"`

}


type GetDynamicDNSResponse struct {
	DynamicDNSInformation onvif.DynamicDNSInformation

}


type SetDynamicDNS struct {
	XMLName string `xml:"wsdl:SetDynamicDNS"`
	Type onvif.DynamicDNSType `xml:"wsdl:Type"`
	Name onvif.DNSName `xml:"wsdl:Name"`
	TTL  xsd.Duration `xml:"wsdl:TTL"`

}


type SetDynamicDNSResponse struct {

}


type GetNetworkInterfaces struct {
	XMLName string `xml:"wsdl:GetNetworkInterfaces"`

}


type GetNetworkInterfacesResponse struct {
	NetworkInterfaces onvif.NetworkInterface

}


type SetNetworkInterfaces struct {
	XMLName string `xml:"wsdl:SetNetworkInterfaces"`
	InterfaceToken onvif.ReferenceToken `xml:"wsdl:InterfaceToken"`
	NetworkInterface onvif.NetworkInterfaceSetConfiguration `xml:"wsdl:NetworkInterface"`

}


type SetNetworkInterfacesResponse struct {
	RebootNeeded xsd.Boolean

}


type GetNetworkProtocols struct {
	XMLName string `xml:"wsdl:GetNetworkProtocols"`

}


type GetNetworkProtocolsResponse struct {
	NetworkProtocols onvif.NetworkProtocol

}


type SetNetworkProtocols struct {
	XMLName string `xml:"wsdl:SetNetworkProtocols"`
	NetworkProtocols onvif.NetworkProtocol `xml:"wsdl:NetworkProtocols"`

}


type SetNetworkProtocolsResponse struct {

}


type GetNetworkDefaultGateway struct {
	XMLName string `xml:"wsdl:GetNetworkDefaultGateway"`

}


type GetNetworkDefaultGatewayResponse struct {
	NetworkGateway onvif.NetworkGateway

}


type SetNetworkDefaultGateway struct {
	XMLName string `xml:"wsdl:SetNetworkDefaultGateway"`
	IPv4Address onvif.IPv4Address `xml:"wsdl:IPv4Address"`
	IPv6Address onvif.IPv6Address `xml:"wsdl:IPv6Address"`

}


type SetNetworkDefaultGatewayResponse struct {

}


type GetZeroConfiguration struct {
	XMLName string `xml:"wsdl:GetZeroConfiguration"`

}


type GetZeroConfigurationResponse struct {
	ZeroConfiguration onvif.NetworkZeroConfiguration

}


type SetZeroConfiguration struct {
	XMLName string `xml:"wsdl:SetZeroConfiguration"`
	InterfaceToken onvif.ReferenceToken `xml:"wsdl:InterfaceToken"`
	Enabled        xsd.Boolean `xml:"wsdl:Enabled"`

}


type SetZeroConfigurationResponse struct {

}


type GetIPAddressFilter struct {
	XMLName string `xml:"wsdl:GetIPAddressFilter"`

}


type GetIPAddressFilterResponse struct {
	IPAddressFilter onvif.IPAddressFilter

}


type SetIPAddressFilter struct {
	XMLName string `xml:"wsdl:SetIPAddressFilter"`
	IPAddressFilter onvif.IPAddressFilter `xml:"wsdl:IPAddressFilter"`

}


type SetIPAddressFilterResponse struct {

}

//This operation adds an IP filter address to a device.
//If the device supports device access control based on
//IP filtering rules (denied or accepted ranges of IP addresses),
//the device shall support adding of IP filtering addresses through
//the AddIPAddressFilter command.
type AddIPAddressFilter struct {
	XMLName string `xml:"wsdl:AddIPAddressFilter"`
	IPAddressFilter onvif.IPAddressFilter `xml:"wsdl:IPAddressFilter"`

}


type AddIPAddressFilterResponse struct {

}


type RemoveIPAddressFilter struct {
	XMLName string `xml:"wsdl:RemoveIPAddressFilter"`
	IPAddressFilter onvif.IPAddressFilter `xml:"onvif:IPAddressFilter"`

}


type RemoveIPAddressFilterResponse struct {

}


type GetAccessPolicy struct {
	XMLName string `xml:"wsdl:GetAccessPolicy"`
}


type GetAccessPolicyResponse struct {
	PolicyFile onvif.BinaryData

}


type SetAccessPolicy struct {
	XMLName string `xml:"wsdl:SetAccessPolicy"`
	PolicyFile onvif.BinaryData `xml:"wsdl:PolicyFile"`

}


type SetAccessPolicyResponse struct {

}


type CreateCertificate struct {
	XMLName string `xml:"wsdl:CreateCertificate"`
	CertificateID  xsd.Token `xml:"wsdl:CertificateID,omitempty"`
	Subject        string `xml:"wsdl:Subject,omitempty"`
	ValidNotBefore xsd.DateTime `xml:"wsdl:ValidNotBefore,omitempty"`
	ValidNotAfter  xsd.DateTime `xml:"wsdl:ValidNotAfter,omitempty"`

}


type CreateCertificateResponse struct {
	NvtCertificate onvif.Certificate

}


type GetCertificates struct {
	XMLName string `xml:"wsdl:GetCertificates"`
}


type GetCertificatesResponse struct {
	NvtCertificate onvif.Certificate

}


type GetCertificatesStatus struct {
	XMLName string `xml:"wsdl:GetCertificatesStatus"`

}


type GetCertificatesStatusResponse struct {
	CertificateStatus onvif.CertificateStatus

}


type SetCertificatesStatus struct {
	XMLName string `xml:"wsdl:SetCertificatesStatus"`
	CertificateStatus onvif.CertificateStatus `xml:"wsdl:CertificateStatus"`

}


type SetCertificatesStatusResponse struct {

}

//TODO: List of CertificateID
type DeleteCertificates struct {
	XMLName string `xml:"wsdl:DeleteCertificates"`
	CertificateID xsd.Token `xml:"wsdl:CertificateID"`

}


type DeleteCertificatesResponse struct {

}

//TODO: Откуда onvif:data = cid:21312413412
type GetPkcs10Request struct {
	XMLName string `xml:"wsdl:GetPkcs10Request"`
	CertificateID xsd.Token `xml:"wsdl:CertificateID"`
	Subject xsd.String `xml:"wsdl:Subject"`
	Attributes onvif.BinaryData `xml:"wsdl:Attributes"`

}


type GetPkcs10RequestResponse struct {
	Pkcs10Request onvif.BinaryData

}

//TODO: one or more NTVCertificate
type LoadCertificates struct {
	XMLName string `xml:"wsdl:LoadCertificates"`
	NVTCertificate onvif.Certificate `xml:"wsdl:NVTCertificate"`

}


type LoadCertificatesResponse struct {

}


type GetClientCertificateMode struct {
	XMLName string `xml:"wsdl:GetClientCertificateMode"`

}


type GetClientCertificateModeResponse struct {
	Enabled xsd.Boolean

}


type SetClientCertificateMode struct {
	XMLName string `xml:"wsdl:SetClientCertificateMode"`
	Enabled xsd.Boolean `xml:"wsdl:Enabled"`

}


type SetClientCertificateModeResponse struct {

}


type GetRelayOutputs struct {
	XMLName string `xml:"wsdl:GetRelayOutputs"`

}


type GetRelayOutputsResponse struct {
	RelayOutputs onvif.RelayOutput

}


type SetRelayOutputSettings struct {
	XMLName string `xml:"wsdl:SetRelayOutputSettings"`
	RelayOutputToken onvif.ReferenceToken `xml:"wsdl:RelayOutputToken"`
	Properties onvif.RelayOutputSettings `xml:"wsdl:Properties"`

}


type SetRelayOutputSettingsResponse struct {

}


type SetRelayOutputState struct {
	XMLName string `xml:"wsdl:SetRelayOutputState"`
	RelayOutputToken onvif.ReferenceToken `xml:"wsdl:RelayOutputToken"`
	LogicalState onvif.RelayLogicalState `xml:"wsdl:LogicalState"`

}


type SetRelayOutputStateResponse struct {

}


type SendAuxiliaryCommand struct {
	XMLName string `xml:"wsdl:SendAuxiliaryCommand"`
	AuxiliaryCommand onvif.AuxiliaryData `xml:"wsdl:AuxiliaryCommand"`

}


type SendAuxiliaryCommandResponse struct {
	AuxiliaryCommandResponse onvif.AuxiliaryData

}


type GetCACertificates struct {
	XMLName string `xml:"wsdl:GetCACertificates"`
}


type GetCACertificatesResponse struct {
	CACertificate onvif.Certificate

}

//TODO: one or more CertificateWithPrivateKey
type LoadCertificateWithPrivateKey struct {
	XMLName string `xml:"wsdl:LoadCertificateWithPrivateKey"`
	CertificateWithPrivateKey onvif.CertificateWithPrivateKey `xml:"wsdl:CertificateWithPrivateKey"`

}


type LoadCertificateWithPrivateKeyResponse struct {

}


type GetCertificateInformation struct {
	XMLName string `xml:"wsdl:GetCertificateInformation"`
	CertificateID xsd.Token `xml:"wsdl:CertificateID"`

}


type GetCertificateInformationResponse struct {
	CertificateInformation onvif.CertificateInformation

}


type LoadCACertificates struct {
	XMLName string `xml:"wsdl:LoadCACertificates"`
	CACertificate onvif.Certificate `xml:"wsdl:CACertificate"`

}


type LoadCACertificatesResponse struct {

}


type CreateDot1XConfiguration struct {
	XMLName string `xml:"wsdl:CreateDot1XConfiguration"`
	Dot1XConfiguration onvif.Dot1XConfiguration `xml:"wsdl:Dot1XConfiguration"`

}


type CreateDot1XConfigurationResponse struct {

}


type SetDot1XConfiguration struct {
	XMLName string `xml:"wsdl:SetDot1XConfiguration"`
	Dot1XConfiguration onvif.Dot1XConfiguration `xml:"wsdl:Dot1XConfiguration"`

}


type SetDot1XConfigurationResponse struct {

}


type GetDot1XConfiguration struct {
	XMLName string `xml:"wsdl:GetDot1XConfiguration"`
	Dot1XConfigurationToken onvif.ReferenceToken `xml:"wsdl:Dot1XConfigurationToken"`

}


type GetDot1XConfigurationResponse struct {
	Dot1XConfiguration onvif.Dot1XConfiguration

}


type GetDot1XConfigurations struct {
	XMLName string `xml:"wsdl:GetDot1XConfigurations"`

}


type GetDot1XConfigurationsResponse struct {
	Dot1XConfiguration onvif.Dot1XConfiguration

}

//TODO: Zero or more Dot1XConfigurationToken
type DeleteDot1XConfiguration struct {
	XMLName string `xml:"wsdl:DeleteDot1XConfiguration"`
	Dot1XConfigurationToken onvif.ReferenceToken `xml:"wsdl:Dot1XConfigurationToken"`

}


type DeleteDot1XConfigurationResponse struct {

}


type GetDot11Capabilities struct {
	XMLName string `xml:"wsdl:GetDot11Capabilities"`

}


type GetDot11CapabilitiesResponse struct {
	Capabilities onvif.Dot11Capabilities

}


type GetDot11Status struct {
	XMLName string `xml:"wsdl:GetDot11Status"`
	InterfaceToken onvif.ReferenceToken `xml:"wsdl:InterfaceToken"`

}


type GetDot11StatusResponse struct {
	Status onvif.Dot11Status

}


type ScanAvailableDot11Networks struct {
	XMLName string `xml:"wsdl:ScanAvailableDot11Networks"`
	InterfaceToken onvif.ReferenceToken `xml:"wsdl:InterfaceToken"`

}


type ScanAvailableDot11NetworksResponse struct {
	Networks onvif.Dot11AvailableNetworks

}


type GetSystemUris struct {
	XMLName string `xml:"wsdl:GetSystemUris"`

}


type GetSystemUrisResponse struct {
	SystemLogUris onvif.SystemLogUriList
	SupportInfoUri xsd.AnyURI
	SystemBackupUri xsd.AnyURI
	Extension xsd.AnyType
}


type StartFirmwareUpgrade struct {
	XMLName string `xml:"wsdl:StartFirmwareUpgrade"`

}


type StartFirmwareUpgradeResponse struct {
	UploadUri xsd.AnyURI
	UploadDelay xsd.Duration
	ExpectedDownTime xsd.Duration

}


type StartSystemRestore struct {
	XMLName string `xml:"wsdl:StartSystemRestore"`

}


type StartSystemRestoreResponse struct {
	UploadUri xsd.AnyURI
	ExpectedDownTime xsd.Duration

}


type GetStorageConfigurations struct {
	XMLName string `xml:"wsdl:GetStorageConfigurations"`

}


type GetStorageConfigurationsResponse struct {
	StorageConfigurations StorageConfiguration

}


type CreateStorageConfiguration struct {
	XMLName string `xml:"wsdl:CreateStorageConfiguration"`
	StorageConfiguration StorageConfigurationData

}


type CreateStorageConfigurationResponse struct {
	Token onvif.ReferenceToken

}


type GetStorageConfiguration struct {
	XMLName string `xml:"wsdl:GetStorageConfiguration"`
	Token onvif.ReferenceToken `xml:"wsdl:Token"`

}


type GetStorageConfigurationResponse struct {
	StorageConfiguration StorageConfiguration

}


type SetStorageConfiguration struct {
	XMLName string `xml:"wsdl:SetStorageConfiguration"`
	StorageConfiguration StorageConfiguration `xml:"wsdl:StorageConfiguration"`

}


type SetStorageConfigurationResponse struct {

}


type DeleteStorageConfiguration struct {
	XMLName string `xml:"wsdl:DeleteStorageConfiguration"`
	Token onvif.ReferenceToken `xml:"wsdl:Token"`

}


type DeleteStorageConfigurationResponse struct {

}


type GetGeoLocation struct {
	XMLName string `xml:"wsdl:GetGeoLocation"`

}


type GetGeoLocationResponse struct {
	Location onvif.LocationEntity

}

//TODO: one or more Location
type SetGeoLocation struct {
	XMLName string `xml:"wsdl:SetGeoLocation"`
	Location onvif.LocationEntity `xml:"wsdl:Location"`

}


type SetGeoLocationResponse struct {

}


type DeleteGeoLocation struct {
	XMLName string `xml:"wsdl:DeleteGeoLocation"`
	Location onvif.LocationEntity `xml:"wsdl:Location"`

}


type DeleteGeoLocationResponse struct {

}