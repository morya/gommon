package base62

import "testing"

func TestFormatInt(t *testing.T) {
	arr := []int64{
		1,
		10003333,
		99998888,
	}

	for _, v := range arr {
		t.Logf("%s", FormatInt(v))
	}
}
