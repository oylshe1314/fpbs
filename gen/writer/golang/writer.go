package golang

import (
	"bytes"
	"fmt"
	"fpbs"
	"fpbs/gen/parser"
	"fpbs/gen/writer"
	"fpbs/util"
	"os"
	"sort"
	"strings"
)

type golangWriter struct {
}

func NewWriter() writer.Writer {
	return &golangWriter{}
}

func (this *golangWriter) Write(outDir string, packages map[string]*parser.Package) error {
	var err error
	if !strings.HasSuffix(outDir, "/") && !strings.HasSuffix(outDir, "\\") {
		outDir = outDir + "/"
	}
	for _, _package := range packages {
		var packDir = outDir + _package.Name
		err = os.MkdirAll(packDir, 777)
		if err != nil {
			return err
		}

		for _, file := range _package.Files {
			err = this.writeFile(_package.Name, packDir, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (this *golangWriter) writeFile(packName, packDir string, file *parser.File) error {
	var buff bytes.Buffer

	buff.WriteString(fmt.Sprintf("package %s\n", packName))
	buff.WriteByte('\n')

	var imports = []string{"ecs/framework/fpbs"}

	for _, _import := range file.Imports {
		imports = append(imports, _import)
	}

	sort.Slice(imports, func(i, j int) bool {
		return imports[i] < imports[j]
	})

	switch {
	case len(imports) == 1:
		buff.WriteString(fmt.Sprintf("import \"%s\"\n", imports[0]))
	case len(imports) > 1:
		buff.WriteString("import (\n")
		for _, _import := range imports {
			buff.WriteString(fmt.Sprintf("\t\"%s\"\n", _import))
		}
		buff.WriteString(")\n")
	}

	for _, _struct := range file.Structs {
		var structName = util.UpperCamelCase(_struct.Name)

		buff.WriteByte('\n')

		// structure define
		if _struct.Comment != "" {
			buff.WriteString(fmt.Sprintf("// %s %s\n", structName, _struct.Comment))
		}
		buff.WriteString(fmt.Sprintf("type %s struct {\n", structName))

		var maxFieldLength = 0
		var maxTypeLength = 0
		var maxTagLength = 0
		for _, field := range _struct.Fields {
			if len(field.Name) > maxFieldLength {
				maxFieldLength = len(field.Name)
			}
			if len(field.Type) > maxTypeLength {
				maxTypeLength = len(field.Type)
			}
			switch {
			case field.Key > 99:
				maxTagLength = 6
			case field.Key > 9:
				maxTagLength = 5
			default:
				maxTagLength = 4
			}
		}

		var fmtStr = fmt.Sprintf("\t%%-%ds %%-%ds `key:\"%%d\" json:\"%%s,omitempty\"`", maxFieldLength, maxTypeLength)

		for _, field := range _struct.Fields {
			var tagLength = 0
			switch {
			case field.Key > 99:
				tagLength = 3
			case field.Key > 9:
				tagLength = 2
			default:
				tagLength = 1
			}
			switch field.ParsedType {
			case fpbs.FieldTypeStruct:
				if field.Comment == "" {
					buff.WriteString(fmt.Sprintf(fmtStr, util.UpperCamelCase(field.Name), "*"+util.UpperCamelCase(field.Type), field.Key, field.Name))
				} else {
					buff.WriteString(fmt.Sprintf(fmtStr+fmt.Sprintf("%%%ds %%s", maxTagLength-tagLength), util.UpperCamelCase(field.Name), "*"+field.Type, field.Key, field.Name, "//", field.Comment))
				}
			case fpbs.FieldTypeStructArray:
				if field.Comment == "" {
					buff.WriteString(fmt.Sprintf(fmtStr, util.UpperCamelCase(field.Name), "[]*"+util.UpperCamelCase(field.Type), field.Key, field.Name))
				} else {
					buff.WriteString(fmt.Sprintf(fmtStr+fmt.Sprintf("%%%ds %%s", maxTagLength-tagLength), util.UpperCamelCase(field.Name), "[]*"+field.Type, field.Key, field.Name, "//", field.Comment))
				}
			default:
				var fieldTypeName = parser.GetGolangFieldTypeName(field.ParsedType)
				if field.Comment == "" {
					buff.WriteString(fmt.Sprintf(fmtStr, util.UpperCamelCase(field.Name), fieldTypeName, field.Key, field.Name))
				} else {
					buff.WriteString(fmt.Sprintf(fmtStr+fmt.Sprintf("%%%ds %%s", maxTagLength-tagLength), util.UpperCamelCase(field.Name), fieldTypeName, field.Key, field.Name, "//", field.Comment))
				}
			}
			buff.WriteString("\n")
		}

		buff.WriteString("}\n")

		//structure write functions
		buff.WriteString(fmt.Sprintf("func Write%s(msg *%s) []byte {\n", structName, structName))
		buff.WriteString(fmt.Sprintf("\treturn fpbs.Write(write%s(msg))\n", structName))
		buff.WriteString("}\n")
		buff.WriteString("\n")
		buff.WriteString(fmt.Sprintf("func write%s(msg *%s) fpbs.Fields {\n", structName, structName))
		buff.WriteString("\tif msg == nil {\n")
		buff.WriteString("\t\treturn nil\n")
		buff.WriteString("\t}\n")
		buff.WriteString("\n")
		buff.WriteString("\tvar fields = fpbs.Fields{}\n")
		for _, field := range _struct.Fields {
			switch field.ParsedType {
			case fpbs.FieldTypeStruct:
				buff.WriteString(fmt.Sprintf("\tfields.PutFields(%d, write%s(msg.%s))\n", field.Key, util.UpperCamelCase(field.Type), util.UpperCamelCase(field.Name)))
			case fpbs.FieldTypeStructArray:
				buff.WriteString(fmt.Sprintf("\tfields.PutFieldsArray(%d, write%sArray(msg.%s))\n", field.Key, util.UpperCamelCase(field.Type), util.UpperCamelCase(field.Name)))
			default:
				var isArray = false
				var fieldTypeName = parser.GetGolangFieldTypeName(field.ParsedType)
				if field.ParsedType > fpbs.FieldTypeStruct {
					isArray = true
					fieldTypeName = parser.GetGolangFieldTypeName(field.ParsedType - fpbs.FieldTypeStruct)
				}
				if isArray {
					buff.WriteString(fmt.Sprintf("\tfields.Put%sArray(%d, msg.%s)\n", util.UpperCamelCase(fieldTypeName), field.Key, util.UpperCamelCase(field.Name)))
				} else {
					buff.WriteString(fmt.Sprintf("\tfields.Put%s(%d, msg.%s)\n", util.UpperCamelCase(fieldTypeName), field.Key, util.UpperCamelCase(field.Name)))
				}
			}
		}
		buff.WriteString("\n")
		buff.WriteString("\treturn fields\n")
		buff.WriteString("}\n")
		buff.WriteString("\n")
		buff.WriteString(fmt.Sprintf("func write%sArray(array []*%s) []fpbs.Fields {\n", structName, structName))
		buff.WriteString("\tif len(array) == 0 {\n")
		buff.WriteString("\t\treturn nil\n")
		buff.WriteString("\t}\n")
		buff.WriteString("\n")
		buff.WriteString("\tvar fieldsArray = make([]fpbs.Fields, len(array))\n")
		buff.WriteString("\tfor i := range array {\n")
		buff.WriteString(fmt.Sprintf("\t\tfieldsArray[i] = write%s(array[i])\n", structName))
		buff.WriteString("\t}\n")
		buff.WriteString("\n")
		buff.WriteString("\treturn fieldsArray\n")
		buff.WriteString("}\n")
		buff.WriteString("\n")

		//structure parse functions
		buff.WriteString(fmt.Sprintf("func Parse%s(p []byte) (*%s, error) {\n", structName, structName))
		buff.WriteString("\tfields, err := fpbs.Parse(p)\n")
		buff.WriteString("\tif err != nil {\n")
		buff.WriteString("\t\treturn nil, err\n")
		buff.WriteString("\t}\n")
		buff.WriteString(fmt.Sprintf("\treturn parse%s(fields), nil\n", structName))
		buff.WriteString("}\n")
		buff.WriteString("\n")
		buff.WriteString(fmt.Sprintf("func parse%s(fields fpbs.Fields) *%s {\n", structName, structName))
		buff.WriteString("\tif fields == nil {\n")
		buff.WriteString("\t\treturn nil\n")
		buff.WriteString("\t}\n")
		buff.WriteString("\n")
		buff.WriteString(fmt.Sprintf("\tvar msg = new(%s)\n", structName))
		for _, field := range _struct.Fields {
			switch field.ParsedType {
			case fpbs.FieldTypeStruct:
				buff.WriteString(fmt.Sprintf("\tmsg.%s = parse%s(fields.GetFields(%d))\n", util.UpperCamelCase(field.Name), util.UpperCamelCase(field.Type), field.Key))
			case fpbs.FieldTypeStructArray:
				buff.WriteString(fmt.Sprintf("\tmsg.%s = parse%sArray(fields.GetFieldsArray(%d))\n", util.UpperCamelCase(field.Name), util.UpperCamelCase(field.Type), field.Key))
			default:
				var isArray = false
				var fieldTypeName = parser.GetGolangFieldTypeName(field.ParsedType)
				if field.ParsedType > fpbs.FieldTypeStruct {
					isArray = true
					fieldTypeName = parser.GetGolangFieldTypeName(field.ParsedType - fpbs.FieldTypeStruct)
				}
				if isArray {
					buff.WriteString(fmt.Sprintf("\tmsg.%s = fields.Get%sArray(%d)\n", util.UpperCamelCase(field.Name), util.UpperCamelCase(fieldTypeName), field.Key))
				} else {
					buff.WriteString(fmt.Sprintf("\tmsg.%s = fields.Get%s(%d)\n", util.UpperCamelCase(field.Name), util.UpperCamelCase(fieldTypeName), field.Key))
				}
			}
		}
		buff.WriteString("\n")
		buff.WriteString("\treturn msg\n")
		buff.WriteString("}\n")
		buff.WriteString("\n")
		buff.WriteString(fmt.Sprintf("func parse%sArray(fieldsArray []fpbs.Fields) []*%s {\n", structName, structName))
		buff.WriteString("\tif len(fieldsArray) == 0 {\n")
		buff.WriteString("\t\treturn nil\n")
		buff.WriteString("\t}\n")
		buff.WriteString("\n")
		buff.WriteString(fmt.Sprintf("\tvar array = make([]*%s, len(fieldsArray))\n", structName))
		buff.WriteString("\tfor i := range fieldsArray {\n")
		buff.WriteString(fmt.Sprintf("\t\tarray[i] = parse%s(fieldsArray[i])\n", structName))
		buff.WriteString("\t}\n")
		buff.WriteString("\treturn array\n")
		buff.WriteString("}\n")
		buff.WriteString("\n")
	}

	var err = os.WriteFile(packDir+"/"+file.Name+".go", buff.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}
