// Copyright (c) 2022-2022 Jean-Francois SMIGIELSKI

package main

import (
	"context"
	"encoding/json"
	wsdiscovery "github.com/use-go/onvif/ws-discovery"
	"net"
	"net/url"
	"os"

	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
)

func discover(ctx context.Context) error {
	type Output struct {
		device.GetDeviceInformationResponse
		Error     error
		Interface string
		Endpoint  string
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return errors.Annotate(err, "lan discovery")
	}

	encoder := json.NewEncoder(os.Stdout)
	for _, itf := range interfaces {
		devices, err := wsdiscovery.GetAvailableDevicesAtSpecificEthernetInterface(itf.Name)
		if err != nil {
			Logger.Warn().Str("itf", itf.Name).Msg("lan discovery failed")
		} else {
			for _, dev := range devices {
				u := dev.GetEndpoint("device")
				parsedUrl, err := url.Parse(u)
				authDev, err := device.NewDevice(device.DeviceParams{
					Xaddr:    parsedUrl.Host,
					Username: "admin",
					Password: "ollyhgqo",
				})
				if err != nil {
					Logger.Warn().Str("itf", itf.Name).Msg("auth failed")
				} else {
					dev = *authDev
				}

				reply, err := device.Call_GetDeviceInformation(ctx, &dev, device.GetDeviceInformation{})
				out := Output{GetDeviceInformationResponse: reply}
				out.Error = err
				out.Interface = itf.Name
				out.Endpoint = parsedUrl.Host
				encoder.Encode(out)
			}
		}
	}
	return nil
}
