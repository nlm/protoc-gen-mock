package protomockstring

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func defaultFieldValueMocker(field *protogen.Field) string {
	return KindDefaultValue(field.Desc.Kind())
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
	if value := nameBasedFieldValueMocker(field); value != "" {
		return value
	}
	return defaultFieldValueMocker(field)
}
