// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a GetIPAddressFilterResponse.
func Call_GetIPAddressFilter(ctx context.Context, dev *Device, request GetIPAddressFilter) (GetIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetIPAddressFilterResponse GetIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetIPAddressFilterResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetIPAddressFilter")
		return reply.Body.GetIPAddressFilterResponse, errors.Annotate(err, "reply")
	}
}
