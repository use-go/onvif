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

// Call_CreatePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a CreatePresetTourResponse.
func Call_CreatePresetTour(ctx context.Context, dev *onvif.Device, request CreatePresetTour) (CreatePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreatePresetTourResponse CreatePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreatePresetTourResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "CreatePresetTour")
		return reply.Body.CreatePresetTourResponse, errors.Annotate(err, "reply")
	}
}
