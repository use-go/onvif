package onvif

// Onvif WebService
const (
	CoreWebService      = "Core"
	DeviceWebService    = "Device"
	EventWebService     = "Event"
	MediaWebService     = "Media"
	StreamingWebService = "Streaming"
	PTZWebService       = "PTZ"
)

// Onvif WebService function
const (
	// WebService - Device
	GetHostname              = "GetHostname"
	SetHostname              = "SetHostname"
	GetDNS                   = "GetDNS"
	SetDNS                   = "SetDNS"
	GetNetworkInterfaces     = "GetNetworkInterfaces"
	SetNetworkInterfaces     = "SetNetworkInterfaces"
	GetNetworkProtocols      = "GetNetworkProtocols"
	SetNetworkProtocols      = "SetNetworkProtocols"
	GetNetworkDefaultGateway = "GetNetworkDefaultGateway"
	SetNetworkDefaultGateway = "SetNetworkDefaultGateway"

	GetDeviceInformation    = "GetDeviceInformation"
	GetSystemDateAndTime    = "GetSystemDateAndTime"
	SetSystemDateAndTime    = "SetSystemDateAndTime"
	SetSystemFactoryDefault = "SetSystemFactoryDefault"
	SystemReboot            = "SystemReboot"

	// WebService - Media
	GetMetadataConfiguration            = "GetMetadataConfiguration"
	GetMetadataConfigurations           = "GetMetadataConfigurations"
	AddMetadataConfiguration            = "AddMetadataConfiguration"
	RemoveMetadataConfiguration         = "RemoveMetadataConfiguration"
	SetMetadataConfiguration            = "SetMetadataConfiguration"
	GetCompatibleMetadataConfigurations = "GetCompatibleMetadataConfigurations"
	GetMetadataConfigurationOptions     = "GetMetadataConfigurationOptions"
)
