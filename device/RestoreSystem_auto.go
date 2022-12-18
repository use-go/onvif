// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_RestoreSystem forwards the call to dev.CallMethod() then parses the payload of the reply as a RestoreSystemResponse.
func Call_RestoreSystem(ctx context.Context, dev *Device, request RestoreSystem) (RestoreSystemResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RestoreSystemResponse RestoreSystemResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RestoreSystemResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "RestoreSystem")
		return reply.Body.RestoreSystemResponse, errors.Annotate(err, "reply")
	}
}
