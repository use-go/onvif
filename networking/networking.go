package networking

import (
	"bytes"
	"net/http"
	"time"
)

// SendSoap send soap message
func SendSoap(endpoint, message string) (*http.Response, error) {
	httpClient := new(http.Client)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// SendSoapWithTimeout send soap message with timeOut
func SendSoapWithTimeout(endpoint string, message []byte, timeout time.Duration) (*http.Response, error) {
	httpClient := &http.Client{
		Timeout: timeout,
	}

	return httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewReader(message))
}
