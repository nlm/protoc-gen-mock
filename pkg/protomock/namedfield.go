package protomock

import (
	"strings"

	fake "github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func nameBasedFieldValueMocker(field *protogen.Field) string {
	switch strings.ToLower(string(field.Desc.Name())) {
	case "id":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.UUID())
		case protoreflect.Fixed32Kind,
			protoreflect.Fixed64Kind,
			protoreflect.Int32Kind,
			protoreflect.Int64Kind,
			protoreflect.Sfixed32Kind,
			protoreflect.Sfixed64Kind,
			protoreflect.Sint32Kind,
			protoreflect.Sint64Kind,
			protoreflect.Uint32Kind,
			protoreflect.Uint64Kind:
			return S(fake.Number(1000000, 9999999))
		}
	case "first_name", "firstname", "first-name":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.FirstName())
		}
	case "last_name", "lastname", "last-name":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.LastName())
		}
	case "name":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.Name())
		}
	case "email", "mail", "e-mail":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.Email())
		}
	case "address":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.Address())
		}
	case "mac",
		"macaddress", "mac-address", "mac_address",
		"macaddr", "mac-addr", "mac_addr":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.MacAddress())
		}
	}
	return ""
}
