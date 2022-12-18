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

// Call_GetDPAddresses forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDPAddressesResponse.
func Call_GetDPAddresses(ctx context.Context, dev *onvif.Device, request GetDPAddresses) (GetDPAddressesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDPAddressesResponse GetDPAddressesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDPAddressesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetDPAddresses")
		return reply.Body.GetDPAddressesResponse, errors.Annotate(err, "reply")
	}
}
