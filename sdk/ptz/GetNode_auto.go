// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/ptz"
)

// Call_GetNode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNodeResponse.
func Call_GetNode(ctx context.Context, dev *onvif.Device, request ptz.GetNode) (ptz.GetNodeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNodeResponse ptz.GetNodeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNodeResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetNode")
		return reply.Body.GetNodeResponse, err
	}
}
