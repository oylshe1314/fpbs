package parser

import "fpbs"

var fieldTypes = map[string]fpbs.FieldType{
	"bool":      fpbs.FieldTypeBool,
	"int8":      fpbs.FieldTypeInt8,
	"int16":     fpbs.FieldTypeInt16,
	"int32":     fpbs.FieldTypeInt32,
	"int64":     fpbs.FieldTypeInt64,
	"uint8":     fpbs.FieldTypeUint8,
	"uint16":    fpbs.FieldTypeUint16,
	"uint32":    fpbs.FieldTypeUint32,
	"uint64":    fpbs.FieldTypeUint64,
	"float32":   fpbs.FieldTypeFloat32,
	"float64":   fpbs.FieldTypeFloat64,
	"string":    fpbs.FieldTypeString,
	"bool[]":    fpbs.FieldTypeBoolArray,
	"int8[]":    fpbs.FieldTypeInt8Array,
	"int16[]":   fpbs.FieldTypeInt16Array,
	"int32[]":   fpbs.FieldTypeInt32Array,
	"int64[]":   fpbs.FieldTypeInt64Array,
	"uint8[]":   fpbs.FieldTypeUint8Array,
	"uint16[]":  fpbs.FieldTypeUint16Array,
	"uint32[]":  fpbs.FieldTypeUint32Array,
	"uint64[]":  fpbs.FieldTypeUint64Array,
	"float32[]": fpbs.FieldTypeFloat32Array,
	"float64[]": fpbs.FieldTypeFloat64Array,
	"string[]":  fpbs.FieldTypeStringArray,
}

var fieldTypeNames = map[fpbs.FieldType]string{
	fpbs.FieldTypeBool:         "bool",
	fpbs.FieldTypeInt8:         "int8",
	fpbs.FieldTypeInt16:        "int16",
	fpbs.FieldTypeInt32:        "int32",
	fpbs.FieldTypeInt64:        "int64",
	fpbs.FieldTypeUint8:        "uint8",
	fpbs.FieldTypeUint16:       "uint16",
	fpbs.FieldTypeUint32:       "uint32",
	fpbs.FieldTypeUint64:       "uint64",
	fpbs.FieldTypeFloat32:      "float32",
	fpbs.FieldTypeFloat64:      "float64",
	fpbs.FieldTypeString:       "string",
	fpbs.FieldTypeBoolArray:    "bool[]",
	fpbs.FieldTypeInt8Array:    "int8[]",
	fpbs.FieldTypeInt16Array:   "int16[]",
	fpbs.FieldTypeInt32Array:   "int32[]",
	fpbs.FieldTypeInt64Array:   "int64[]",
	fpbs.FieldTypeUint8Array:   "uint8[]",
	fpbs.FieldTypeUint16Array:  "uint16[]",
	fpbs.FieldTypeUint32Array:  "uint32[]",
	fpbs.FieldTypeUint64Array:  "uint64[]",
	fpbs.FieldTypeFloat32Array: "float32[]",
	fpbs.FieldTypeFloat64Array: "float64[]",
	fpbs.FieldTypeStringArray:  "string[]",
}

var golangFieldTypes = map[string]fpbs.FieldType{
	"bool":      fpbs.FieldTypeBool,
	"int8":      fpbs.FieldTypeInt8,
	"int16":     fpbs.FieldTypeInt16,
	"int32":     fpbs.FieldTypeInt32,
	"int64":     fpbs.FieldTypeInt64,
	"uint8":     fpbs.FieldTypeUint8,
	"uint16":    fpbs.FieldTypeUint16,
	"uint32":    fpbs.FieldTypeUint32,
	"uint64":    fpbs.FieldTypeUint64,
	"float32":   fpbs.FieldTypeFloat32,
	"float64":   fpbs.FieldTypeFloat64,
	"string":    fpbs.FieldTypeString,
	"[]bool":    fpbs.FieldTypeBoolArray,
	"[]int8":    fpbs.FieldTypeInt8Array,
	"[]int16":   fpbs.FieldTypeInt16Array,
	"[]int32":   fpbs.FieldTypeInt32Array,
	"[]int64":   fpbs.FieldTypeInt64Array,
	"[]uint8":   fpbs.FieldTypeUint8Array,
	"[]uint16":  fpbs.FieldTypeUint16Array,
	"[]uint32":  fpbs.FieldTypeUint32Array,
	"[]uint64":  fpbs.FieldTypeUint64Array,
	"[]float32": fpbs.FieldTypeFloat32Array,
	"[]float64": fpbs.FieldTypeFloat64Array,
	"[]string":  fpbs.FieldTypeStringArray,
}

