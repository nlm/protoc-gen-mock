package protomock

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const mockMaxDepth = 100
const mockRepeatedCount = 3

func newMessage(desc protoreflect.MessageDescriptor) protoreflect.ProtoMessage {
	mt, err := protoregistry.GlobalTypes.FindMessageByName(desc.FullName())
	if err != nil {
		// We should never panic, as if the field is present in the message,
		// it should be present in the registry
		panic(err)
	}
	return mt.New().Interface()
}

func mockList(field protoreflect.FieldDescriptor, lst protoreflect.List) {
	switch field.Kind() {
	// FIXME: other kinds
	case protoreflect.MessageKind:
		for i := 0; i < mockRepeatedCount; i++ {
			sm := lst.AppendMutable()
			mockMessage(sm.Message().Interface(), 0)
		}
	default:
		for i := 0; i < mockRepeatedCount; i++ {
			lst.Append(mockField(field))
		}
	}
}

func mockMessage(msg proto.Message, depth int) {
	fields := msg.ProtoReflect().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		// fmt.Println(field.FullName())
		if field.IsList() {
			lst := msg.ProtoReflect().Mutable(field).List()
			mockList(field, lst)
			continue
		}
		if field.IsMap() {
			fmt.Println(field.FullName())
			fmt.Println("FIXME")
			continue
		}
		if field.Kind() != protoreflect.MessageKind {
			msg.ProtoReflect().Set(field, mockField(field))
			continue
		}
		// protect from infinite recursion
		if depth >= mockMaxDepth {
			continue
		}
		sm := newMessage(field.Message())
		mockMessage(sm, depth+1)
		msg.ProtoReflect().Set(field, protoreflect.ValueOf(sm.ProtoReflect()))
	}
}
