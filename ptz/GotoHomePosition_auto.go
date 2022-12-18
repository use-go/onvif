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

// Call_GotoHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a GotoHomePositionResponse.
func Call_GotoHomePosition(ctx context.Context, dev *device.Device, request GotoHomePosition) (GotoHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoHomePositionResponse GotoHomePositionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GotoHomePositionResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GotoHomePosition")
		return reply.Body.GotoHomePositionResponse, errors.Annotate(err, "reply")
	}
}
