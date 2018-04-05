package networking

import (
	"net/http"
	"bytes"
	"io/ioutil"
)

func SendSoap(endpoint, message string) string {
	httpClient := new(http.Client)

	resp, _ := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))

	b, _ := ioutil.ReadAll(resp.Body)

	//log.Println(string(b))

	return string(b)
}
