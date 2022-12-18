// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"context"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/ptz"
)

type PtzOutput struct { 
	Configurations                           *ptz.GetConfigurationsResponse
	Nodes                                    *ptz.GetNodesResponse
	ServiceCapabilities                      *ptz.GetServiceCapabilitiesResponse
}

func detailPtz(ctx context.Context, dev *device.Device) PtzOutput {
	var out PtzOutput
	calls := make([]func(c context.Context), 0)

	calls = append(calls, func(c context.Context) {
		if p, err := ptz.Call_GetConfigurations(c, dev, ptz.GetConfigurations {}); err == nil {
			out.Configurations = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Configurations").Msg("ptz")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := ptz.Call_GetNodes(c, dev, ptz.GetNodes {}); err == nil {
			out.Nodes = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "Nodes").Msg("ptz")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := ptz.Call_GetServiceCapabilities(c, dev, ptz.GetServiceCapabilities {}); err == nil {
			out.ServiceCapabilities = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "ServiceCapabilities").Msg("ptz")
		}
	})

	runAll(ctx, calls...)
	return out
}
