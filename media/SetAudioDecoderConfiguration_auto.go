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

// Call_SetAudioDecoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioDecoderConfigurationResponse.
func Call_SetAudioDecoderConfiguration(ctx context.Context, dev *device.Device, request SetAudioDecoderConfiguration) (SetAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioDecoderConfigurationResponse SetAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioDecoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetAudioDecoderConfiguration")
		return reply.Body.SetAudioDecoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
