package networking

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"github.com/pkg/errors"
	"strconv"
)

func SendSoap(endpoint, message string) (string, error) {
	httpClient := new(http.Client)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error: got HTTP response status " + strconv.Itoa(resp.StatusCode))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//log.Println(resp.StatusCode)

	return string(b),nil
}
