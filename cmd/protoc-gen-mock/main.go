package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func KindDefaultValue(kind protoreflect.Kind) string {
	switch kind {
	case protoreflect.BoolKind:
		return fmt.Sprint(false)
	case protoreflect.StringKind:
		return "\"string\""
	case protoreflect.BytesKind:
		return "[]byte(\"bytes\")"
	case protoreflect.EnumKind:
		return "0"
	case protoreflect.DoubleKind,
		protoreflect.Fixed32Kind,
		protoreflect.Fixed64Kind,
		protoreflect.FloatKind,
		protoreflect.Int32Kind,
		protoreflect.Int64Kind,
		protoreflect.Sfixed32Kind,
		protoreflect.Sfixed64Kind,
		protoreflect.Sint32Kind,
		protoreflect.Sint64Kind,
		protoreflect.Uint32Kind,
		protoreflect.Uint64Kind:
		return "0"
	case protoreflect.MessageKind:
		// FIXME
		return "nil"
	default:
		return ""
	}
}

func MockFieldValue(field *protogen.Field) string {
	return KindDefaultValue(field.Desc.Kind())
}

func Generate(gen *protogen.Plugin) error {
	log.Print("----- START PLUGIN -----")
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		fileName := file.GeneratedFilenamePrefix + ".mock.go"
		genFile := gen.NewGeneratedFile(fileName, "")
		log.Print("----- BEGIN FILE ", file.Desc.Path(), " -----")

		// packages
		genFile.P("package ", file.GoPackageName)
		genFile.P("import \"google.golang.org/grpc\"")
		genFile.P("import \"context\"")

		// Services
		for _, s := range file.Services {
			baseServerName := s.GoName + "Server"
			mockServerName := "Mock" + baseServerName
			mockResponsesName := mockServerName + "Responses"
			mockErrorsName := mockServerName + "Errors"

			// type MockXServerResponses
			genFile.P("type ", mockResponsesName, " struct {")
			for _, m := range s.Methods {
				genFile.P(m.GoName, " *", m.Output.GoIdent.GoName)
			}
			genFile.P("}")

			// type MockXServerErrors
			genFile.P("type ", mockErrorsName, " struct {")
			for _, m := range s.Methods {
				genFile.P(m.GoName, " error")
			}
			genFile.P("}")

			// type MockXServer
			genFile.P("type ", mockServerName, " struct {")
			genFile.P("Unimplemented", baseServerName)
			genFile.P("responses ", mockResponsesName)
			genFile.P("errors ", mockErrorsName)
			genFile.P("}")

			for _, m := range s.Methods {
				genFile.P("// func (ms *", mockServerName, ") Register", m.GoName, "Response(response *", m.Output.GoIdent.GoName, ") {")
				genFile.P("// ms.responses.", m.GoName, " = response")
				genFile.P("// }")
			}

			genFile.P("func (ms *", mockServerName, ") RegisterMockResponse(method string, response any) {")
			genFile.P("switch method {")
			for _, m := range s.Methods {
				genFile.P("case \"", m.GoName, "\":")
				genFile.P("switch r := response.(type) {")
				genFile.P("case error:")
				genFile.P("ms.errors.", m.GoName, " = r")
				genFile.P("case *", m.Output.GoIdent.GoName, ":")
				genFile.P("ms.responses.", m.GoName, " = r")
				genFile.P("default:")
				genFile.P("panic(\"wrong argument type for this method\")")
				genFile.P("}")
			}
			genFile.P("default:")
			genFile.P("panic(\"unknown method name: \" + method)")
			genFile.P("}")
			genFile.P("}")

			// Methods
			for _, m := range s.Methods {
				genFile.P("func (ms *", mockServerName, ") ", m.GoName, "(ctx context.Context, req *", m.Input.GoIdent.GoName, ") (*", m.Output.GoIdent.GoName, ", error) {")

				// check registered errors
				genFile.P("if ms.errors.", m.GoName, " != nil {")
				genFile.P("return nil, ms.errors.", m.GoName)
				genFile.P("}")

				// check registered response
				genFile.P("if ms.responses.", m.GoName, " != nil {")
				genFile.P("return ms.responses.", m.GoName, ", nil")
				genFile.P("}")

				// return defaut value
				genFile.P("return &", m.Output.GoIdent.GoName, "{")
				for _, f := range m.Output.Fields {
					o := f.Desc.Options().ProtoReflect().GetUnknown()
					genFile.P("// Options: ", o, " ", o.IsValid())
					genFile.P(f.GoName, ": ", MockFieldValue(f), ",")
				}
				genFile.P("}, nil")
				genFile.P("}")
			}

			// func RegisterXServer()
			genFile.P("func Register", mockServerName, "(s grpc.ServiceRegistrar) (*", mockServerName, ") {")
			genFile.P("ms := &", mockServerName, "{}")
			genFile.P("Register", baseServerName, "(s, ms)")
			genFile.P("return ms")
			genFile.P("}")
		}

		// genFile.P("import \"fmt\"")
		// for _, message := range file.Messages {
		// 	genFile.P("func (*", message.GoIdent.GoName, ") HelloWorld() {")
		// 	genFile.P("fmt.Println(\"Hello, ", message.GoIdent.GoName, "!\")")
		// 	genFile.P("}")
		// }
		log.Print("----- END FILE ", file.Desc.Path(), " -----")
	}
	log.Print("----- END PLUGIN -----")
	return nil
}

func main() {
	//log.SetOutput(os.Stderr)
	log.SetOutput(ioutil.Discard)
	protogen.Options{}.Run(Generate)
}
