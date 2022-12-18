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

// Call_GetEventProperties forwards the call to dev.CallMethod() then parses the payload of the reply as a GetEventPropertiesResponse.
func Call_GetEventProperties(ctx context.Context, dev *device.Device, request GetEventProperties) (GetEventPropertiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEventPropertiesResponse GetEventPropertiesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetEventPropertiesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetEventProperties")
		return reply.Body.GetEventPropertiesResponse, errors.Annotate(err, "reply")
	}
}
