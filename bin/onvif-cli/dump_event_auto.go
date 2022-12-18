// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package main

import (
	"context"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/event"
)

type EventOutput struct { 
	EventProperties                          *event.GetEventPropertiesResponse
	ServiceCapabilities                      *event.GetServiceCapabilitiesResponse
}

func detailEvent(ctx context.Context, dev *device.Device) EventOutput {
	var out EventOutput
	calls := make([]func(c context.Context), 0)

	calls = append(calls, func(c context.Context) {
		if p, err := event.Call_GetEventProperties(c, dev, event.GetEventProperties {}); err == nil {
			out.EventProperties = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "EventProperties").Msg("event")
		}
	})

	calls = append(calls, func(c context.Context) {
		if p, err := event.Call_GetServiceCapabilities(c, dev, event.GetServiceCapabilities {}); err == nil {
			out.ServiceCapabilities = &p
		} else {
			Logger.Trace().Err(err).Str("rpc", "ServiceCapabilities").Msg("event")
		}
	})

	runAll(ctx, calls...)
	return out
}
