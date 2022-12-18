// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_SetRemoteUser forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRemoteUserResponse.
func Call_SetRemoteUser(ctx context.Context, dev *Device, request SetRemoteUser) (SetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteUserResponse SetRemoteUserResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetRemoteUserResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetRemoteUser")
		return reply.Body.SetRemoteUserResponse, errors.Annotate(err, "reply")
	}
}
