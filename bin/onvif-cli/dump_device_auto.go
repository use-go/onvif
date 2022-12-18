// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"context"
	"github.com/use-go/onvif/device"
)

type DeviceOutput struct { 
	AccessPolicy                             *device.GetAccessPolicyResponse
	CACertificates                           *device.GetCACertificatesResponse
	Certificates                             *device.GetCertificatesResponse
	CertificatesStatus                       *device.GetCertificatesStatusResponse
	ClientCertificateMode                    *device.GetClientCertificateModeResponse
	DNS                                      *device.GetDNSResponse
	DPAddresses                              *device.GetDPAddressesResponse
	DeviceInformation                        *device.GetDeviceInformationResponse
	DiscoveryMode                            *device.GetDiscoveryModeResponse
	Dot11Capabilities                        *device.GetDot11CapabilitiesResponse
	Dot1XConfigurations                      *device.GetDot1XConfigurationsResponse
	DynamicDNS                               *device.GetDynamicDNSResponse
	EndpointReference                        *device.GetEndpointReferenceResponse
	GeoLocation                              *device.GetGeoLocationResponse
	Hostname                                 *device.GetHostnameResponse
	IPAddressFilter                          *device.GetIPAddressFilterResponse
	NTP                                      *device.GetNTPResponse
	NetworkDefaultGateway                    *device.GetNetworkDefaultGatewayResponse
	NetworkInterfaces                        *device.GetNetworkInterfacesResponse
	NetworkProtocols                         *device.GetNetworkProtocolsResponse
	RelayOutputs                             *device.GetRelayOutputsResponse
	RemoteDiscoveryMode                      *device.GetRemoteDiscoveryModeResponse
	RemoteUser                               *device.GetRemoteUserResponse
	Scopes                                   *device.GetScopesResponse
	ServiceCapabilities                      *device.GetServiceCapabilitiesResponse
	StorageConfigurations                    *device.GetStorageConfigurationsResponse
	SystemBackup                             *device.GetSystemBackupResponse
	SystemDateAndTime                        *device.GetSystemDateAndTimeResponse
	SystemSupportInformation                 *device.GetSystemSupportInformationResponse
	SystemUris                               *device.GetSystemUrisResponse
	Users                                    *device.GetUsersResponse
	WsdlUrl                                  *device.GetWsdlUrlResponse
	ZeroConfiguration                        *device.GetZeroConfigurationResponse
}

func detailDevice(ctx context.Context, dev *device.Device) DeviceOutput {
	var out DeviceOutput
	calls := make([]func(c context.Context), 0)

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetAccessPolicy(c, dev, device.GetAccessPolicy {}); err == nil {
			out.AccessPolicy = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "AccessPolicy").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetCACertificates(c, dev, device.GetCACertificates {}); err == nil {
			out.CACertificates = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "CACertificates").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetCertificates(c, dev, device.GetCertificates {}); err == nil {
			out.Certificates = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Certificates").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetCertificatesStatus(c, dev, device.GetCertificatesStatus {}); err == nil {
			out.CertificatesStatus = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "CertificatesStatus").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetClientCertificateMode(c, dev, device.GetClientCertificateMode {}); err == nil {
			out.ClientCertificateMode = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "ClientCertificateMode").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetDNS(c, dev, device.GetDNS {}); err == nil {
			out.DNS = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "DNS").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetDPAddresses(c, dev, device.GetDPAddresses {}); err == nil {
			out.DPAddresses = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "DPAddresses").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetDeviceInformation(c, dev, device.GetDeviceInformation {}); err == nil {
			out.DeviceInformation = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "DeviceInformation").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetDiscoveryMode(c, dev, device.GetDiscoveryMode {}); err == nil {
			out.DiscoveryMode = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "DiscoveryMode").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetDot11Capabilities(c, dev, device.GetDot11Capabilities {}); err == nil {
			out.Dot11Capabilities = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Dot11Capabilities").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetDot1XConfigurations(c, dev, device.GetDot1XConfigurations {}); err == nil {
			out.Dot1XConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Dot1XConfigurations").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetDynamicDNS(c, dev, device.GetDynamicDNS {}); err == nil {
			out.DynamicDNS = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "DynamicDNS").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetEndpointReference(c, dev, device.GetEndpointReference {}); err == nil {
			out.EndpointReference = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "EndpointReference").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetGeoLocation(c, dev, device.GetGeoLocation {}); err == nil {
			out.GeoLocation = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "GeoLocation").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetHostname(c, dev, device.GetHostname {}); err == nil {
			out.Hostname = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Hostname").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetIPAddressFilter(c, dev, device.GetIPAddressFilter {}); err == nil {
			out.IPAddressFilter = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "IPAddressFilter").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetNTP(c, dev, device.GetNTP {}); err == nil {
			out.NTP = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "NTP").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetNetworkDefaultGateway(c, dev, device.GetNetworkDefaultGateway {}); err == nil {
			out.NetworkDefaultGateway = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "NetworkDefaultGateway").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetNetworkInterfaces(c, dev, device.GetNetworkInterfaces {}); err == nil {
			out.NetworkInterfaces = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "NetworkInterfaces").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetNetworkProtocols(c, dev, device.GetNetworkProtocols {}); err == nil {
			out.NetworkProtocols = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "NetworkProtocols").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetRelayOutputs(c, dev, device.GetRelayOutputs {}); err == nil {
			out.RelayOutputs = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "RelayOutputs").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetRemoteDiscoveryMode(c, dev, device.GetRemoteDiscoveryMode {}); err == nil {
			out.RemoteDiscoveryMode = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "RemoteDiscoveryMode").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetRemoteUser(c, dev, device.GetRemoteUser {}); err == nil {
			out.RemoteUser = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "RemoteUser").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetScopes(c, dev, device.GetScopes {}); err == nil {
			out.Scopes = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Scopes").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetServiceCapabilities(c, dev, device.GetServiceCapabilities {}); err == nil {
			out.ServiceCapabilities = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "ServiceCapabilities").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetStorageConfigurations(c, dev, device.GetStorageConfigurations {}); err == nil {
			out.StorageConfigurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "StorageConfigurations").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetSystemBackup(c, dev, device.GetSystemBackup {}); err == nil {
			out.SystemBackup = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "SystemBackup").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetSystemDateAndTime(c, dev, device.GetSystemDateAndTime {}); err == nil {
			out.SystemDateAndTime = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "SystemDateAndTime").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetSystemSupportInformation(c, dev, device.GetSystemSupportInformation {}); err == nil {
			out.SystemSupportInformation = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "SystemSupportInformation").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetSystemUris(c, dev, device.GetSystemUris {}); err == nil {
			out.SystemUris = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "SystemUris").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetUsers(c, dev, device.GetUsers {}); err == nil {
			out.Users = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Users").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetWsdlUrl(c, dev, device.GetWsdlUrl {}); err == nil {
			out.WsdlUrl = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "WsdlUrl").Msg("device")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := device.Call_GetZeroConfiguration(c, dev, device.GetZeroConfiguration {}); err == nil {
			out.ZeroConfiguration = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "ZeroConfiguration").Msg("device")
		}
	})

	runAll(ctx, calls...)
	return out
}
