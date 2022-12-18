// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_SendAuxiliaryCommand forwards the call to dev.CallMethod() then parses the payload of the reply as a SendAuxiliaryCommandResponse.
func Call_SendAuxiliaryCommand(ctx context.Context, dev *onvif.Device, request SendAuxiliaryCommand) (SendAuxiliaryCommandResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SendAuxiliaryCommandResponse SendAuxiliaryCommandResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SendAuxiliaryCommandResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SendAuxiliaryCommand")
		return reply.Body.SendAuxiliaryCommandResponse, errors.Annotate(err, "reply")
	}
}
