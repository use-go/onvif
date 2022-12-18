// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package event

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_Subscribe forwards the call to dev.CallMethod() then parses the payload of the reply as a SubscribeResponse.
func Call_Subscribe(ctx context.Context, dev *device.Device, request Subscribe) (SubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SubscribeResponse SubscribeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SubscribeResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "Subscribe")
		return reply.Body.SubscribeResponse, errors.Annotate(err, "reply")
	}
}
