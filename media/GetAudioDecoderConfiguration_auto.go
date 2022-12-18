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

// Call_GetAudioDecoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationResponse.
func Call_GetAudioDecoderConfiguration(ctx context.Context, dev *onvif.Device, request GetAudioDecoderConfiguration) (GetAudioDecoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationResponse GetAudioDecoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetAudioDecoderConfiguration")
		return reply.Body.GetAudioDecoderConfigurationResponse, errors.Annotate(err, "reply")
	}
}
