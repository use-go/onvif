// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_SetNetworkDefaultGateway forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNetworkDefaultGatewayResponse.
func Call_SetNetworkDefaultGateway(ctx context.Context, dev *Device, request SetNetworkDefaultGateway) (SetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkDefaultGatewayResponse SetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNetworkDefaultGatewayResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "SetNetworkDefaultGateway")
		return reply.Body.SetNetworkDefaultGatewayResponse, errors.Annotate(err, "reply")
	}
}
