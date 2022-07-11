package main

import (
	"strings"

	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/compiler/protogen"
)

func init() {
	RegisterFieldValueMocker(nameBasedFieldValueMocker)
}

func genRandomName() string {
	return cases.Title(language.Und).String(strings.Replace(namegenerator.GetRandomName(), "-", " ", 1))
}

func genRandomEmail() string {
	return strings.ToLower(strings.Replace(namegenerator.GetRandomName(), "-", ".", 1)) + "@example.com"
}

func nameBasedFieldValueMocker(field *protogen.Field) string {
	switch strings.ToLower(string(field.Desc.Name())) {
	case "name":
		return "\"" + genRandomName() + "\""
	case "email":
		return "\"" + genRandomEmail() + "\""
	}
	return ""
}
