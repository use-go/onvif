// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_GetCompatibleVideoEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoEncoderConfigurationsResponse.
func Call_GetCompatibleVideoEncoderConfigurations(ctx context.Context, dev *device.Device, request GetCompatibleVideoEncoderConfigurations) (GetCompatibleVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoEncoderConfigurationsResponse GetCompatibleVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleVideoEncoderConfigurations")
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
