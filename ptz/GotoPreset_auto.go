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

// Call_GotoPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a GotoPresetResponse.
func Call_GotoPreset(ctx context.Context, dev *device.Device, request GotoPreset) (GotoPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoPresetResponse GotoPresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GotoPresetResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GotoPreset")
		return reply.Body.GotoPresetResponse, errors.Annotate(err, "reply")
	}
}
