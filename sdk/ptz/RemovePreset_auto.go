// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/ptz"
)

// Call_RemovePreset forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePresetResponse.
func Call_RemovePreset(ctx context.Context, dev *onvif.Device, request ptz.RemovePreset) (ptz.RemovePresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePresetResponse ptz.RemovePresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePresetResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemovePreset")
		return reply.Body.RemovePresetResponse, err
	}
}
