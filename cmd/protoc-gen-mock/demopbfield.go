package main

import (
	"github.com/google/uuid"
	"github.com/nlm/protoc-gen-mock/demopb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func init() {
	//protomock.RegisterFieldValueMocker(demopbFieldValueMocker)
}

func demopbFieldValueMocker(field *protogen.Field) string {
	switch proto.GetExtension(field.Desc.Options().(*descriptorpb.FieldOptions), demopb.E_Type).(demopb.StandardFieldType) {
	case demopb.StandardFieldType_ip, demopb.StandardFieldType_ipv4:
		return "\"192.0.2.1\""
	case demopb.StandardFieldType_ipnet:
		return "\"192.0.2.0/24\""
	case demopb.StandardFieldType_uuid:
		return "\"" + uuid.NewString() + "\""
	case demopb.StandardFieldType_ipv6:
		return "\"2001:db8:9070:b0ef::/128\""
	case demopb.StandardFieldType_size:
		return "\"42\""
	default:
		return ""
	}
}
