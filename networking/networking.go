package networking

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/juju/errors"
)

// SendSoap send soap message
func SendSoap(httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	fmt.Printf("SendSoap message:%s\n", message)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, errors.Annotate(err, "Post")
	}

	return resp, nil
}
