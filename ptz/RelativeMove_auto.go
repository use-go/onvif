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

// Call_RelativeMove forwards the call to dev.CallMethod() then parses the payload of the reply as a RelativeMoveResponse.
func Call_RelativeMove(ctx context.Context, dev *device.Device, request RelativeMove) (RelativeMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RelativeMoveResponse RelativeMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RelativeMoveResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RelativeMove")
		return reply.Body.RelativeMoveResponse, errors.Annotate(err, "reply")
	}
}
