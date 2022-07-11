package protomock

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Q(str string) string {
	return "\"" + str + "\""
}

func nameBasedFieldValueMocker(field *protogen.Field) string {
	switch strings.ToLower(string(field.Desc.Name())) {
	case "id":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(deterministicRandom.RandUUID())
		case protoreflect.Fixed32Kind:
			return Q("42")
		}
	case "name":
		return Q(deterministicRandom.RandName())
	case "email":
		return Q(deterministicRandom.RandEmail())
	}
	return ""
}
