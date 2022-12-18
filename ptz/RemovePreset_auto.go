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

// Call_RemovePreset forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePresetResponse.
func Call_RemovePreset(ctx context.Context, dev *onvif.Device, request RemovePreset) (RemovePresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePresetResponse RemovePresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePresetResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemovePreset")
		return reply.Body.RemovePresetResponse, errors.Annotate(err, "reply")
	}
}
