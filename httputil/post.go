package httputil

import (
	"io"
	"net/http"
	"time"
)

func HttpPost(url string, body io.Reader, headers http.Header) ([]byte, error) {
	return HttpPostWithTimeout(time.Second*5, url, body, headers)
}

func HttpPostWithTimeout(timeout time.Duration, url string, body io.Reader, headers http.Header) ([]byte, error) {
	var c = &http.Client{
		Timeout: timeout,
	}
	var req, _ = http.NewRequest("POST", url, body)
	if headers != nil {
		req.Header = headers
	}

	resp, err := c.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	return data, err
}
