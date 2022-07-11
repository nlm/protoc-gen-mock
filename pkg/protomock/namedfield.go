package protomock

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/google/uuid"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	int32IdMax = 100000000
)

func genRandomName() string {
	return cases.Title(language.Und).String(strings.Replace(namegenerator.GetRandomName(), "-", " ", 1))
}

func genRandomEmail() string {
	return strings.ToLower(strings.Replace(namegenerator.GetRandomName(), "-", ".", 1)) + "@example.com"
}

func genRandomUUID() string {
	return uuid.NewString()
}

func nameBasedFieldValueMocker(field *protogen.Field) string {
	switch strings.ToLower(string(field.Desc.Name())) {
	case "id":
		switch field.Desc.Kind() {
		case protoreflect.StringKind:
			return "\"" + genRandomUUID() + "\""
		case protoreflect.Fixed32Kind:
			return "\"" + fmt.Sprint(rand.Int31n(int32IdMax)) + "\""
		}
	case "name":
		return "\"" + genRandomName() + "\""
	case "email":
		return "\"" + genRandomEmail() + "\""
	}
	return ""
}
