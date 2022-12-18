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

// Call_GetVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoEncoderConfigurationResponse.
func Call_GetVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request GetVideoEncoderConfiguration) (GetVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationResponse GetVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoEncoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetVideoEncoderConfiguration")
		return reply.Body.GetVideoEncoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
