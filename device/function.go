package device

type GetHostnameFunction struct{}

func (function *GetHostnameFunction) Request() interface{} {
	return &GetHostname{}
}

func (function *GetHostnameFunction) Response() interface{} {
	return &GetHostnameResponse{}
}

type SetHostnameFunction struct{}

func (function *SetHostnameFunction) Request() interface{} {
	return &SetHostname{}
}

func (function *SetHostnameFunction) Response() interface{} {
	return &SetHostnameResponse{}
}

type GetDNSFunction struct{}

func (function *GetDNSFunction) Request() interface{} {
	return &GetDNS{}
}

func (function *GetDNSFunction) Response() interface{} {
	return &GetDNSResponse{}
}

type SetDNSFunction struct{}

func (function *SetDNSFunction) Request() interface{} {
	return &SetDNS{}
}

func (function *SetDNSFunction) Response() interface{} {
	return &SetDNSResponse{}
}

type GetNetworkInterfacesFunction struct{}

func (function *GetNetworkInterfacesFunction) Request() interface{} {
	return &GetNetworkInterfaces{}
}

func (function *GetNetworkInterfacesFunction) Response() interface{} {
	return &GetNetworkInterfacesResponse{}
}

type SetNetworkInterfacesFunction struct{}

func (function *SetNetworkInterfacesFunction) Request() interface{} {
	return &SetNetworkInterfaces{}
}

func (function *SetNetworkInterfacesFunction) Response() interface{} {
	return &SetNetworkInterfacesResponse{}
}

type GetNetworkProtocolsFunction struct{}

func (function *GetNetworkProtocolsFunction) Request() interface{} {
	return &GetNetworkProtocols{}
}

func (function *GetNetworkProtocolsFunction) Response() interface{} {
	return &GetNetworkProtocolsResponse{}
}

type SetNetworkProtocolsFunction struct{}

func (function *SetNetworkProtocolsFunction) Request() interface{} {
	return &SetNetworkProtocols{}
}

func (function *SetNetworkProtocolsFunction) Response() interface{} {
	return &SetNetworkProtocolsResponse{}
}

type GetNetworkDefaultGatewayFunction struct{}

func (function *GetNetworkDefaultGatewayFunction) Request() interface{} {
	return &GetNetworkDefaultGateway{}
}

func (function *GetNetworkDefaultGatewayFunction) Response() interface{} {
	return &GetNetworkDefaultGatewayResponse{}
}

type SetNetworkDefaultGatewayFunction struct{}

func (function *SetNetworkDefaultGatewayFunction) Request() interface{} {
	return &SetNetworkDefaultGateway{}
}

func (function *SetNetworkDefaultGatewayFunction) Response() interface{} {
	return &SetNetworkDefaultGatewayResponse{}
}

type GetDeviceInformationFunction struct{}

func (function *GetDeviceInformationFunction) Request() interface{} {
	return &GetDeviceInformation{}
}

func (function *GetDeviceInformationFunction) Response() interface{} {
	return &GetDeviceInformationResponse{}
}

type GetSystemDateAndTimeFunction struct{}

func (function *GetSystemDateAndTimeFunction) Request() interface{} {
	return &GetSystemDateAndTime{}
}

func (function *GetSystemDateAndTimeFunction) Response() interface{} {
	return &GetSystemDateAndTimeResponse{}
}

type SetSystemDateAndTimeFunction struct{}

func (function *SetSystemDateAndTimeFunction) Request() interface{} {
	return &SetSystemDateAndTime{}
}

func (function *SetSystemDateAndTimeFunction) Response() interface{} {
	return &SetSystemDateAndTimeResponse{}
}

type SetSystemFactoryDefaultFunction struct{}

func (function *SetSystemFactoryDefaultFunction) Request() interface{} {
	return &SetSystemFactoryDefault{}
}

func (function *SetSystemFactoryDefaultFunction) Response() interface{} {
	return &SetSystemFactoryDefaultResponse{}
}

type SystemRebootFunction struct{}

func (function *SystemRebootFunction) Request() interface{} {
	return &SystemReboot{}
}

func (function *SystemRebootFunction) Response() interface{} {
	return &SystemRebootResponse{}
}
