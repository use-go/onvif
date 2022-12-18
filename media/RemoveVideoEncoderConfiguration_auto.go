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

// Call_RemoveVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveVideoEncoderConfigurationResponse.
func Call_RemoveVideoEncoderConfiguration(ctx context.Context, dev *device.Device, request RemoveVideoEncoderConfiguration) (RemoveVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoEncoderConfigurationResponse RemoveVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveVideoEncoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemoveVideoEncoderConfiguration")
		return reply.Body.RemoveVideoEncoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
