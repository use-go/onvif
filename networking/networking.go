package networking

import (
	"net/http"
	"bytes"
)

func SendSoap(endpoint, message string) (*http.Response, error) {
	httpClient := new(http.Client)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, err
	}

	/*if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusBadRequest {
		return "", errors.New("error: got HTTP response status " + strconv.Itoa(resp.StatusCode))
	}*/
	//b, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return resp, err
	//}
	//fmt.Println(endpoint)
	//fmt.Println(string(b))
	//log.Println(resp.StatusCode)

	return resp,nil
}
