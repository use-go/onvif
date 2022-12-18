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

// Call_AddAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioOutputConfigurationResponse.
func Call_AddAudioOutputConfiguration(ctx context.Context, dev *device.Device, request AddAudioOutputConfiguration) (AddAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioOutputConfigurationResponse AddAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddAudioOutputConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "AddAudioOutputConfiguration")
		return reply.Body.AddAudioOutputConfigurationResponse, errors.Annotate(err, "reply")
	}
}
