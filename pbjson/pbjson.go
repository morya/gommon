package pbjson

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var opt = protojson.MarshalOptions{Multiline: false, EmitUnpopulated: true}

func PBJsonToStr(m proto.Message) string {
	return opt.Format(m)
}

func PBJsonToBytes(m proto.Message) []byte {
	data, _ := opt.Marshal(m)
	return data
}
