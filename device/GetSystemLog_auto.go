// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetSystemLog forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemLogResponse.
func Call_GetSystemLog(ctx context.Context, dev *Device, request GetSystemLog) (GetSystemLogResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemLogResponse GetSystemLogResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemLogResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetSystemLog")
		return reply.Body.GetSystemLogResponse, errors.Annotate(err, "reply")
	}
}
