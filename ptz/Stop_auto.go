// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_Stop forwards the call to dev.CallMethod() then parses the payload of the reply as a StopResponse.
func Call_Stop(ctx context.Context, dev *onvif.Device, request Stop) (StopResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopResponse StopResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StopResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "Stop")
		return reply.Body.StopResponse, errors.Annotate(err, "reply")
	}
}
