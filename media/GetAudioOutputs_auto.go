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

// Call_GetAudioOutputs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputsResponse.
func Call_GetAudioOutputs(ctx context.Context, dev *device.Device, request GetAudioOutputs) (GetAudioOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputsResponse GetAudioOutputsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputsResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetAudioOutputs")
		return reply.Body.GetAudioOutputsResponse, errors.Annotate(err, "reply")
	}
}
