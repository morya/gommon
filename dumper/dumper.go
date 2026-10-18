// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// this file is first built inside my personal project ex

package dumper

import (
	"bytes"
	"encoding/json"
)

func DumpJson(obj interface{}) string {
	data, _ := json.Marshal(obj)
	return string(data)
}

func DumpJsonIndent(obj interface{}) string {
	b := &bytes.Buffer{}
	enc := json.NewEncoder(b)
	enc.SetIndent("", "  ")
	enc.Encode(obj)
	return b.String()
}
