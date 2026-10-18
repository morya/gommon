package b64

import (
	"bytes"
	"testing"
)

func TestB64Encode(t *testing.T) {
	data := []byte(`some-demo-string`)
	right := []byte(`c29tZS1kZW1vLXN0cmluZw==`)
	result := Base64Encode(data)
	t.Logf("v = %v", result)

	if bytes.Compare([]byte(result), right) != 0 {
		t.Errorf("compare failed")
	}
}
