package protomock

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

const populateMaxDepth = 100

// populate is the recurstion-protecting implementation of the Populate function
func populate(msg proto.Message, depth int) proto.Message {
	for i := 0; i < msg.ProtoReflect().Descriptor().Fields().Len(); i++ {
		field := msg.ProtoReflect().Descriptor().Fields().Get(i)
		if field.Kind() != protoreflect.MessageKind {
			continue
		}
		// Lookup is needed, because the descriptor does not directly link the concrete type
		mt, err := protoregistry.GlobalTypes.FindMessageByName(field.Message().FullName())
		if err != nil {
			// We should never panic, as if the field is present in the message,
			// it should be present in the registry
			panic(err)
		}
		// protect from infinite recursion
		if depth < populateMaxDepth {
			m := populate(mt.New().Interface(), depth+1).ProtoReflect()
			msg.ProtoReflect().Set(field, protoreflect.ValueOf(m))
		}
	}
	return msg
}

// Populate will recursively populate all the Message fields of a Message
func Populate(msg proto.Message) proto.Message {
	return populate(msg, 0)
}
