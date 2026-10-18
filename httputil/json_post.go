package httputil

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/morya/gommon/gostring"
)

func JsonPost(url string, req interface{}, ack interface{}) error {
	body := bytes.NewReader(gostring.JsonEncode(req))
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Accept", "application/json")
	data, err := HttpPost(url, body, headers)
	if err != nil {
		return err
	}

	if ack != nil {
		err = json.Unmarshal(data, ack)
	}
	return err
}
