// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetHostname forwards the call to dev.CallMethod() then parses the payload of the reply as a GetHostnameResponse.
func Call_GetHostname(ctx context.Context, dev *Device, request GetHostname) (GetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetHostnameResponse GetHostnameResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetHostnameResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetHostname")
		return reply.Body.GetHostnameResponse, errors.Annotate(err, "reply")
	}
}