var golangFieldTypeNames = map[fpbs.FieldType]string{
	fpbs.FieldTypeBool:         "bool",
	fpbs.FieldTypeInt8:         "int8",
	fpbs.FieldTypeInt16:        "int16",
	fpbs.FieldTypeInt32:        "int32",
	fpbs.FieldTypeInt64:        "int64",
	fpbs.FieldTypeUint8:        "uint8",
	fpbs.FieldTypeUint16:       "uint16",
	fpbs.FieldTypeUint32:       "uint32",
	fpbs.FieldTypeUint64:       "uint64",
	fpbs.FieldTypeFloat32:      "float32",
	fpbs.FieldTypeFloat64:      "float64",
	fpbs.FieldTypeString:       "string",
	fpbs.FieldTypeBoolArray:    "[]bool",
	fpbs.FieldTypeInt8Array:    "[]int8",
	fpbs.FieldTypeInt16Array:   "[]int16",
	fpbs.FieldTypeInt32Array:   "[]int32",
	fpbs.FieldTypeInt64Array:   "[]int64",
	fpbs.FieldTypeUint8Array:   "[]uint8",
	fpbs.FieldTypeUint16Array:  "[]uint16",
	fpbs.FieldTypeUint32Array:  "[]uint32",
	fpbs.FieldTypeUint64Array:  "[]uint64",
	fpbs.FieldTypeFloat32Array: "[]float32",
	fpbs.FieldTypeFloat64Array: "[]float64",
	fpbs.FieldTypeStringArray:  "[]string",
}

var csharpFieldTypes = map[string]fpbs.FieldType{
	"bool":     fpbs.FieldTypeBool,
	"sbyte":    fpbs.FieldTypeInt8,
	"short":    fpbs.FieldTypeInt16,
	"int":      fpbs.FieldTypeInt32,
	"long":     fpbs.FieldTypeInt64,
	"byte":     fpbs.FieldTypeUint8,
	"ushort":   fpbs.FieldTypeUint16,
	"uint":     fpbs.FieldTypeUint32,
	"ulong":    fpbs.FieldTypeUint64,
	"float":    fpbs.FieldTypeFloat32,
	"double":   fpbs.FieldTypeFloat64,
	"string":   fpbs.FieldTypeString,
	"bool[]":   fpbs.FieldTypeBoolArray,
	"sbyte[]":  fpbs.FieldTypeInt8Array,
	"short[]":  fpbs.FieldTypeInt16Array,
	"int[]":    fpbs.FieldTypeInt32Array,
	"long[]":   fpbs.FieldTypeInt64Array,
	"byte[]":   fpbs.FieldTypeUint8Array,
	"ushort[]": fpbs.FieldTypeUint16Array,
	"uint[]":   fpbs.FieldTypeUint32Array,
	"ulong[]":  fpbs.FieldTypeUint64Array,
	"float[]":  fpbs.FieldTypeFloat32Array,
	"double[]": fpbs.FieldTypeFloat64Array,
	"string[]": fpbs.FieldTypeStringArray,
}

var csharpFieldTypeNames = map[fpbs.FieldType]string{
	fpbs.FieldTypeBool:         "bool",
	fpbs.FieldTypeInt8:         "sbyte",
	fpbs.FieldTypeInt16:        "short",
	fpbs.FieldTypeInt32:        "int",
	fpbs.FieldTypeInt64:        "long",
	fpbs.FieldTypeUint8:        "byte",
	fpbs.FieldTypeUint16:       "ushort",
	fpbs.FieldTypeUint32:       "uint",
	fpbs.FieldTypeUint64:       "ulong",
	fpbs.FieldTypeFloat32:      "float",
	fpbs.FieldTypeFloat64:      "double",
	fpbs.FieldTypeString:       "string",
	fpbs.FieldTypeBoolArray:    "bool[]",
	fpbs.FieldTypeInt8Array:    "sbyte[]",
	fpbs.FieldTypeInt16Array:   "short[]",
	fpbs.FieldTypeInt32Array:   "int[]",
	fpbs.FieldTypeInt64Array:   "long[]",
	fpbs.FieldTypeUint8Array:   "byte[]",
	fpbs.FieldTypeUint16Array:  "ushort[]",
	fpbs.FieldTypeUint32Array:  "uint[]",
	fpbs.FieldTypeUint64Array:  "ulong[]",
	fpbs.FieldTypeFloat32Array: "float[]",
	fpbs.FieldTypeFloat64Array: "double[]",
	fpbs.FieldTypeStringArray:  "string[]",
}

func GetFieldType(name string) fpbs.FieldType {
	tipe, ok := fieldTypes[name]
	if ok {
		return tipe
	} else {
		return fpbs.FieldTypeNone
	}
}

func GetFieldTypeName(tipe fpbs.FieldType) string {
	return fieldTypeNames[tipe]
}

func GetGolangFieldType(name string) fpbs.FieldType {
	tipe, ok := golangFieldTypes[name]
	if ok {
		return tipe
	} else {
		return fpbs.FieldTypeNone
	}
}

func GetGolangFieldTypeName(tipe fpbs.FieldType) string {
	return golangFieldTypeNames[tipe]
}

func GetCsharpFieldType(name string) fpbs.FieldType {
	tipe, ok := csharpFieldTypes[name]
	if ok {
		return tipe
	} else {
		return fpbs.FieldTypeNone
	}
}

func GetCsharpFieldTypeName(tipe fpbs.FieldType) string {
	return csharpFieldTypeNames[tipe]
}
