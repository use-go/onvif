// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetNetworkInterfaces forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkInterfacesResponse.
func Call_GetNetworkInterfaces(ctx context.Context, dev *Device, request GetNetworkInterfaces) (GetNetworkInterfacesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkInterfacesResponse GetNetworkInterfacesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNetworkInterfacesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetNetworkInterfaces")
		return reply.Body.GetNetworkInterfacesResponse, errors.Annotate(err, "reply")
	}
}
