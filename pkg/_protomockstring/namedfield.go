package protomockstring

import (
	"strings"

	fake "github.com/brianvoe/gofakeit/v6"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var sepRemover = strings.NewReplacer("-", "", "_", "", ".", "")

func nameBasedFieldValueMocker(field *protogen.Field) string {
	switch strings.ToLower(sepRemover.Replace(string(field.Desc.Name()))) {
	case "id", "uid", "uuid":
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
	case "firstname":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.FirstName())
		}
	case "lastname":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.LastName())
		}
	case "name", "fullname":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.Name())
		}
	case "mail", "email":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.Email())
		}
	case "address", "streetaddress":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.Address())
		}
	case "mac", "macaddress", "macaddr":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.MacAddress())
		}
	case "description":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return Q(fake.LoremIpsumSentence(5))
		}
	}
	return ""
}
