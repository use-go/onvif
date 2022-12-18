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

// Call_GetCertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificatesResponse.
func Call_GetCertificates(ctx context.Context, dev *onvif.Device, request GetCertificates) (GetCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesResponse GetCertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCertificatesResponse, errors.Annotate(err, "call")
	} else {
		err = networking.ReadAndParse(ctx, httpReply, &reply, "GetCertificates")
		return reply.Body.GetCertificatesResponse, errors.Annotate(err, "reply")
	}
}
