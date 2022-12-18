// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_GetConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationOptionsResponse.
func Call_GetConfigurationOptions(ctx context.Context, dev *device.Device, request GetConfigurationOptions) (GetConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationOptionsResponse GetConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetConfigurationOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetConfigurationOptions")
		return reply.Body.GetConfigurationOptionsResponse, errors.Annotate(err, "reply")
	}
}
