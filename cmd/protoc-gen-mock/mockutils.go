package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtoKindDefaultValue(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.BoolKind:
		return fmt.Sprint(false)
	case protoreflect.StringKind:
		return "\"string\""
	case protoreflect.BytesKind:
		return "[]byte(\"bytes\")"
	case protoreflect.EnumKind:
		return "0"
	case protoreflect.DoubleKind,
		protoreflect.Fixed32Kind,
		protoreflect.Fixed64Kind,
		protoreflect.FloatKind,
		protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sint64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind:
		return "0"
	case protoreflect.MessageKind:
		// FIXME
		return "nil"
	default:
		return ""
	}
}

func defaultFieldValueMocker(field *protogen.Field) string {
	return ProtoKindDefaultValue(field.Desc.Kind())
}

type FieldValueMocker func(*protogen.Field) string

var fieldValueMockers []FieldValueMocker

func RegisterFieldValueMocker(fvm FieldValueMocker) {
	fieldValueMockers = append(fieldValueMockers, fvm)
}

func MockFieldValue(field *protogen.Field) string {
	for _, fvm := range fieldValueMockers {
		if value := fvm(field); value != "" {
			return value
		}
	}
	return defaultFieldValueMocker(field)
}
