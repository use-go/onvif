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

// Call_GetPresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPresetTourResponse.
func Call_GetPresetTour(ctx context.Context, dev *device.Device, request GetPresetTour) (GetPresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPresetTourResponse GetPresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPresetTourResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetPresetTour")
		return reply.Body.GetPresetTourResponse, errors.Annotate(err, "reply")
	}
}
