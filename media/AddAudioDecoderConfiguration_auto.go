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

// Call_AddAudioDecoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioDecoderConfigurationResponse.
func Call_AddAudioDecoderConfiguration(ctx context.Context, dev *device.Device, request AddAudioDecoderConfiguration) (AddAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioDecoderConfigurationResponse AddAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddAudioDecoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "AddAudioDecoderConfiguration")
		return reply.Body.AddAudioDecoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
