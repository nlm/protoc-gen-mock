package protomock

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FieldMocker func(protoreflect.FieldDescriptor) any

var fieldMockers []FieldMocker

func init() {
	registerFieldMocker(nameBasedFieldMocker, randomFieldMocker)
}

func registerFieldMocker(fm ...FieldMocker) {
	fieldMockers = append(fieldMockers, fm...)
}

func mockField(field protoreflect.FieldDescriptor) protoreflect.Value {
	for _, fm := range fieldMockers {
		if v := fm(field); v != nil {
			return protoreflect.ValueOf(v)
		}
	}
	return field.Default()
}
