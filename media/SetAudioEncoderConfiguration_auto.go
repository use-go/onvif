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

// Call_SetAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioEncoderConfigurationResponse.
func Call_SetAudioEncoderConfiguration(ctx context.Context, dev *device.Device, request SetAudioEncoderConfiguration) (SetAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioEncoderConfigurationResponse SetAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioEncoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetAudioEncoderConfiguration")
		return reply.Body.SetAudioEncoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
