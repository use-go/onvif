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

// Call_GetPkcs10Request forwards the call to dev.CallMethod() then parses the payload of the reply as a GetPkcs10RequestResponse.
func Call_GetPkcs10Request(ctx context.Context, dev *onvif.Device, request GetPkcs10Request) (GetPkcs10RequestResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetPkcs10RequestResponse GetPkcs10RequestResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetPkcs10RequestResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetPkcs10Request")
		return reply.Body.GetPkcs10RequestResponse, errors.Annotate(err, "reply")
	}
}
