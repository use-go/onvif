// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a GetUsersResponse.
func Call_GetUsers(ctx context.Context, dev *Device, request GetUsers) (GetUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetUsersResponse GetUsersResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetUsersResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetUsers")
		return reply.Body.GetUsersResponse, errors.Annotate(err, "reply")
	}
}
