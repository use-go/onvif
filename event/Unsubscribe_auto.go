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

// Call_Unsubscribe forwards the call to dev.CallMethod() then parses the payload of the reply as a UnsubscribeResponse.
func Call_Unsubscribe(ctx context.Context, dev *device.Device, request Unsubscribe) (UnsubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UnsubscribeResponse UnsubscribeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.UnsubscribeResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "Unsubscribe")
		return reply.Body.UnsubscribeResponse, errors.Annotate(err, "reply")
	}
}
