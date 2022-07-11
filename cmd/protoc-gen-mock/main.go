package main

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/compiler/protogen"
)

var generatedErrors = []ErrorDecl{
	{"ErrWrongArgType", "wrong argument type for this method"},
	{"ErrUnknownMethod", "unknown method name"},
	{"ErrEmptyResponse", "empty response to register"},
}

var generatedImports = []ImportDecl{
	{"context", "context"},
	{"errors", "errors"},
	{"grpc", "google.golang.org/grpc"},
	{"protojson", "google.golang.org/protobuf/encoding/protojson"},
	{"status", "google.golang.org/genproto/googleapis/rpc/status"},
	{"spb", "google.golang.org/grpc/status"},
}

func genHeader(file *protogen.File, genFile *protogen.GeneratedFile) error {
	genFile.P("// Code generated by protoc-gen-mock. DO NOT EDIT.")
	genFile.P("// source: ", *file.Proto.Name)
	genFile.P("")
	genFile.P("package ", file.GoPackageName)
	return nil
}

func genImports(file *protogen.File, genFile *protogen.GeneratedFile) error {
	genFile.P("import (")
	for _, gi := range generatedImports {
		genFile.P(gi)
	}
	genFile.P(")")
	return nil
}

func genErrors(file *protogen.File, genFile *protogen.GeneratedFile) error {
	genFile.P("var (")
	for _, ge := range generatedErrors {
		genFile.P(ge)
	}
	genFile.P(")")
	return nil
}

