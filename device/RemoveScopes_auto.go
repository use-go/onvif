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

// Call_RemoveScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveScopesResponse.
func Call_RemoveScopes(ctx context.Context, dev *onvif.Device, request RemoveScopes) (RemoveScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveScopesResponse RemoveScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveScopesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RemoveScopes")
		return reply.Body.RemoveScopesResponse, errors.Annotate(err, "reply")
	}
}
