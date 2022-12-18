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

// Call_RemoveAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioOutputConfigurationResponse.
func Call_RemoveAudioOutputConfiguration(ctx context.Context, dev *device.Device, request RemoveAudioOutputConfiguration) (RemoveAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioOutputConfigurationResponse RemoveAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveAudioOutputConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemoveAudioOutputConfiguration")
		return reply.Body.RemoveAudioOutputConfigurationResponse, errors.Annotate(err, "reply")
	}
}
