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

// Call_GetSystemSupportInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemSupportInformationResponse.
func Call_GetSystemSupportInformation(ctx context.Context, dev *onvif.Device, request GetSystemSupportInformation) (GetSystemSupportInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemSupportInformationResponse GetSystemSupportInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemSupportInformationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetSystemSupportInformation")
		return reply.Body.GetSystemSupportInformationResponse, errors.Annotate(err, "reply")
	}
}
