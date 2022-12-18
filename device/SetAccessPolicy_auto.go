// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_SetAccessPolicy forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAccessPolicyResponse.
func Call_SetAccessPolicy(ctx context.Context, dev *Device, request SetAccessPolicy) (SetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAccessPolicyResponse SetAccessPolicyResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAccessPolicyResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetAccessPolicy")
		return reply.Body.SetAccessPolicyResponse, errors.Annotate(err, "reply")
	}
}
