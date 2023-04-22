package json

import (
	"encoding/json"
	"fpbs"
	"fpbs/errors"
	"fpbs/gen/parser"
	"fpbs/util"
	"os"
	"strings"
)

type _Package struct {
	Name    string     `json:"package"`
	Imports []string   `json:"imports"`
	Enums   []*_Enum   `json:"enums"`
	Structs []*_Struct `json:"structs"`

	enums           map[string]struct{}
	structs         map[string]struct{}
	importedStructs map[string]struct{}

	importedPackages map[string]struct{}
}

type _Enum struct {
	Name    string       `json:"name"`
	Type    string       `json:"type"`
	Comment string       `json:"comment"`
	Items   []*_EnumItem `json:"values"`

	items map[string]struct{}
}

type _EnumItem struct {
	Name    string      `json:"name"`
	Value   interface{} `json:"value"`
	Comment string      `json:"comment"`
}

type _Struct struct {
	Name    string          `json:"name"`
	Comment string          `json:"comment"`
	Fields  []*_StructField `json:"fields"`

	fields map[string]struct{}
	keys   map[byte]struct{}
}

type _StructField struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Key     byte   `json:"key"`
	Comment string `json:"comment"`

	fieldType fpbs.FieldType
	fieldKey  fpbs.FieldKey
}

type jsonParser struct {
}

func NewParser() parser.Parser {
	return &jsonParser{}
}

func (this *jsonParser) ParseFiles(files []string, incDirs []string) (map[string]*parser.Package, error) {
	var parsedPackages = map[string]*parser.Package{}
	for _, file := range files {
		_package, err := this.parseFile(file, incDirs)
		if err != nil {
			return nil, err
		}

		var parsedFile = &parser.File{
			Name: util.Filename(file),
		}
		for _, _struct := range _package.Structs {
			var parsedStruct = &parser.Struct{
				Name:    _struct.Name,
				Comment: _struct.Comment,
			}

			for _, field := range _struct.Fields {
				parsedStruct.Fields = append(parsedStruct.Fields, &parser.StructField{
					Name:       field.Name,
					Type:       field.Type,
					Key:        field.Key,
					Comment:    field.Comment,
					ParsedType: field.fieldType,
				})
			}

			parsedFile.Structs = append(parsedFile.Structs, parsedStruct)
		}

		for importPackageName := range _package.importedPackages {
			parsedFile.Imports = append(parsedFile.Imports, importPackageName)
		}

		var parsedPackage = parsedPackages[_package.Name]
		if parsedPackage == nil {
			parsedPackage = &parser.Package{Name: _package.Name}
			parsedPackages[_package.Name] = parsedPackage
		}

		parsedPackage.Files = append(parsedPackage.Files, parsedFile)
	}
	return parsedPackages, nil
}

func (this *jsonParser) parseFiles(files []string, incDirs []string) (map[string]*_Package, error) {
	var packages = map[string]*_Package{}
	for _, file := range files {
		_package, err := this.parseFile(file, incDirs)
		if err != nil {
			return nil, err
		}

		packages[file] = _package
	}
	return packages, nil
}

func (this *jsonParser) parseFile(file string, incDirs []string) (*_Package, error) {
	var p, err = os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var _package = &_Package{enums: map[string]struct{}{}, structs: map[string]struct{}{}, importedStructs: map[string]struct{}{}}
	err = json.Unmarshal(p, _package)
	if err != nil {
		return nil, err
	}

	if !parser.CheckPackageName(_package.Name) {
		return nil, errors.Errorf("file: %s, illegal package name: %s", file, _package.Name)
	}

	var importPackages = map[string]*_Package{}
	for _, importFile := range _package.Imports {
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
	}

	_package.structs = map[string]struct{}{}
	for _, _struct := range _package.Structs {
		_, ok := _package.structs[_struct.Name]
		if ok {
			return nil, errors.Errorf("file: %s, repeated struct name: %s", file, _struct.Name)
		}

		if !parser.CheckStructName(_struct.Name) {
			return nil, errors.Errorf("file: %s, illegal struct name: %s", file, _struct.Name)
		}

		_struct.fields = map[string]struct{}{}
		_struct.keys = map[byte]struct{}{}
		for _, field := range _struct.Fields {
			_, ok = _package.structs[_struct.Name]
			if ok {
				return nil, errors.Errorf("file: %s, struct: %s, illegal field name: %s", file, _struct.Name, field.Name)
			}
			if !parser.CheckFieldName(field.Name) {
				return nil, errors.Errorf("file: %s, struct: %s, illegal field name: %s", file, _struct.Name, field.Name)
			}
			if field.Key == 0 {
				return nil, errors.Errorf("file: %s, struct: %s, field: %s, could not be 0", file, _struct.Name, field.Name)
			}
			_, ok = _struct.keys[field.Key]
			if ok {
				return nil, errors.Errorf("file: %s, struct: %s, field: %s, repeated field key: %d", file, _struct.Name, field.Name, field.fieldKey)
			}

			field.fieldType = parser.GetFieldType(field.Type)
			if field.fieldType == fpbs.FieldTypeNone {
				var isArray = false
				var fieldTypeName = field.Type
				if strings.HasSuffix(fieldTypeName, "[]") {
					isArray = true
					fieldTypeName = fieldTypeName[:len(fieldTypeName)-2]
				}

				_, ok = _package.structs[fieldTypeName]
				if ok {
					if isArray {
						field.fieldType = fpbs.FieldTypeStructArray
					} else {
						field.fieldType = fpbs.FieldTypeStruct
					}
				} else {
					for _, importPackage := range importPackages {
						_, ok = importPackage.structs[fieldTypeName]
						if ok {
							if importPackage.Name != _package.Name {
								_package.importedStructs[importPackage.Name] = struct{}{}
							}
							if isArray {
								field.fieldType = fpbs.FieldTypeStructArray
							} else {
								field.fieldType = fpbs.FieldTypeStruct
							}
						}
					}
				}

				field.Type = fieldTypeName

				if field.fieldType == fpbs.FieldTypeNone {
					return nil, errors.Errorf("file: '%s', struct: '%s', undefined type: %s", file, _struct.Name, fieldTypeName)
				}
			}

			field.fieldKey = fpbs.FieldKey(field.Key)

			_struct.keys[field.Key] = struct{}{}
			_struct.fields[field.Name] = struct{}{}
		}
		_package.structs[_struct.Name] = struct{}{}
	}

	return _package, nil
}

func (this *jsonParser) findImport(importFile string, incDirs []string) (string, error) {
	var files []string
	for _, dir := range incDirs {
		file := dir + importFile
		_, err := os.Stat(file)
		if err == nil {
			files = append(files, file)
		}
	}

	if len(files) == 0 {
		return "", errors.Errorf("could not find the file: %s", importFile)
	}

	if len(files) > 1 {
		return "", errors.Errorf("found multiple files with the same name: %s.\n%s\n", importFile, strings.Join(files, "\n"))
	}

	return files[0], nil
}
