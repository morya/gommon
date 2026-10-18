package httputil

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

func HttpGetTimeout(timeout time.Duration, url string, params url.Values, headers http.Header) ([]byte, error) {
	var c = &http.Client{
		Timeout: timeout,
	}
	if params != nil {
		url += "?" + params.Encode()
	}
	var req, _ = http.NewRequest("GET", url, nil)
	if headers != nil {
		req.Header = headers
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return body, err
}

func HttpGet(url string, params url.Values, headers http.Header) ([]byte, error) {
	return HttpGetTimeout(time.Second*5, url, params, headers)
}
