package protomock

import (
	"github.com/nlm/protoc-gen-mock/pkg/pb/mockpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func optionBasedScalarValueMocker(field protoreflect.FieldDescriptor) any {
	// TODO: handle field.Name == "key" && field.ContainingMessage().IsMapEntry()
	switch proto.GetExtension(field.Options(), mockpb.E_Type).(mockpb.MockFieldType) {
	case mockpb.MockFieldType_ip, mockpb.MockFieldType_ipv4:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().IPv4Address()
		}
	case mockpb.MockFieldType_ipv6:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().IPv6Address()
		}
	case mockpb.MockFieldType_mac_address:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().MacAddress()
		}
	case mockpb.MockFieldType_uuid:
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().UUID()
		}
	}
	return nil
}
