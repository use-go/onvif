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

// Call_RemoveAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioEncoderConfigurationResponse.
func Call_RemoveAudioEncoderConfiguration(ctx context.Context, dev *device.Device, request RemoveAudioEncoderConfiguration) (RemoveAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioEncoderConfigurationResponse RemoveAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemoveAudioEncoderConfiguration")
		return reply.Body.RemoveAudioEncoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
