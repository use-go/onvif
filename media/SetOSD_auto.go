// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/networking"
)

// Call_SetOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a SetOSDResponse.
func Call_SetOSD(ctx context.Context, dev *onvif.Device, request SetOSD) (SetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetOSDResponse SetOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetOSDResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetOSD")
		return reply.Body.SetOSDResponse, errors.Annotate(err, "reply")
	}
}
