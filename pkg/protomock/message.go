package protomock

import (
	"github.com/nlm/protoc-gen-mock/pkg/pb/mockpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const mockMaxDepth = 100
const mockRepeatedCount = 3

func getRepetitions(field protoreflect.FieldDescriptor) int {
	v := proto.GetExtension(field.Options(), mockpb.E_Items).(uint32)
	if v > 0 {
		return int(v)
	}
	return mockRepeatedCount
}

func newMessage(desc protoreflect.MessageDescriptor) protoreflect.ProtoMessage {
	mt, err := protoregistry.GlobalTypes.FindMessageByName(desc.FullName())
	if err != nil {
		// We should never panic, as if the field is present in the message,
		// it should be present in the registry
		panic(err)
	}
	return mt.New().Interface()
}

// mockList mocks a list field
func mockList(msg proto.Message, field protoreflect.FieldDescriptor, depth int) {
	lst := msg.ProtoReflect().Mutable(field).List()
	// check tags for repetition
	switch field.Kind() {
	// FIXME: other kinds
	case protoreflect.MessageKind:
		for i := 0; i < getRepetitions(field); i++ {
			sm := lst.AppendMutable()
			mockMessage(sm.Message().Interface(), 0)
		}
	default:
		for i := 0; i < getRepetitions(field); i++ {
			lst.Append(mockScalar(field))
		}
	}
}

// mockMap mocks a map field
func mockMap(msg proto.Message, field protoreflect.FieldDescriptor, depth int) {
	mp := msg.ProtoReflect().Mutable(field).Map()
	for i := 0; i < getRepetitions(field); i++ {
		// Key
		var mapKey protoreflect.MapKey
		switch field.MapKey().Kind() {
		case protoreflect.MessageKind:
			// actually cannot happen
			m := newMessage(field.Message())
			mockMessage(m, depth)
			mapKey = protoreflect.MapKey(protoreflect.ValueOfMessage(m.ProtoReflect()))
		default:
			mapKey = protoreflect.MapKey(mockScalar(field.MapKey()))
		}
		// Value
		switch field.MapValue().Kind() {
		case protoreflect.MessageKind:
			mapValue := mp.Mutable(mapKey)
			mockMessage(mapValue.Message().Interface(), depth)
			// TODO: list / map ?
		default:
			mapValue := mockScalar(field.MapValue())
			mp.Set(mapKey, mapValue)
		}
	}
}

// mockUnary mocks an unary message type
func mockUnary(msg proto.Message, field protoreflect.FieldDescriptor, depth int) {
	switch field.Kind() {
	case protoreflect.MessageKind:
		sm := newMessage(field.Message())
		mockMessage(sm, depth)
		msg.ProtoReflect().Set(field, protoreflect.ValueOf(sm.ProtoReflect()))
	default:
		msg.ProtoReflect().Set(field, mockScalar(field))
	}
}

// mockField mocks a field
func mockField(msg proto.Message, field protoreflect.FieldDescriptor, depth int) {
	switch {
	case field.IsList():
		mockList(msg, field, depth)
	case field.IsMap():
		mockMap(msg, field, depth)
	default:
		mockUnary(msg, field, depth)
	}
}

// mockMessage mocks a Message field
func mockMessage(msg proto.Message, depth int) {
	if depth >= mockMaxDepth {
		return
	}
	fields := msg.ProtoReflect().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		mockField(msg, field, depth+1)
	}
}
