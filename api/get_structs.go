package api

import (
	"errors"

	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
	"github.com/use-go/onvif/ptz"
)

func getPTZStructByName(name string) (interface{}, error) {
	switch name {
	case "GetServiceCapabilities":
		return &ptz.GetServiceCapabilities{}, nil
	case "GetNodes":
		return &ptz.GetNodes{}, nil
	case "GetNode":
		return &ptz.GetNode{}, nil
	case "GetConfiguration":
		return &ptz.GetConfiguration{}, nil
	case "GetConfigurations":
		return &ptz.GetConfigurations{}, nil
	case "SetConfiguration":
		return &ptz.SetConfiguration{}, nil
	case "GetConfigurationOptions":
		return &ptz.GetConfigurationOptions{}, nil
	case "SendAuxiliaryCommand":
		return &ptz.SendAuxiliaryCommand{}, nil
	case "GetPresets":
		return &ptz.GetPresets{}, nil
	case "SetPreset":
		return &ptz.SetPreset{}, nil
	case "RemovePreset":
		return &ptz.RemovePreset{}, nil
	case "GotoPreset":
		return &ptz.GotoPreset{}, nil
	case "GotoHomePosition":
		return &ptz.GotoHomePosition{}, nil
	case "SetHomePosition":
		return &ptz.SetHomePosition{}, nil
	case "ContinuousMove":
		return &ptz.ContinuousMove{}, nil
	case "RelativeMove":
		return &ptz.RelativeMove{}, nil
	case "GetStatus":
		return &ptz.GetStatus{}, nil
	case "AbsoluteMove":
		return &ptz.AbsoluteMove{}, nil
	case "GeoMove":
		return &ptz.GeoMove{}, nil
	case "Stop":
		return &ptz.Stop{}, nil
	case "GetPresetTours":
		return &ptz.GetPresetTours{}, nil
	case "GetPresetTour":
		return &ptz.GetPresetTour{}, nil
	case "GetPresetTourOptions":
		return &ptz.GetPresetTourOptions{}, nil
	case "CreatePresetTour":
		return &ptz.CreatePresetTour{}, nil
	case "ModifyPresetTour":
		return &ptz.ModifyPresetTour{}, nil
	case "OperatePresetTour":
		return &ptz.OperatePresetTour{}, nil
	case "RemovePresetTour":
		return &ptz.RemovePresetTour{}, nil
	case "GetCompatibleConfigurations":
		return &ptz.GetCompatibleConfigurations{}, nil
	default:
		return nil, errors.New("there is no such method in the PTZ service")
	}
}

