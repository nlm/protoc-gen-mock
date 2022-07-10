package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

type Import struct {
	Alias string
	Name  string
}

func (i Import) String() string {
	return fmt.Sprintf("%s \"%s\"", i.Alias, i.Name)
}

type GenFunc func(file *protogen.File, genFile *protogen.GeneratedFile) error
