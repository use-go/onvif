package networking

import (
	"net/http"
	"bytes"
	"io/ioutil"
)

func SendSoap(endpoint, message string) (string, error) {
	httpClient := new(http.Client)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b),nil
}