func Generate(gen *protogen.Plugin) error {
	log.Print("----- BEGIN PLUGIN -----")
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}
		fileName := file.GeneratedFilenamePrefix + ".mock.go"
		genFile := gen.NewGeneratedFile(fileName, file.GoImportPath)
		log.Print("----- BEGIN FILE ", file.Desc.Path(), " -----")

		generators := []GenFunc{
			genHeader,
			genImports,
			genErrors,
		}
		for _, f := range generators {
			if err := f(file, genFile); err != nil {
				return err
			}
			genFile.P("")
		}

		// Services
		for _, s := range file.Services {
			baseServerName := s.GoName + "Server"
			mockServerName := "Mock" + baseServerName

			// type MockXServer
			genFile.P("type ", mockServerName, " struct {")
			genFile.P("Unimplemented", baseServerName)
			// contents
			genFile.P("contents struct {")
			for _, m := range s.Methods {
				genFile.P(m.GoName, " *", genFile.QualifiedGoIdent(m.Output.GoIdent))
			}
			genFile.P("}")
			// errors
			genFile.P("errors struct {")
			for _, m := range s.Methods {
				genFile.P(m.GoName, " error")
			}
			genFile.P("}")
			// callbacks
			genFile.P("callbacks struct {")
			for _, m := range s.Methods {
				genFile.P(m.GoName, " func(*", mockServerName, ")")
			}
			genFile.P("}")
			genFile.P("}")

			// RegisterMockResponse
			genFile.P("// RegisterMockResponse registers a response that is return at method invocation.")
			genFile.P("func (ms *", mockServerName, ") RegisterMockResponse(method string, response any) error {")
			genFile.P("switch method {")
			for _, m := range s.Methods {
				genFile.P("case \"", m.GoName, "\":")
				genFile.P("switch r := response.(type) {")
				genFile.P("case error:")
				genFile.P("ms.errors.", m.GoName, " = r")
				genFile.P("case *", genFile.QualifiedGoIdent(m.Output.GoIdent), ":")
				genFile.P("ms.contents.", m.GoName, " = r")
				genFile.P("default:")
				genFile.P("return ErrWrongArgType")
				genFile.P("}")
			}
			genFile.P("default:")
			genFile.P("return ErrUnknownMethod")
			genFile.P("}")
			genFile.P("return nil")
			genFile.P("}")

			// RegisterMockCallback
			genFile.P("// RegisterMockCallback registers a callback that is called after method invocation.")
			genFile.P("func (ms *", mockServerName, ") RegisterMockCallback(method string, callback func(*", mockServerName, ")) error {")
			genFile.P("switch method {")
			for _, m := range s.Methods {
				genFile.P("case \"", m.GoName, "\":")
				genFile.P("ms.callbacks.", m.GoName, " = callback")
			}
			genFile.P("default:")
			genFile.P("return ErrUnknownMethod")
			genFile.P("}")
			genFile.P("return nil")
			genFile.P("}")

			// RegisterJSONMockContent
			genFile.P("// RegisterJSONMockContent registers a JSON string as a Mock content,")
			genFile.P("// making sure that the format is respected")
			genFile.P("func (ms *", mockServerName, ") RegisterJSONMockContent(method string, payload []byte) error {")
			genFile.P("switch method {")
			for _, m := range s.Methods {
				genFile.P("case \"", m.GoName, "\":")
				genFile.P("var content = new(", genFile.QualifiedGoIdent(m.Output.GoIdent), ")")
				genFile.P("if err := protojson.Unmarshal(payload, content); err != nil {")
				genFile.P("return err")
				genFile.P("}")
				genFile.P("ms.contents.", m.GoName, " = content")
			}
			genFile.P("default:")
			genFile.P("return ErrUnknownMethod")
			genFile.P("}")
			genFile.P("return nil")
			genFile.P("}")

			// RegisterJSONMockStatus
			genFile.P("// RegisterJSONMockStatus registers a JSON string as a Mock status,")
			genFile.P("// making sure that the format is respected")
			genFile.P("func (ms *", mockServerName, ") RegisterJSONMockStatus(method string, payload []byte) error {")
			genFile.P("switch method {")
			for _, m := range s.Methods {
				genFile.P("case \"", m.GoName, "\":")
				genFile.P("var sta = new(status.Status)")
				genFile.P("if err := protojson.Unmarshal(payload, sta); err != nil {")
				genFile.P("return err")
				genFile.P("}")
				genFile.P("ms.errors.", m.GoName, " = spb.ErrorProto(sta)")
			}
			genFile.P("default:")
			genFile.P("return ErrUnknownMethod")
			genFile.P("}")
			genFile.P("return nil")
			genFile.P("}")

			// Mocked Methods
			for _, m := range s.Methods {
				genFile.P("func (ms *", mockServerName, ") ", m.GoName, "(",
					"ctx context.Context,",
					"req *", genFile.QualifiedGoIdent(m.Input.GoIdent), ")",
					"(*", genFile.QualifiedGoIdent(m.Output.GoIdent), ", error) {")

				// defer callback if present
				genFile.P("if ms.callbacks.", m.GoName, " != nil {")
				genFile.P("defer ms.callbacks.", m.GoName, "(ms)")
				genFile.P("}")

				// check registered errors
				genFile.P("if ms.errors.", m.GoName, " != nil {")
				genFile.P("return nil, ms.errors.", m.GoName)
				genFile.P("}")

				// check registered response
				genFile.P("if ms.contents.", m.GoName, " != nil {")
				genFile.P("return ms.contents.", m.GoName, ", nil")
				genFile.P("}")

				// return defaut value
				genFile.P("return &", genFile.QualifiedGoIdent(m.Output.GoIdent), "{")
				for _, f := range m.Output.Fields {
					genFile.P(f.GoName, ": ", MockFieldValue(f), ",")
				}
				genFile.P("}, nil")
				genFile.P("}")
				genFile.P("")
			}

			// func RegisterXServer()
			genFile.P("func Register", mockServerName, "(s grpc.ServiceRegistrar) (*", mockServerName, ") {")
			genFile.P("ms := &", mockServerName, "{}")
			genFile.P("Register", baseServerName, "(s, ms)")
			genFile.P("return ms")
			genFile.P("}")
		}

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
