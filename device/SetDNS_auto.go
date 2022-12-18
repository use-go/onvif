// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_SetDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDNSResponse.
func Call_SetDNS(ctx context.Context, dev *Device, request SetDNS) (SetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDNSResponse SetDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetDNSResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetDNS")
		return reply.Body.SetDNSResponse, errors.Annotate(err, "reply")
	}
}
