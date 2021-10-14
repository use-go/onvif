package main

import (
	"errors"
	"fmt"
	"github.com/use-go/onvif"

	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
)

func GetResponseStructByCommand(command string) (interface{}, error) {
	switch command {
	// WebService - Device - Network Configuration
	case onvif.GetHostname:
		return &device.GetHostnameResponse{}, nil
	case onvif.SetHostname:
		return &device.SetHostnameResponse{}, nil
	case onvif.GetDNS:
		return &device.GetDNSResponse{}, nil
	case onvif.SetDNS:
		return &device.SetDNSResponse{}, nil
	case onvif.GetNetworkInterfaces:
		return &device.GetNetworkInterfacesResponse{}, nil
	case onvif.SetNetworkInterfaces:
		return &device.SetNetworkInterfacesResponse{}, nil
	case onvif.GetNetworkProtocols:
		return &device.GetNetworkProtocolsResponse{}, nil
	case onvif.SetNetworkProtocols:
		return &device.SetNetworkProtocolsResponse{}, nil
	case onvif.GetNetworkDefaultGateway:
		return &device.GetNetworkDefaultGatewayResponse{}, nil
	case onvif.SetNetworkDefaultGateway:
		return &device.SetNetworkDefaultGatewayResponse{}, nil

	// WebService - Device - System Function
	case onvif.GetDeviceInformation:
		return &device.GetDeviceInformationResponse{}, nil
	case onvif.GetSystemDateAndTime:
		return &device.GetSystemDateAndTimeResponse{}, nil
	case onvif.SetSystemDateAndTime:
		return &device.SetSystemDateAndTimeResponse{}, nil
	case onvif.SetSystemFactoryDefault:
		return &device.SetSystemFactoryDefaultResponse{}, nil
	case onvif.SystemReboot:
		return &device.SystemRebootResponse{}, nil

	// WebService - Media - Metadata Configuration
	case onvif.GetMetadataConfiguration:
		return &media.GetMetadataConfigurationResponse{}, nil
	case onvif.GetMetadataConfigurations:
		return &media.GetMetadataConfigurationsResponse{}, nil
	case onvif.AddMetadataConfiguration:
		return &media.AddMetadataConfigurationResponse{}, nil
	case onvif.RemoveMetadataConfiguration:
		return &media.RemoveMetadataConfigurationResponse{}, nil
	case onvif.SetMetadataConfiguration:
		return &media.SetMetadataConfigurationResponse{}, nil
	case onvif.GetCompatibleMetadataConfigurations:
		return &media.GetCompatibleMetadataConfigurationsResponse{}, nil
	case onvif.GetMetadataConfigurationOptions:
		return &media.GetMetadataConfigurationOptionsResponse{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("not support the command '%s'", command))
	}
}
