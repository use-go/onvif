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

// Call_SetScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a SetScopesResponse.
func Call_SetScopes(ctx context.Context, dev *onvif.Device, request SetScopes) (SetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetScopesResponse SetScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetScopesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetScopes")
		return reply.Body.SetScopesResponse, errors.Annotate(err, "reply")
	}
}
