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

// Call_Renew forwards the call to dev.CallMethod() then parses the payload of the reply as a RenewResponse.
func Call_Renew(ctx context.Context, dev *device.Device, request Renew) (RenewResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RenewResponse RenewResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RenewResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "Renew")
		return reply.Body.RenewResponse, errors.Annotate(err, "reply")
	}
}
