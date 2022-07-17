package main

/*
import (
	"strings"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/nlm/protoc-gen-mock/demopb"
	"github.com/nlm/protoc-gen-mock/pkg/_protomockstring"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func init() {
	protomockstring.RegisterFieldValueMocker(demopbNameBasedFieldValueMocker)
	protomockstring.RegisterFieldValueMocker(demopbAnnotationFieldValueMocker)
}

var Q = protomockstring.Q

func ParentName(field *protogen.Field) string {
	if field.Parent == nil {
		return ""
	}
	return string(field.Parent.Desc.Name())
}

const (
	StringValue = "google.protobuf.StringValue"
)

func demopbNameBasedFieldValueMocker(field *protogen.Field) string {
	switch strings.ToLower(string(field.Desc.Name())) {
	case "organization_id", "project_id":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.UUID())
		case protoreflect.MessageKind:
			if field.Desc.FullName() == StringValue {
				return Q(fake.UUID())
			}
		}
	case "created_at", "updated_at":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.Date())
		}
	}
	return ""
}

func demopbAnnotationFieldValueMocker(field *protogen.Field) string {
	switch proto.GetExtension(field.Desc.Options().(*descriptorpb.FieldOptions), demopb.E_Type).(demopb.StandardFieldType) {
	case demopb.StandardFieldType_ip, demopb.StandardFieldType_ipv4:
		return Q(fake.IPv4Address())
	case demopb.StandardFieldType_ipnet:
		return Q(fake.IPv4Address() + "/32")
	case demopb.StandardFieldType_uuid:
		return Q(fake.UUID())
	case demopb.StandardFieldType_ipv6:
		return Q(fake.IPv6Address())
	case demopb.StandardFieldType_size:
		return Q("42")
	default:
		return ""
	}
}
*/
