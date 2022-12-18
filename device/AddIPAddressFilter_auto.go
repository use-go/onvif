// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_AddIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a AddIPAddressFilterResponse.
func Call_AddIPAddressFilter(ctx context.Context, dev *Device, request AddIPAddressFilter) (AddIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddIPAddressFilterResponse AddIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddIPAddressFilterResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "AddIPAddressFilter")
		return reply.Body.AddIPAddressFilterResponse, errors.Annotate(err, "reply")
	}
}
