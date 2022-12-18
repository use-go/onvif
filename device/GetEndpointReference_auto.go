// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetEndpointReference forwards the call to dev.CallMethod() then parses the payload of the reply as a GetEndpointReferenceResponse.
func Call_GetEndpointReference(ctx context.Context, dev *Device, request GetEndpointReference) (GetEndpointReferenceResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEndpointReferenceResponse GetEndpointReferenceResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetEndpointReferenceResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetEndpointReference")
		return reply.Body.GetEndpointReferenceResponse, errors.Annotate(err, "reply")
	}
}
