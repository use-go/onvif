// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_SystemReboot forwards the call to dev.CallMethod() then parses the payload of the reply as a SystemRebootResponse.
func Call_SystemReboot(ctx context.Context, dev *Device, request SystemReboot) (SystemRebootResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SystemRebootResponse SystemRebootResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SystemRebootResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SystemReboot")
		return reply.Body.SystemRebootResponse, errors.Annotate(err, "reply")
	}
}
