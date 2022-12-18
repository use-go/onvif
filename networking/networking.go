package networking

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"github.com/juju/errors"
	"io/ioutil"
	"net/http"
)

// SendSoap send soap message
func SendSoap(httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, errors.Annotate(err, "Post")
	}

	return resp, nil
}

func ReadAndParse(ctx context.Context, httpReply *http.Response, reply interface{}, tag string) error {
	if httpReply.StatusCode/100 != 2 {
		return fmt.Errorf("unexpected status %v instead of 2XX", httpReply.StatusCode)
	}
	if b, err := ioutil.ReadAll(httpReply.Body); err != nil {
		return errors.Annotate(err, "read")
	} else {
		err = xml.Unmarshal(b, reply)
		return errors.Annotate(err, "decode")
	}
}
