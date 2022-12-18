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

// Call_GetAudioOutputConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationOptionsResponse.
func Call_GetAudioOutputConfigurationOptions(ctx context.Context, dev *onvif.Device, request GetAudioOutputConfigurationOptions) (GetAudioOutputConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationOptionsResponse GetAudioOutputConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetAudioOutputConfigurationOptions")
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, errors.Annotate(err, "reply")
	}
}
