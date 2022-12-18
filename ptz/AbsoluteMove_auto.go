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

// Call_AbsoluteMove forwards the call to dev.CallMethod() then parses the payload of the reply as a AbsoluteMoveResponse.
func Call_AbsoluteMove(ctx context.Context, dev *device.Device, request AbsoluteMove) (AbsoluteMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AbsoluteMoveResponse AbsoluteMoveResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AbsoluteMoveResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "AbsoluteMove")
		return reply.Body.AbsoluteMoveResponse, errors.Annotate(err, "reply")
	}
}
