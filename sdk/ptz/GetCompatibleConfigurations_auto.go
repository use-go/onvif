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

// Call_GetCompatibleConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleConfigurationsResponse.
func Call_GetCompatibleConfigurations(ctx context.Context, dev *onvif.Device, request ptz.GetCompatibleConfigurations) (ptz.GetCompatibleConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleConfigurationsResponse ptz.GetCompatibleConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleConfigurationsResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleConfigurations")
		return reply.Body.GetCompatibleConfigurationsResponse, err
	}
}
