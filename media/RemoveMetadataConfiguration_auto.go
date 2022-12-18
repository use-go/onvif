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

// Call_RemoveMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveMetadataConfigurationResponse.
func Call_RemoveMetadataConfiguration(ctx context.Context, dev *device.Device, request RemoveMetadataConfiguration) (RemoveMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveMetadataConfigurationResponse RemoveMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemoveMetadataConfiguration")
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Annotate(err, "reply")
	}
}
