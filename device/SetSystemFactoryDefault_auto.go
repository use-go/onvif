// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_SetSystemFactoryDefault forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSystemFactoryDefaultResponse.
func Call_SetSystemFactoryDefault(ctx context.Context, dev *onvif.Device, request SetSystemFactoryDefault) (SetSystemFactoryDefaultResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemFactoryDefaultResponse SetSystemFactoryDefaultResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetSystemFactoryDefault")
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Annotate(err, "reply")
	}
}
