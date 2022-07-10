package main

import "fmt"

type Import struct {
	Alias string
	Name  string
}

func (i Import) String() string {
	return fmt.Sprintf("%s \"%s\"", i.Alias, i.Name)
}
