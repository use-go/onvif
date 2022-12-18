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

// Call_GetGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetGeoLocationResponse.
func Call_GetGeoLocation(ctx context.Context, dev *onvif.Device, request GetGeoLocation) (GetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetGeoLocationResponse GetGeoLocationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetGeoLocationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetGeoLocation")
		return reply.Body.GetGeoLocationResponse, errors.Annotate(err, "reply")
	}
}