func getDeviceStructByName(name string) (interface{}, error) {
	switch name {
	case "GetServices":
		return &device.GetServices{}, nil
	case "GetServiceCapabilities":
		return &device.GetServiceCapabilities{}, nil
	case "GetDeviceInformation":
		return &device.GetDeviceInformation{}, nil
	case "SetSystemDateAndTime":
		return &device.SetSystemDateAndTime{}, nil
	case "GetSystemDateAndTime":
		return &device.GetSystemDateAndTime{}, nil
	case "SetSystemFactoryDefault":
		return &device.SetSystemFactoryDefault{}, nil
	case "UpgradeSystemFirmware":
		return &device.UpgradeSystemFirmware{}, nil
	case "SystemReboot":
		return &device.SystemReboot{}, nil
	case "RestoreSystem":
		return &device.RestoreSystem{}, nil
	case "GetSystemBackup":
		return &device.GetSystemBackup{}, nil
	case "GetSystemLog":
		return &device.GetSystemLog{}, nil
	case "GetSystemSupportInformation":
		return &device.GetSystemSupportInformation{}, nil
	case "GetScopes":
		return &device.GetScopes{}, nil
	case "SetScopes":
		return &device.SetScopes{}, nil
	case "AddScopes":
		return &device.AddScopes{}, nil
	case "RemoveScopes":
		return &device.RemoveScopes{}, nil
	case "GetDiscoveryMode":
		return &device.GetDiscoveryMode{}, nil
	case "SetDiscoveryMode":
		return &device.SetDiscoveryMode{}, nil
	case "GetRemoteDiscoveryMode":
		return &device.GetRemoteDiscoveryMode{}, nil
	case "SetRemoteDiscoveryMode":
		return &device.SetRemoteDiscoveryMode{}, nil
	case "GetDPAddresses":
		return &device.GetDPAddresses{}, nil
	case "SetDPAddresses":
		return &device.SetDPAddresses{}, nil
	case "GetEndpointReference":
		return &device.GetEndpointReference{}, nil
	case "GetRemoteUser":
		return &device.GetRemoteUser{}, nil
	case "SetRemoteUser":
		return &device.SetRemoteUser{}, nil
	case "GetUsers":
		return &device.GetUsers{}, nil
	case "CreateUsers":
		return &device.CreateUsers{}, nil
	case "DeleteUsers":
		return &device.DeleteUsers{}, nil
	case "SetUser":
		return &device.SetUser{}, nil
	case "GetWsdlUrl":
		return &device.GetWsdlUrl{}, nil
	case "GetCapabilities":
		return &device.GetCapabilities{}, nil
	case "GetHostname":
		return &device.GetHostname{}, nil
	case "SetHostname":
		return &device.SetHostname{}, nil
	case "SetHostnameFromDHCP":
		return &device.SetHostnameFromDHCP{}, nil
	case "GetDNS":
		return &device.GetDNS{}, nil
	case "SetDNS":
		return &device.SetDNS{}, nil
	case "GetNTP":
		return &device.GetNTP{}, nil
	case "SetNTP":
		return &device.SetNTP{}, nil
	case "GetDynamicDNS":
		return &device.GetDynamicDNS{}, nil
	case "SetDynamicDNS":
		return &device.SetDynamicDNS{}, nil
	case "GetNetworkInterfaces":
		return &device.GetNetworkInterfaces{}, nil
	case "SetNetworkInterfaces":
		return &device.SetNetworkInterfaces{}, nil
	case "GetNetworkProtocols":
		return &device.GetNetworkProtocols{}, nil
	case "SetNetworkProtocols":
		return &device.SetNetworkProtocols{}, nil
	case "GetNetworkDefaultGateway":
		return &device.GetNetworkDefaultGateway{}, nil
	case "SetNetworkDefaultGateway":
		return &device.SetNetworkDefaultGateway{}, nil
	case "GetZeroConfiguration":
		return &device.GetZeroConfiguration{}, nil
	case "SetZeroConfiguration":
		return &device.SetZeroConfiguration{}, nil
	case "GetIPAddressFilter":
		return &device.GetIPAddressFilter{}, nil
	case "SetIPAddressFilter":
		return &device.SetIPAddressFilter{}, nil
	case "AddIPAddressFilter":
		return &device.AddIPAddressFilter{}, nil
	case "RemoveIPAddressFilter":
		return &device.RemoveIPAddressFilter{}, nil
	case "GetAccessPolicy":
		return &device.GetAccessPolicy{}, nil
	case "SetAccessPolicy":
		return &device.SetAccessPolicy{}, nil
	case "CreateCertificate":
		return &device.CreateCertificate{}, nil
	case "GetCertificates":
		return &device.GetCertificates{}, nil
	case "GetCertificatesStatus":
		return &device.GetCertificatesStatus{}, nil
	case "SetCertificatesStatus":
		return &device.SetCertificatesStatus{}, nil
	case "DeleteCertificates":
		return &device.DeleteCertificates{}, nil
	case "GetPkcs10Request":
		return &device.GetPkcs10Request{}, nil
	case "LoadCertificates":
		return &device.LoadCertificates{}, nil
	case "GetClientCertificateMode":
		return &device.GetClientCertificateMode{}, nil
	case "SetClientCertificateMode":
		return &device.SetClientCertificateMode{}, nil
	case "GetRelayOutputs":
		return &device.GetRelayOutputs{}, nil
	case "SetRelayOutputSettings":
		return &device.SetRelayOutputSettings{}, nil
	case "SetRelayOutputState":
		return &device.SetRelayOutputState{}, nil
	case "SendAuxiliaryCommand":
		return &device.SendAuxiliaryCommand{}, nil
	case "GetCACertificates":
		return &device.GetCACertificates{}, nil
	case "LoadCertificateWithPrivateKey":
		return &device.LoadCertificateWithPrivateKey{}, nil
	case "GetCertificateInformation":
		return &device.GetCertificateInformation{}, nil
	case "LoadCACertificates":
		return &device.LoadCACertificates{}, nil
	case "CreateDot1XConfiguration":
		return &device.CreateDot1XConfiguration{}, nil
	case "SetDot1XConfiguration":
		return &device.SetDot1XConfiguration{}, nil
	case "GetDot1XConfiguration":
		return &device.GetDot1XConfiguration{}, nil
	case "GetDot1XConfigurations":
		return &device.GetDot1XConfigurations{}, nil
	case "DeleteDot1XConfiguration":
		return &device.DeleteDot1XConfiguration{}, nil
	case "GetDot11Capabilities":
		return &device.GetDot11Capabilities{}, nil
	case "GetDot11Status":
		return &device.GetDot11Status{}, nil
	case "ScanAvailableDot11Networks":
		return &device.ScanAvailableDot11Networks{}, nil
	case "GetSystemUris":
		return &device.GetSystemUris{}, nil
	case "StartFirmwareUpgrade":
		return &device.StartFirmwareUpgrade{}, nil
	case "StartSystemRestore":
		return &device.StartSystemRestore{}, nil
	case "GetStorageConfigurations":
		return &device.GetStorageConfigurations{}, nil
	case "CreateStorageConfiguration":
		return &device.CreateStorageConfiguration{}, nil
	case "GetStorageConfiguration":
		return &device.GetStorageConfiguration{}, nil
	case "SetStorageConfiguration":
		return &device.SetStorageConfiguration{}, nil
	case "DeleteStorageConfiguration":
		return &device.DeleteStorageConfiguration{}, nil
	case "GetGeoLocation":
		return &device.GetGeoLocation{}, nil
	case "SetGeoLocation":
		return &device.SetGeoLocation{}, nil
	case "DeleteGeoLocation":
		return &device.DeleteGeoLocation{}, nil
	default:
		return nil, errors.New("there is no such method in the Device service")
	}
}

