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

// Call_PullMessages forwards the call to dev.CallMethod() then parses the payload of the reply as a PullMessagesResponse.
func Call_PullMessages(ctx context.Context, dev *device.Device, request PullMessages) (PullMessagesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			PullMessagesResponse PullMessagesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.PullMessagesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "PullMessages")
		return reply.Body.PullMessagesResponse, errors.Annotate(err, "reply")
	}
}
