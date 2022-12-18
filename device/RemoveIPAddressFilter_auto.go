// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_RemoveIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveIPAddressFilterResponse.
func Call_RemoveIPAddressFilter(ctx context.Context, dev *Device, request RemoveIPAddressFilter) (RemoveIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveIPAddressFilterResponse RemoveIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveIPAddressFilterResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemoveIPAddressFilter")
		return reply.Body.RemoveIPAddressFilterResponse, errors.Annotate(err, "reply")
	}
}