func getMediaStructByName(name string) (interface{}, error) {
	switch name {
	case "GetServiceCapabilities":
		return &media.GetServiceCapabilities{}, nil
	case "GetVideoSources":
		return &media.GetVideoSources{}, nil
	case "GetAudioSources":
		return &media.GetAudioSources{}, nil
	case "GetAudioOutputs":
		return &media.GetAudioOutputs{}, nil
	case "CreateProfile":
		return &media.CreateProfile{}, nil
	case "GetProfile":
		return &media.GetProfile{}, nil
	case "GetProfiles":
		return &media.GetProfiles{}, nil
	case "AddVideoEncoderConfiguration":
		return &media.AddVideoEncoderConfiguration{}, nil
	case "RemoveVideoEncoderConfiguration":
		return &media.RemoveVideoEncoderConfiguration{}, nil
	case "AddVideoSourceConfiguration":
		return &media.AddVideoSourceConfiguration{}, nil
	case "RemoveVideoSourceConfiguration":
		return &media.RemoveVideoSourceConfiguration{}, nil
	case "AddAudioEncoderConfiguration":
		return &media.AddAudioEncoderConfiguration{}, nil
	case "RemoveAudioEncoderConfiguration":
		return &media.RemoveAudioEncoderConfiguration{}, nil
	case "AddAudioSourceConfiguration":
		return &media.AddAudioSourceConfiguration{}, nil
	case "RemoveAudioSourceConfiguration":
		return &media.RemoveAudioSourceConfiguration{}, nil
	case "AddPTZConfiguration":
		return &media.AddPTZConfiguration{}, nil
	case "RemovePTZConfiguration":
		return &media.RemovePTZConfiguration{}, nil
	case "AddVideoAnalyticsConfiguration":
		return &media.AddVideoAnalyticsConfiguration{}, nil
	case "RemoveVideoAnalyticsConfiguration":
		return &media.RemoveVideoAnalyticsConfiguration{}, nil
	case "AddMetadataConfiguration":
		return &media.AddMetadataConfiguration{}, nil
	case "RemoveMetadataConfiguration":
		return &media.RemoveMetadataConfiguration{}, nil
	case "AddAudioOutputConfiguration":
		return &media.AddAudioOutputConfiguration{}, nil
	case "RemoveAudioOutputConfiguration":
		return &media.RemoveAudioOutputConfiguration{}, nil
	case "AddAudioDecoderConfiguration":
		return &media.AddAudioDecoderConfiguration{}, nil
	case "RemoveAudioDecoderConfiguration":
		return &media.RemoveAudioDecoderConfiguration{}, nil
	case "DeleteProfile":
		return &media.DeleteProfile{}, nil
	case "GetVideoSourceConfigurations":
		return &media.GetVideoSourceConfigurations{}, nil
	case "GetVideoEncoderConfigurations":
		return &media.GetVideoEncoderConfigurations{}, nil
	case "GetAudioSourceConfigurations":
		return &media.GetAudioSourceConfigurations{}, nil
	case "GetAudioEncoderConfigurations":
		return &media.GetAudioEncoderConfigurations{}, nil
	case "GetVideoAnalyticsConfigurations":
		return &media.GetVideoAnalyticsConfigurations{}, nil
	case "GetMetadataConfigurations":
		return &media.GetMetadataConfigurations{}, nil
	case "GetAudioOutputConfigurations":
		return &media.GetAudioOutputConfigurations{}, nil
	case "GetAudioDecoderConfigurations":
		return &media.GetAudioDecoderConfigurations{}, nil
	case "GetVideoSourceConfiguration":
		return &media.GetVideoSourceConfiguration{}, nil
	case "GetVideoEncoderConfiguration":
		return &media.GetVideoEncoderConfiguration{}, nil
	case "GetAudioSourceConfiguration":
		return &media.GetAudioSourceConfiguration{}, nil
	case "GetAudioEncoderConfiguration":
		return &media.GetAudioEncoderConfiguration{}, nil
	case "GetVideoAnalyticsConfiguration":
		return &media.GetVideoAnalyticsConfiguration{}, nil
	case "GetMetadataConfiguration":
		return &media.GetMetadataConfiguration{}, nil
	case "GetAudioOutputConfiguration":
		return &media.GetAudioOutputConfiguration{}, nil
	case "GetAudioDecoderConfiguration":
		return &media.GetAudioDecoderConfiguration{}, nil
	case "GetCompatibleVideoEncoderConfigurations":
		return &media.GetCompatibleVideoEncoderConfigurations{}, nil
	case "GetCompatibleVideoSourceConfigurations":
		return &media.GetCompatibleVideoSourceConfigurations{}, nil
	case "GetCompatibleAudioEncoderConfigurations":
		return &media.GetCompatibleAudioEncoderConfigurations{}, nil
	case "GetCompatibleAudioSourceConfigurations":
		return &media.GetCompatibleAudioSourceConfigurations{}, nil
	case "GetCompatibleVideoAnalyticsConfigurations":
		return &media.GetCompatibleVideoAnalyticsConfigurations{}, nil
	case "GetCompatibleMetadataConfigurations":
		return &media.GetCompatibleMetadataConfigurations{}, nil
	case "GetCompatibleAudioOutputConfigurations":
		return &media.GetCompatibleAudioOutputConfigurations{}, nil
	case "GetCompatibleAudioDecoderConfigurations":
		return &media.GetCompatibleAudioDecoderConfigurations{}, nil
	case "SetVideoSourceConfiguration":
		return &media.SetVideoSourceConfiguration{}, nil
	case "SetVideoEncoderConfiguration":
		return &media.SetVideoEncoderConfiguration{}, nil
	case "SetAudioSourceConfiguration":
		return &media.SetAudioSourceConfiguration{}, nil
	case "SetAudioEncoderConfiguration":
		return &media.SetAudioEncoderConfiguration{}, nil
	case "SetVideoAnalyticsConfiguration":
		return &media.SetVideoAnalyticsConfiguration{}, nil
	case "SetMetadataConfiguration":
		return &media.SetMetadataConfiguration{}, nil
	case "SetAudioOutputConfiguration":
		return &media.SetAudioOutputConfiguration{}, nil
	case "SetAudioDecoderConfiguration":
		return &media.SetAudioDecoderConfiguration{}, nil
	case "GetVideoSourceConfigurationOptions":
		return &media.GetVideoSourceConfigurationOptions{}, nil
	case "GetVideoEncoderConfigurationOptions":
		return &media.GetVideoEncoderConfigurationOptions{}, nil
	case "GetAudioSourceConfigurationOptions":
		return &media.GetAudioSourceConfigurationOptions{}, nil
	case "GetAudioEncoderConfigurationOptions":
		return &media.GetAudioEncoderConfigurationOptions{}, nil
	case "GetMetadataConfigurationOptions":
		return &media.GetMetadataConfigurationOptions{}, nil
	case "GetAudioOutputConfigurationOptions":
		return &media.GetAudioOutputConfigurationOptions{}, nil
	case "GetAudioDecoderConfigurationOptions":
		return &media.GetAudioDecoderConfigurationOptions{}, nil
	case "GetGuaranteedNumberOfVideoEncoderInstances":
		return &media.GetGuaranteedNumberOfVideoEncoderInstances{}, nil
	case "GetStreamUri":
		return &media.GetStreamUri{}, nil
	case "StartMulticastStreaming":
		return &media.StartMulticastStreaming{}, nil
	case "StopMulticastStreaming":
		return &media.StopMulticastStreaming{}, nil
	case "SetSynchronizationPoint":
		return &media.SetSynchronizationPoint{}, nil
	case "GetSnapshotUri":
		return &media.GetSnapshotUri{}, nil
	case "GetVideoSourceModes":
		return &media.GetVideoSourceModes{}, nil
	case "SetVideoSourceMode":
		return &media.SetVideoSourceMode{}, nil
	case "GetOSDs":
		return &media.GetOSDs{}, nil
	case "GetOSD":
		return &media.GetOSD{}, nil
	case "GetOSDOptions":
		return &media.GetOSDOptions{}, nil
	case "SetOSD":
		return &media.SetOSD{}, nil
	case "CreateOSD":
		return &media.CreateOSD{}, nil
	case "DeleteOSD":
		return &media.DeleteOSD{}, nil
	default:
		return nil, errors.New("there is no such method in the Media service")
	}

}
