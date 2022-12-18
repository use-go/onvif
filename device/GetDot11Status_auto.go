// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetDot11Status forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot11StatusResponse.
func Call_GetDot11Status(ctx context.Context, dev *Device, request GetDot11Status) (GetDot11StatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11StatusResponse GetDot11StatusResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot11StatusResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetDot11Status")
		return reply.Body.GetDot11StatusResponse, errors.Annotate(err, "reply")
	}
}
