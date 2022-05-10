// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/ptz"
)

// Call_RemovePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePresetTourResponse.
func Call_RemovePresetTour(ctx context.Context, dev *onvif.Device, request ptz.RemovePresetTour) (ptz.RemovePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePresetTourResponse ptz.RemovePresetTourResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePresetTourResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemovePresetTour")
		return reply.Body.RemovePresetTourResponse, errors.Annotate(err, "reply")
	}
}
