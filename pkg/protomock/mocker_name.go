package protomock

import (
	"strings"

	"github.com/nlm/protoc-gen-mock/pkg/pbutils"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var sepRemover = strings.NewReplacer("-", "", "_", "", ".", "")

func normalizeFieldName(name protoreflect.Name) string {
	return strings.ToLower(sepRemover.Replace(string(name)))
}

func convertNumeral[T float32 | float64 | int | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](value T, kind protoreflect.Kind) any {
	switch kind {
	case protoreflect.DoubleKind:
		return float64(value)
	case protoreflect.FloatKind:
		return float32(value)
	case protoreflect.Sfixed32Kind, protoreflect.Int32Kind:
		return int32(value)
	case protoreflect.Sfixed64Kind, protoreflect.Int64Kind:
		return int64(value)
	case protoreflect.Fixed32Kind, protoreflect.Uint32Kind:
		return uint32(value)
	case protoreflect.Fixed64Kind, protoreflect.Uint64Kind:
		return uint64(value)
	}
	panic("invalid conversion")
}

func nameBasedFieldMocker(field protoreflect.FieldDescriptor) any {
	switch normalizeFieldName(field.Name()) {
	case "id", "uid", "uuid":
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().UUID()
		}
	case "firstname":
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().FirstName()
		}
	case "lastname":
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().LastName()
		}
	case "name", "fullname":
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Name()
		}
	case "mail", "email":
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Email()
		}
	case "address", "streetaddress":
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().Address()
		}
	case "mac", "macaddress", "macaddr":
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().MacAddress()
		}
	case "description":
		switch field.Kind() {
		case protoreflect.StringKind:
			return Faker().LoremIpsumSentence(5)
		}
	case "totalcount", "count":
		if pbutils.IsNumeralKind(field.Kind()) {
			return convertNumeral(mockRepeatedCount, field.Kind())
		}
	}
	return nil
}