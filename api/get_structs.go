package api

import (
	"github.com/yakovlevdmv/goonvif/PTZ"
	"errors"
	"github.com/yakovlevdmv/goonvif/Device"
)

func getPTZStructByName(name string) (interface{}, error) {
	switch name {
	case "GetServiceCapabilities":
		return &PTZ.GetServiceCapabilities{}, nil
	case "GetNodes":
		return &PTZ.GetNodes{}, nil
	case "GetNode":
		return &PTZ.GetNode{}, nil
	case "GetConfiguration":
		return &PTZ.GetConfiguration{}, nil
	case "GetConfigurations":
		return &PTZ.GetConfigurations{}, nil
	case "SetConfiguration":
		return &PTZ.SetConfiguration{}, nil
	case "GetConfigurationOptions":
		return &PTZ.GetConfigurationOptions{}, nil
	case "SendAuxiliaryCommand":
		return &PTZ.SendAuxiliaryCommand{}, nil
	case "GetPresets":
		return &PTZ.GetPresets{}, nil
	case "SetPreset":
		return &PTZ.SetPreset{}, nil
	case "RemovePreset":
		return &PTZ.RemovePreset{}, nil
	case "GotoPreset":
		return &PTZ.GotoPreset{}, nil
	case "GotoHomePosition":
		return &PTZ.GotoHomePosition{}, nil
	case "SetHomePosition":
		return &PTZ.SetHomePosition{}, nil
	case "ContinuousMove":
		return &PTZ.ContinuousMove{}, nil
	case "RelativeMove":
		return &PTZ.RelativeMove{}, nil
	case "GetStatus":
		return &PTZ.GetStatus{}, nil
	case "AbsoluteMove":
		return &PTZ.AbsoluteMove{}, nil
	case "GeoMove":
		return &PTZ.GeoMove{}, nil
	case "Stop":
		return &PTZ.Stop{}, nil
	case "GetPresetTours":
		return &PTZ.GetPresetTours{}, nil
	case "GetPresetTour":
		return &PTZ.GetPresetTour{}, nil
	case "GetPresetTourOptions":
		return &PTZ.GetPresetTourOptions{}, nil
	case "CreatePresetTour":
		return &PTZ.CreatePresetTour{}, nil
	case "ModifyPresetTour":
		return &PTZ.ModifyPresetTour{}, nil
	case "OperatePresetTour":
		return &PTZ.OperatePresetTour{}, nil
	case "RemovePresetTour":
		return &PTZ.RemovePresetTour{}, nil
	case "GetCompatibleConfigurations":
		return &PTZ.GetCompatibleConfigurations{}, nil
	default:
		return nil, errors.New("there is no such method in the PTZ service")
	}
}

