package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

type ImportDecl struct {
	Alias string
	Name  string
}

func (d ImportDecl) String() string {
	return fmt.Sprintf("%s \"%s\"", d.Alias, d.Name)
}

type ErrorDecl struct {
	Name    string
	Message string
}

func (d ErrorDecl) String() string {
	return fmt.Sprintf("%s = errors.New(\"%s\")", d.Name, d.Message)
}

type GenFunc func(file *protogen.File, genFile *protogen.GeneratedFile) error
