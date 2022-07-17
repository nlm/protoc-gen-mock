package protomock

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ScalarValueMocker func(protoreflect.FieldDescriptor) any

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

func mockScalar(field protoreflect.FieldDescriptor) protoreflect.Value {
	for _, fm := range scalarValueMockers {
		if v := fm(field); v != nil {
			return protoreflect.ValueOf(v)
		}
	}
	return field.Default()
}
