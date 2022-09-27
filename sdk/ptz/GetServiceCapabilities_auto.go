// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/ptz"
)

// Call_GetServiceCapabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetServiceCapabilitiesResponse.
func Call_GetServiceCapabilities(ctx context.Context, dev *onvif.Device, request ptz.GetServiceCapabilities) (ptz.GetServiceCapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServiceCapabilitiesResponse ptz.GetServiceCapabilitiesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetServiceCapabilitiesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetServiceCapabilities")
		return reply.Body.GetServiceCapabilitiesResponse, err
	}
}
