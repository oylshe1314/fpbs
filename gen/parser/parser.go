package parser

import (
	"fpbs"
)

type Package struct {
	Name  string
	Files []*File
}

type File struct {
	Name    string
	Structs []*Struct
	Imports []string
}

type Enum[T any] struct {
	Name    string
	Type    string
	Comment string
	Values  []*EnumValue[T] `json:"values"`
}

type EnumValue[T any] struct {
	Name    string `json:"name"`
	Value   T      `json:"value"`
	Comment string `json:"comment"`
}

type Struct struct {
	Name    string
	Comment string
	Fields  []*StructField
}

type StructField struct {
	Name    string
	Type    string
	Key     byte
	Comment string

	ParsedType fpbs.FieldType
}

type Parser interface {
	ParseFiles(files []string, incDirs []string) (map[string]*Package, error)
}