func getDeviceStructByName(name string) (interface{}, error) {
	switch name {
	case "GetServices":
		return &Device.GetServices{}, nil
	case "GetServiceCapabilities":
		return &Device.GetServiceCapabilities{}, nil
	case "GetDeviceInformation":
		return &Device.GetDeviceInformation{}, nil
	case "SetSystemDateAndTime":
		return &Device.SetSystemDateAndTime{}, nil
	case "GetSystemDateAndTime":
		return &Device.GetSystemDateAndTime{}, nil
	case "SetSystemFactoryDefault":
		return &Device.SetSystemFactoryDefault{}, nil
	case "UpgradeSystemFirmware":
		return &Device.UpgradeSystemFirmware{}, nil
	case "SystemReboot":
		return &Device.SystemReboot{}, nil
	case "RestoreSystem":
		return &Device.RestoreSystem{}, nil
	case "GetSystemBackup":
		return &Device.GetSystemBackup{}, nil
	case "GetSystemLog":
		return &Device.GetSystemLog{}, nil
	case "GetSystemSupportInformation":
		return &Device.GetSystemSupportInformation{}, nil
	case "GetScopes":
		return &Device.GetScopes{}, nil
	case "SetScopes":
		return &Device.SetScopes{}, nil
	case "AddScopes":
		return &Device.AddScopes{}, nil
	case "RemoveScopes":
		return &Device.RemoveScopes{}, nil
	case "GetDiscoveryMode":
		return &Device.GetDiscoveryMode{}, nil
	case "SetDiscoveryMode":
		return &Device.SetDiscoveryMode{}, nil
	case "GetRemoteDiscoveryMode":
		return &Device.GetRemoteDiscoveryMode{}, nil
	case "SetRemoteDiscoveryMode":
		return &Device.SetRemoteDiscoveryMode{}, nil
	case "GetDPAddresses":
		return &Device.GetDPAddresses{}, nil
	case "SetDPAddresses":
		return &Device.SetDPAddresses{}, nil
	case "GetEndpointReference":
		return &Device.GetEndpointReference{}, nil
	case "GetRemoteUser":
		return &Device.GetRemoteUser{}, nil
	case "SetRemoteUser":
		return &Device.SetRemoteUser{}, nil
	case "GetUsers":
		return &Device.GetUsers{}, nil
	case "CreateUsers":
		return &Device.CreateUsers{}, nil
	case "DeleteUsers":
		return &Device.DeleteUsers{}, nil
	case "SetUser":
		return &Device.SetUser{}, nil
	case "GetWsdlUrl":
		return &Device.GetWsdlUrl{}, nil
	case "GetCapabilities":
		return &Device.GetCapabilities{}, nil
	case "GetHostname":
		return &Device.GetHostname{}, nil
	case "SetHostname":
		return &Device.SetHostname{}, nil
	case "SetHostnameFromDHCP":
		return &Device.SetHostnameFromDHCP{}, nil
	case "GetDNS":
		return &Device.GetDNS{}, nil
	case "SetDNS":
		return &Device.SetDNS{}, nil
	case "GetNTP":
		return &Device.GetNTP{}, nil
	case "SetNTP":
		return &Device.SetNTP{}, nil
	case "GetDynamicDNS":
		return &Device.GetDynamicDNS{}, nil
	case "SetDynamicDNS":
		return &Device.SetDynamicDNS{}, nil
	case "GetNetworkInterfaces":
		return &Device.GetNetworkInterfaces{}, nil
	case "SetNetworkInterfaces":
		return &Device.SetNetworkInterfaces{}, nil
	case "GetNetworkProtocols":
		return &Device.GetNetworkProtocols{}, nil
	case "SetNetworkProtocols":
		return &Device.SetNetworkProtocols{}, nil
	case "GetNetworkDefaultGateway":
		return &Device.GetNetworkDefaultGateway{}, nil
	case "SetNetworkDefaultGateway":
		return &Device.SetNetworkDefaultGateway{}, nil
	case "GetZeroConfiguration":
		return &Device.GetZeroConfiguration{}, nil
	case "SetZeroConfiguration":
		return &Device.SetZeroConfiguration{}, nil
	case "GetIPAddressFilter":
		return &Device.GetIPAddressFilter{}, nil
	case "SetIPAddressFilter":
		return &Device.SetIPAddressFilter{}, nil
	case "AddIPAddressFilter":
		return &Device.AddIPAddressFilter{}, nil
	case "RemoveIPAddressFilter":
		return &Device.RemoveIPAddressFilter{}, nil
	case "GetAccessPolicy":
		return &Device.GetAccessPolicy{}, nil
	case "SetAccessPolicy":
		return &Device.SetAccessPolicy{}, nil
	case "CreateCertificate":
		return &Device.CreateCertificate{}, nil
	case "GetCertificates":
		return &Device.GetCertificates{}, nil
	case "GetCertificatesStatus":
		return &Device.GetCertificatesStatus{}, nil
	case "SetCertificatesStatus":
		return &Device.SetCertificatesStatus{}, nil
	case "DeleteCertificates":
		return &Device.DeleteCertificates{}, nil
	case "GetPkcs10Request":
		return &Device.GetPkcs10Request{}, nil
	case "LoadCertificates":
		return &Device.LoadCertificates{}, nil
	case "GetClientCertificateMode":
		return &Device.GetClientCertificateMode{}, nil
	case "SetClientCertificateMode":
		return &Device.SetClientCertificateMode{}, nil
	case "GetRelayOutputs":
		return &Device.GetRelayOutputs{}, nil
	case "SetRelayOutputSettings":
		return &Device.SetRelayOutputSettings{}, nil
	case "SetRelayOutputState":
		return &Device.SetRelayOutputState{}, nil
	case "SendAuxiliaryCommand":
		return &Device.SendAuxiliaryCommand{}, nil
	case "GetCACertificates":
		return &Device.GetCACertificates{}, nil
	case "LoadCertificateWithPrivateKey":
		return &Device.LoadCertificateWithPrivateKey{}, nil
	case "GetCertificateInformation":
		return &Device.GetCertificateInformation{}, nil
	case "LoadCACertificates":
		return &Device.LoadCACertificates{}, nil
	case "CreateDot1XConfiguration":
		return &Device.CreateDot1XConfiguration{}, nil
	case "SetDot1XConfiguration":
		return &Device.SetDot1XConfiguration{}, nil
	case "GetDot1XConfiguration":
		return &Device.GetDot1XConfiguration{}, nil
	case "GetDot1XConfigurations":
		return &Device.GetDot1XConfigurations{}, nil
	case "DeleteDot1XConfiguration":
		return &Device.DeleteDot1XConfiguration{}, nil
	case "GetDot11Capabilities":
		return &Device.GetDot11Capabilities{}, nil
	case "GetDot11Status":
		return &Device.GetDot11Status{}, nil
	case "ScanAvailableDot11Networks":
		return &Device.ScanAvailableDot11Networks{}, nil
	case "GetSystemUris":
		return &Device.GetSystemUris{}, nil
	case "StartFirmwareUpgrade":
		return &Device.StartFirmwareUpgrade{}, nil
	case "StartSystemRestore":
		return &Device.StartSystemRestore{}, nil
	case "GetStorageConfigurations":
		return &Device.GetStorageConfigurations{}, nil
	case "CreateStorageConfiguration":
		return &Device.CreateStorageConfiguration{}, nil
	case "GetStorageConfiguration":
		return &Device.GetStorageConfiguration{}, nil
	case "SetStorageConfiguration":
		return &Device.SetStorageConfiguration{}, nil
	case "DeleteStorageConfiguration":
		return &Device.DeleteStorageConfiguration{}, nil
	case "GetGeoLocation":
		return &Device.GetGeoLocation{}, nil
	case "SetGeoLocation":
		return &Device.SetGeoLocation{}, nil
	case "DeleteGeoLocation":
		return &Device.DeleteGeoLocation{}, nil

	default:
		return nil, errors.New("there is no such method in the Device service")
	}
}