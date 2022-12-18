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

// Call_CreateUsers forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateUsersResponse.
func Call_CreateUsers(ctx context.Context, dev *onvif.Device, request CreateUsers) (CreateUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateUsersResponse CreateUsersResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateUsersResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "CreateUsers")
		return reply.Body.CreateUsersResponse, errors.Annotate(err, "reply")
	}
}
