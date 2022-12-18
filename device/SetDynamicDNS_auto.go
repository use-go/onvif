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

// Call_SetDynamicDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDynamicDNSResponse.
func Call_SetDynamicDNS(ctx context.Context, dev *onvif.Device, request SetDynamicDNS) (SetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDynamicDNSResponse SetDynamicDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetDynamicDNSResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetDynamicDNS")
		return reply.Body.SetDynamicDNSResponse, errors.Annotate(err, "reply")
	}
}
