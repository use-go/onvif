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

// Call_DeleteOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteOSDResponse.
func Call_DeleteOSD(ctx context.Context, dev *device.Device, request DeleteOSD) (DeleteOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteOSDResponse DeleteOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteOSDResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "DeleteOSD")
		return reply.Body.DeleteOSDResponse, errors.Annotate(err, "reply")
	}
}
