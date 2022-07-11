package protomock

import (
	fake "github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func KindDefaultValue(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.BoolKind:
		return S([...]bool{false, true}[fake.Number(0, 1)])
	case protoreflect.StringKind:
		return Q(fake.LoremIpsumSentence(5))
	case protoreflect.BytesKind:
		return S("[]byte(\"" + fake.LoremIpsumSentence(5) + "\")")
	case protoreflect.EnumKind:
		return S(0)
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
		return S(0)
	case protoreflect.MessageKind:
		// FIXME
		return "nil"
	default:
		return ""
	}
}

func kindDefaultValueMocker(field *protogen.Field) string {
	return KindDefaultValue(field.Desc.Kind())
}
