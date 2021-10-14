package main

import (
	"errors"
	"fmt"
	"github.com/use-go/onvif"

	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/media"
)

func GetRequestStructByCommand(command string) (interface{}, error) {
	switch command {
	// WebService - Device - Network Configuration
	case onvif.GetHostname:
		return &device.GetHostname{}, nil
	case onvif.SetHostname:
		return &device.SetHostname{}, nil
	case onvif.GetDNS:
		return &device.GetDNS{}, nil
	case onvif.SetDNS:
		return &device.SetDNS{}, nil
	case onvif.GetNetworkInterfaces:
		return &device.GetNetworkInterfaces{}, nil
	case onvif.SetNetworkInterfaces:
		return &device.SetNetworkInterfaces{}, nil
	case onvif.GetNetworkProtocols:
		return &device.GetNetworkProtocols{}, nil
	case onvif.SetNetworkProtocols:
		return &device.SetNetworkProtocols{}, nil
	case onvif.GetNetworkDefaultGateway:
		return &device.GetNetworkDefaultGateway{}, nil
	case onvif.SetNetworkDefaultGateway:
		return &device.SetNetworkDefaultGateway{}, nil

	// WebService - Device - System Function
	case onvif.GetDeviceInformation:
		return &device.GetDeviceInformation{}, nil
	case onvif.GetSystemDateAndTime:
		return &device.GetSystemDateAndTime{}, nil
	case onvif.SetSystemDateAndTime:
		return &device.SetSystemDateAndTime{}, nil
	case onvif.SetSystemFactoryDefault:
		return &device.SetSystemFactoryDefault{}, nil
	case onvif.SystemReboot:
		return &device.SystemReboot{}, nil

	// WebService - Media - Metadata Configuration
	case onvif.GetMetadataConfiguration:
		return &media.GetMetadataConfiguration{}, nil
	case onvif.GetMetadataConfigurations:
		return &media.GetMetadataConfigurations{}, nil
	case onvif.AddMetadataConfiguration:
		return &media.AddMetadataConfiguration{}, nil
	case onvif.RemoveMetadataConfiguration:
		return &media.RemoveMetadataConfiguration{}, nil
	case onvif.SetMetadataConfiguration:
		return &media.SetMetadataConfiguration{}, nil
	case onvif.GetCompatibleMetadataConfigurations:
		return &media.GetCompatibleMetadataConfigurations{}, nil
	case onvif.GetMetadataConfigurationOptions:
		return &media.GetMetadataConfigurationOptions{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("not support the command '%s'", command))
	}
}
