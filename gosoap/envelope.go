package gosoap

import (
	"encoding/xml"
)

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header  SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Header"`

	Headers []interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Fault"`

	Code   *SOAPFaultCode   `xml:",omitempty"`
	Reason *SOAPFaultReason `xml:",omitempty"`
}

// UnmarshalXML the response body
// https://github.com/faceterteam/onvif4go/blob/master/soap/types.go#L46
// https://play.golang.org/p/FRzdAFrXZ1
func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://www.w3.org/2003/05/soap-envelope" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

type SOAPFaultCode struct {
	Value   string         `xml:"Value"`
	Subcode *SOAPFaultCode `xml:"Subcode,omitempty"`
}

type SOAPFaultReason struct {
	Text string `xml:"Text"`
}

func NewSOAPEnvelope(content interface{}) *SOAPEnvelope {
	return &SOAPEnvelope{
		Body: SOAPBody{
			Content: content,
		},
	}
}
