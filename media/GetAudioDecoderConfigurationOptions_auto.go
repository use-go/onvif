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

// Call_GetAudioDecoderConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationOptionsResponse.
func Call_GetAudioDecoderConfigurationOptions(ctx context.Context, dev *onvif.Device, request GetAudioDecoderConfigurationOptions) (GetAudioDecoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationOptionsResponse GetAudioDecoderConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetAudioDecoderConfigurationOptions")
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, errors.Annotate(err, "reply")
	}
}
