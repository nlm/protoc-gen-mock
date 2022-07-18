package protomock

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ScalarValueMocker func(fieldDescriptor protoreflect.FieldDescriptor, fieldOptions proto.Message) any

var scalarValueMockers []ScalarValueMocker

func init() {
	registerScalarValueMocker(
		optionBasedScalarValueMocker,
		nameBasedScalarValueMocker,
		randomScalarValueMocker,
	)
}

func registerScalarValueMocker(fm ...ScalarValueMocker) {
	scalarValueMockers = append(scalarValueMockers, fm...)
}

func mockScalar(field protoreflect.FieldDescriptor, fieldOptions proto.Message) protoreflect.Value {
	for _, fm := range scalarValueMockers {
		if v := fm(field, fieldOptions); v != nil {
			return protoreflect.ValueOf(v)
		}
	}
	return field.Default()
}
