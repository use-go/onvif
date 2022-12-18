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

// Call_SetHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHomePositionResponse.
func Call_SetHomePosition(ctx context.Context, dev *device.Device, request SetHomePosition) (SetHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHomePositionResponse SetHomePositionResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetHomePositionResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetHomePosition")
		return reply.Body.SetHomePositionResponse, errors.Annotate(err, "reply")
	}
}
