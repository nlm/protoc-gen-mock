package protomock

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	numberLowValue      = 1000000
	numberHighValue     = 9999999
	loremIpsumWordCount = 5
)

func randomScalarValueMocker(field protoreflect.FieldDescriptor, fieldOptions proto.Message) any {
	switch field.Kind() {
	// Scalar Types
	// https://developers.google.com/protocol-buffers/docs/proto3#scalar
	case protoreflect.DoubleKind:
		return float64(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.FloatKind:
		return float32(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Int32Kind:
		return int32(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Int64Kind:
		return int64(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Uint32Kind:
		return uint32(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Uint64Kind:
		return uint64(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Sint32Kind:
		return int32(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Sint64Kind:
		return int64(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Fixed32Kind:
		return uint32(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Fixed64Kind:
		return uint64(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Sfixed32Kind:
		return int32(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.Sfixed64Kind:
		return int64(Faker().Number(numberLowValue, numberHighValue))
	case protoreflect.BoolKind:
		return bool(Faker().Bool())
	case protoreflect.StringKind:
		return string(Faker().LoremIpsumSentence(loremIpsumWordCount))
	case protoreflect.BytesKind:
		return []byte(Faker().LoremIpsumSentence(loremIpsumWordCount))
	// Enum Type
	case protoreflect.EnumKind:
		return field.DefaultEnumValue()
		// return field.Enum().Values().Get(int(Faker().Int64() % int64(field.Enum().Values().Len()))).Number()
	}
	return nil
}
