package fpbs

type FieldType byte

const (
	FieldTypeNone FieldType = iota
	FieldTypeBool
	FieldTypeInt8
	FieldTypeInt16
	FieldTypeInt32
	FieldTypeInt64
	FieldTypeUint8
	FieldTypeUint16
	FieldTypeUint32
	FieldTypeUint64
	FieldTypeFloat32
	FieldTypeFloat64
	FieldTypeString
	FieldTypeStruct
	FieldTypeBoolArray
	FieldTypeInt8Array
	FieldTypeInt16Array
	FieldTypeInt32Array
	FieldTypeInt64Array
	FieldTypeUint8Array
	FieldTypeUint16Array
	FieldTypeUint32Array
	FieldTypeUint64Array
	FieldTypeFloat32Array
	FieldTypeFloat64Array
	FieldTypeStringArray
	FieldTypeStructArray
)

type FieldKey byte
