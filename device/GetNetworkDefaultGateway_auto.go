// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetNetworkDefaultGateway forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkDefaultGatewayResponse.
func Call_GetNetworkDefaultGateway(ctx context.Context, dev *Device, request GetNetworkDefaultGateway) (GetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkDefaultGatewayResponse GetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetNetworkDefaultGateway")
		return reply.Body.GetNetworkDefaultGatewayResponse, errors.Annotate(err, "reply")
	}
}
