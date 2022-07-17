package pbutils

import "google.golang.org/protobuf/reflect/protoreflect"

func IsNumeralKind(kind protoreflect.Kind) bool {
	switch kind {
	case protoreflect.DoubleKind, protoreflect.FloatKind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Int32Kind,
		protoreflect.Sfixed64Kind, protoreflect.Int64Kind,
		protoreflect.Fixed32Kind, protoreflect.Uint32Kind,
		protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
		return true
	default:
		return false
	}
}
