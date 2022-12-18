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

// Call_GetStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStatusResponse.
func Call_GetStatus(ctx context.Context, dev *device.Device, request GetStatus) (GetStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStatusResponse GetStatusResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStatusResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetStatus")
		return reply.Body.GetStatusResponse, errors.Annotate(err, "reply")
	}
}
