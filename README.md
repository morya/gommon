# gommon

A collection of small, independent Go utility packages, merged from the util layer of
[dreamans/syncd](https://github.com/dreamans/syncd) and other day-to-day code.

[![CI](https://img.shields.io/github/actions/workflow/status/morya/gommon/ci.yml?branch=main)](https://github.com/morya/gommon/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/morya/gommon)](https://goreportcard.com/report/github.com/morya/gommon)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Overview

gommon is a utility collection, not a framework. Each package does one narrow job and
can be imported on its own, so you only pay for what you use. Most packages are thin
wrappers over the standard library or a single well-known dependency; none of them
pulls in the rest of the module. Requires Go 1.24 or later.

## Install

```shell
go get github.com/morya/gommon
```

Then import the package you need:

```go
import "github.com/morya/gommon/base62"
```

## Packages

| Package | Description | Notes |
|---|---|---|
| [b64](b64/) | Base64 standard and URL encode/decode helpers | Wraps `encoding/base64` |
| [base62](base62/) | Radix-62 encoding/decoding for bytes and integers | Custom alphabets via `NewEncoding` |
| [config](config/) | Config-friendly types | `Duration` marshals as `"1m"` in JSON/YAML |
| [datetime](datetime/) | Common `time.Duration` constants | `Day`, `Week`, `ThirtyDays` |
| [dumper](dumper/) | One-line JSON dumps for debugging | `DumpJson`, `DumpJsonIndent` |
| [eagle_validator](eagle_validator/) | Struct validation with Chinese error messages | Built on `go-playground/validator/v10` |
| [goslice](goslice/) | Slice contains/intersect/complement helpers | |
| [gostring](gostring/) | String conversion, join/split, random and UTF-8 helpers | |
| [httputil](httputil/) | One-shot GET/POST/JSON requests returning the response body | 5s default timeout |
| [humanize](humanize/) | Human-readable byte sizes and a byte-size parser | SI (`Bytes`) and IEC (`IBytes`) |
| [midnight](midnight/) | Midnight (start of day) for now or any timestamp | |
| [mytrace](mytrace/) | Current/caller function names from `runtime` | Debug logging aid |
| [pbjson](pbjson/) | Compact JSON for protobuf messages | Emits unpopulated fields |
| [string_array](string_array/) | `[]string` implementing `sql.Scanner` / `driver.Valuer` | JSONB columns, also a `flag.Value` |
| [strmd5](strmd5/) | MD5 hex digest of a string, full or 16-char short form | |
| [wait_group_wrapper](wait_group_wrapper/) | `sync.WaitGroup` with `Wrap(func())` running Add/Done for you | |

## Usage

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/morya/gommon/base62"
	"github.com/morya/gommon/httputil"
	"github.com/morya/gommon/wait_group_wrapper"
)

func main() {
	// base62: compact radix-62 encoding for bytes and integers
	enc := base62.EncodeToString([]byte("gommon"))
	num, _ := base62.ParseUint(base62.FormatUint(42))
	fmt.Println(enc, num)

	// httputil: one-shot requests, the body comes back as bytes
	params := url.Values{"q": []string{"golang"}}
	body, err := httputil.HttpGet("https://example.com/api", params, nil)
	if err == nil {
		fmt.Println(len(body))
	}

	var ack map[string]any
	if err := httputil.JsonPost("https://example.com/api", map[string]any{"id": 7}, &ack); err == nil {
		fmt.Println(ack)
	}

	// wait_group_wrapper: Add/Done handled for you
	var wg wait_group_wrapper.WaitGroupWrapper
	wg.Wrap(func() { fmt.Println("done in goroutine") })
	wg.Wait()
}
```

`config.Duration` reads durations written as strings in config files:

```go
type Server struct {
	ReadTimeout config.Duration `json:"read_timeout" yaml:"read_timeout"`
}

var s Server
_ = json.Unmarshal([]byte(`{"read_timeout": "30s"}`), &s)
fmt.Println(s.ReadTimeout.Duration) // 30s
```

## Development

```shell
go build ./...
go vet ./...
go test ./...
```

CI (`.github/workflows/ci.yml`) runs build, vet, and test on every push and pull request.

## License

MIT, see [LICENSE](LICENSE). Copyright (c) morya.

## Acknowledgments

- Most packages originate from the `util` directory of [dreamans/syncd](https://github.com/dreamans/syncd) (MIT).
- The `humanize` package contains code copied from [dustin/go-humanize](https://github.com/dustin/go-humanize) (MIT).
