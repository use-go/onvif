// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_GetVideoEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoEncoderConfigurationsResponse.
func Call_GetVideoEncoderConfigurations(ctx context.Context, dev *onvif.Device, request GetVideoEncoderConfigurations) (GetVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationsResponse GetVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetVideoEncoderConfigurations")
		return reply.Body.GetVideoEncoderConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
