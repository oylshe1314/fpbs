package json

import (
	"bytes"
	"fpbs"
	"fpbs/errors"
	"fpbs/gen/parser"
	"fpbs/util"
	"os"
	"strconv"
	"unicode"
)

type _Package struct {
	Name    string
	Imports []string
	Enums   []*_Enum
	Structs []*_Struct

	enums           map[string]struct{}
	structs         map[string]struct{}
	importedStructs map[string]struct{}

	importedPackages map[string]struct{}
}

type _Enum struct {
	Name    string
	Type    string
	Comment string
	Items   []*_EnumItem

	items map[string]struct{}
}

type _EnumItem struct {
	Name    string
	Value   interface{}
	Comment string
}

type _Struct struct {
	Name    string
	Comment string
	Fields  []*_StructField

	fields map[string]struct{}
	keys   map[byte]struct{}
}

type _StructField struct {
	Name    string
	Type    string
	Key     byte
	Comment string

	fieldType fpbs.FieldType
	fieldKey  fpbs.FieldKey
}

type idlParser struct {
}

func (this *idlParser) ParseFiles(files []string, incDirs []string) (map[string]*parser.Package, error) {

	var packages = map[string]*parser.Package{}
	for _, file := range files {
		_package, err := this.parseFile(file, incDirs)
		if err != nil {
			return nil, err
		}

		var parsedFile = &parser.File{
			Name: util.Filename(file),
		}

		var parsedPackage = packages[_package.Name]
		if parsedPackage == nil {
			parsedPackage = &parser.Package{
				Name: _package.Name,
			}
			packages[parsedPackage.Name] = parsedPackage
		}
		parsedPackage.Files = append(parsedPackage.Files, parsedFile)
	}

	return nil, nil
}

func (this *idlParser) parseFile(file string, incDirs []string) (*_Package, error) {
	p, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var runes = bytes.Runes(p)

	var word []rune
	var rowNum = 1
	var columnNum = 0

	var status = statusPackageDefine
	var _package *_Package
	//var _enum *_Enum
	var _struct *_Struct

	var importPackages = map[string]*_Package{}
	for _, r := range runes {
		columnNum += 1
		if unicode.IsSpace(r) {
			if r == '\n' {
				rowNum += 1
				columnNum = 0
			}

			if len(word) == 0 {
				continue
			}

			switch status {
			case statusPackageDefine:
				if string(word) == "package" {
					status = statusPackageDefining
				} else {
					return nil, errors.Error("should define package at first")
				}
			case statusPackageDefining:
				var statement = string(word)
				if parser.CheckPackageName(statement) {
					_package = &_Package{
						Name:             statement,
						enums:            map[string]struct{}{},
						structs:          map[string]struct{}{},
						importedStructs:  map[string]struct{}{},
						importedPackages: map[string]struct{}{},
					}
					status = statusImportDefine
				} else {
					return nil, errors.Errorf("%s:%d:%d: unexpected statement '%s'", file, rowNum, columnNum, statement)
				}
			case statusImportDefine, statusEnumDefine, statusStructDefine:
				var statement = string(word)
				switch statement {
				case "import":
					status = statusImportDefining
				case "enum":
					status = statusEnumDefining
				case "struct":
					status = statusStructDefining
				default:
					return nil, errors.Errorf("%s:%d:%d: unexpected statement '%s'", file, rowNum, columnNum, statement)
				}
			case statusImportDefining:
				var statement = string(word)
				importFile, err := strconv.Unquote(statement)
				if err != nil {
					return nil, errors.Errorf("%s:%d:%d: unexpected statement '%s'", file, rowNum, columnNum, statement)
				}

				importFile, err = this.findImport(importFile, incDirs)
				if err != nil {
					return nil, err
				}

				importPackage, err := this.parseFile(importFile, incDirs)
				if err != nil {
					return nil, err
				}

				for importStructName := range importPackage.structs {
					for importedFile, importedPackage := range importPackages {
						_, ok := importedPackage.structs[importStructName]
						if ok {
							return nil, errors.Errorf("file: %s, import file: %s,  the struct: %s was defined in file: %s", file, importFile, importStructName, importedFile)
						}
					}
				}

				importPackages[importFile] = importPackage
				if importPackage.Name != _package.Name {
					_package.importedPackages[importPackage.Name] = struct{}{}
				}
			case statusStructDefining:
				var statement = string(word)
				if statement == "{" {
					if _struct != nil {
						status = statusFieldsDefine
					} else {
						return nil, errors.Errorf("%s:%d:%d: unexpected statement '%s'", file, rowNum, columnNum, statement)
					}
				} else {
					if parser.CheckStructName(statement) {
						_struct = &_Struct{
							Name:   statement,
							fields: map[string]struct{}{},
							keys:   map[byte]struct{}{},
						}
					} else {
						return nil, errors.Errorf("%s:%d:%d: unexpected statement '%s'", file, rowNum, columnNum, statement)
					}
				}
			}
		} else {
			word = append(word, r)
		}
	}

	return nil, nil
}

func (this *idlParser) findImport(file string, incDirs []string) (string, error) {
	return "", nil
}
