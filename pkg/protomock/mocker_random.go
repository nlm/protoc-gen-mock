package protomock

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	numberLowValue      = 1000000
	numberMaxValue      = 9999999
	loremIpsumWordCount = 5
)

func randomFieldMocker(field protoreflect.FieldDescriptor) any {
	switch field.Kind() {
	// Scalar Types
	// https://developers.google.com/protocol-buffers/docs/proto3#scalar
	case protoreflect.DoubleKind:
		return float64(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.FloatKind:
		return float32(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Int32Kind:
		return int32(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Int64Kind:
		return int64(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Uint32Kind:
		return uint32(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Uint64Kind:
		return uint64(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Sint32Kind:
		return int32(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Sint64Kind:
		return int64(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Fixed32Kind:
		return uint32(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Fixed64Kind:
		return uint64(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Sfixed32Kind:
		return int32(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.Sfixed64Kind:
		return int64(Faker().Number(numberLowValue, numberMaxValue))
	case protoreflect.BoolKind:
		return [...]bool{false, true}[Faker().Number(0, 1)]
	case protoreflect.StringKind:
		return string(Faker().LoremIpsumSentence(loremIpsumWordCount))
	case protoreflect.BytesKind:
		return []byte(Faker().LoremIpsumSentence(loremIpsumWordCount))
	// Message Type
	case protoreflect.MessageKind:
		return field.Default()
	// Enum Type
	case protoreflect.EnumKind:
		// FIXME
		return field.DefaultEnumValue()

	}
	// panic(fmt.Sprintf("unhandled kind: %v", field.Kind().GoString()))
	return nil
}
