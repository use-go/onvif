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

// Call_GetAudioSources forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourcesResponse.
func Call_GetAudioSources(ctx context.Context, dev *onvif.Device, request GetAudioSources) (GetAudioSourcesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourcesResponse GetAudioSourcesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioSourcesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetAudioSources")
		return reply.Body.GetAudioSourcesResponse, errors.Annotate(err, "reply")
	}
}
