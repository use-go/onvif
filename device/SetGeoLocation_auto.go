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

// Call_SetGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a SetGeoLocationResponse.
func Call_SetGeoLocation(ctx context.Context, dev *onvif.Device, request SetGeoLocation) (SetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetGeoLocationResponse SetGeoLocationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetGeoLocationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetGeoLocation")
		return reply.Body.SetGeoLocationResponse, errors.Annotate(err, "reply")
	}
}
