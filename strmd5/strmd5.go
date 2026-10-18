package strmd5

import (
	"crypto/md5"
	"encoding/hex"
)

func StrMd5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}

func StrMd5Short(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))[8:24]
}
