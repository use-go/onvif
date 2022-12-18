// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/juju/errors"
	"github.com/use-go/onvif/networking"
)

// Call_GetCertificateInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificateInformationResponse.
func Call_GetCertificateInformation(ctx context.Context, dev *Device, request GetCertificateInformation) (GetCertificateInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificateInformationResponse GetCertificateInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCertificateInformationResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetCertificateInformation")
		return reply.Body.GetCertificateInformationResponse, errors.Annotate(err, "reply")
	}
}
