// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/device"
	"github.com/use-go/onvif/networking"
)

// Call_GetNodes forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNodesResponse.
func Call_GetNodes(ctx context.Context, dev *device.Device, request GetNodes) (GetNodesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodesResponse GetNodesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNodesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetNodes")
		return reply.Body.GetNodesResponse, errors.Annotate(err, "reply")
	}
}
