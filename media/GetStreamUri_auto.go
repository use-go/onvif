// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_GetStreamUri forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStreamUriResponse.
func Call_GetStreamUri(ctx context.Context, dev *device.Device, request GetStreamUri) (GetStreamUriResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStreamUriResponse GetStreamUriResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStreamUriResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetStreamUri")
		return reply.Body.GetStreamUriResponse, errors.Annotate(err, "reply")
	}
}
