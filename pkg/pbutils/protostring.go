package pbutils

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtoString(m proto.Message) string {
	bytes, err := protojson.MarshalOptions{
		UseProtoNames:   true,
		UseEnumNumbers:  true,
		EmitUnpopulated: true,
	}.Marshal(m)
	if err != nil {
		return ""
	}
	return string(bytes)
}
