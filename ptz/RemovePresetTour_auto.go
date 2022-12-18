// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_RemovePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePresetTourResponse.
func Call_RemovePresetTour(ctx context.Context, dev *device.Device, request RemovePresetTour) (RemovePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePresetTourResponse RemovePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePresetTourResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemovePresetTour")
		return reply.Body.RemovePresetTourResponse, errors.Annotate(err, "reply")
	}
}
