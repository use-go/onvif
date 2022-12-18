// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_ModifyPresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a ModifyPresetTourResponse.
func Call_ModifyPresetTour(ctx context.Context, dev *onvif.Device, request ModifyPresetTour) (ModifyPresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			ModifyPresetTourResponse ModifyPresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.ModifyPresetTourResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "ModifyPresetTour")
		return reply.Body.ModifyPresetTourResponse, errors.Annotate(err, "reply")
	}
}
