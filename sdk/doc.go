package sdk

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func ReadAndParse(ctx context.Context, httpReply *http.Response, reply interface{}, tag string) error {
	if b, err := ioutil.ReadAll(httpReply.Body); err != nil {
		return err
	} else {
		err = xml.Unmarshal(b, reply)
		return err
	}
}
