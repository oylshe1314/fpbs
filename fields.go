package fpbs

import "fpbs/util"

type Fields map[FieldKey]*util.Pair[FieldType, interface{}]

func (fields Fields) Get(key FieldKey) *util.Pair[FieldType, interface{}] {
	return fields[key]
}

func (fields Fields) Put(key FieldKey, pair *util.Pair[FieldType, interface{}]) {
	fields[key] = pair
}

func (fields Fields) GetBool(key FieldKey) bool {
	if pair := fields.Get(key); pair != nil || pair.Value != nil {
		return pair.Value.(bool)
	}
	return false
}

func (fields Fields) GetInt8(key FieldKey) int8 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(int8)
	}
	return 0
}

func (fields Fields) GetInt16(key FieldKey) int16 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(int16)
	}
	return 0
}

func (fields Fields) GetInt32(key FieldKey) int32 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(int32)
	}
	return 0
}

func (fields Fields) GetInt64(key FieldKey) int64 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(int64)
	}
	return 0
}

func (fields Fields) GetUint8(key FieldKey) uint8 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(uint8)
	}
	return 0
}

func (fields Fields) GetUint16(key FieldKey) uint16 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(uint16)
	}
	return 0
}

func (fields Fields) GetUint32(key FieldKey) uint32 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(uint32)
	}
	return 0
}

func (fields Fields) GetUint64(key FieldKey) uint64 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(uint64)
	}
	return 0
}

func (fields Fields) GetFloat32(key FieldKey) float32 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(float32)
	}
	return 0
}

func (fields Fields) GetFloat64(key FieldKey) float64 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(float64)
	}
	return 0
}

func (fields Fields) GetString(key FieldKey) string {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(string)
	}
	return ""
}

func (fields Fields) GetFields(key FieldKey) Fields {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.(Fields)
	}
	return nil
}

func (fields Fields) GetBoolArray(key FieldKey) []bool {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]bool)
	}
	return nil
}

func (fields Fields) GetInt8Array(key FieldKey) []int8 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]int8)
	}
	return nil
}

func (fields Fields) GetInt16Array(key FieldKey) []int16 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]int16)
	}
	return nil
}

func (fields Fields) GetInt32Array(key FieldKey) []int32 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]int32)
	}
	return nil
}

func (fields Fields) GetInt64Array(key FieldKey) []int64 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]int64)
	}
	return nil
}

func (fields Fields) GetUint8Array(key FieldKey) []uint8 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]uint8)
	}
	return nil
}

func (fields Fields) GetUint16Array(key FieldKey) []uint16 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]uint16)
	}
	return nil
}

func (fields Fields) GetUint32Array(key FieldKey) []uint32 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]uint32)
	}
	return nil
}

func (fields Fields) GetUint64Array(key FieldKey) []uint64 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]uint64)
	}
	return nil
}

func (fields Fields) GetFloat32Array(key FieldKey) []float32 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]float32)
	}
	return nil
}

func (fields Fields) GetFloat64Array(key FieldKey) []float64 {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]float64)
	}
	return nil
}

func (fields Fields) GetStringArray(key FieldKey) []string {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]string)
	}
	return nil
}

func (fields Fields) GetFieldsArray(key FieldKey) []Fields {
	if pair := fields.Get(key); pair != nil && pair.Value != nil {
		return pair.Value.([]Fields)
	}
	return nil
}

func (fields Fields) PutBool(key FieldKey, val bool) {
	if val {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeBool, Value: val})
	}
}

func (fields Fields) PutInt8(key FieldKey, val int8) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeInt8, Value: val})
	}
}

func (fields Fields) PutInt16(key FieldKey, val int16) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeInt16, Value: val})
	}
}

func (fields Fields) PutInt32(key FieldKey, val int32) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeInt32, Value: val})
	}
}

func (fields Fields) PutInt64(key FieldKey, val int64) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeInt64, Value: val})
	}
}

func (fields Fields) PutUint8(key FieldKey, val uint8) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeUint8, Value: val})
	}
}

func (fields Fields) PutUint16(key FieldKey, val uint16) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeUint16, Value: val})
	}
}

func (fields Fields) PutUint32(key FieldKey, val uint32) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeUint32, Value: val})
	}
}

func (fields Fields) PutUint64(key FieldKey, val uint64) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeUint64, Value: val})
	}
}

func (fields Fields) PutFloat32(key FieldKey, val float32) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeFloat32, Value: val})
	}
}

func (fields Fields) PutFloat64(key FieldKey, val float64) {
	if val != 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeFloat64, Value: val})
	}
}

func (fields Fields) PutString(key FieldKey, val string) {
	if len(val) > 0 {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeString, Value: val})
	}
}

func (fields Fields) PutFields(key FieldKey, val Fields) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeStruct, Value: val})
	}
}

func (fields Fields) PutBoolArray(key FieldKey, val []bool) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeBoolArray, Value: val})
	}
}

func (fields Fields) PutInt8Array(key FieldKey, val []int8) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeInt8Array, Value: val})
	}
}

func (fields Fields) PutInt16Array(key FieldKey, val []int16) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeInt16Array, Value: val})
	}
}

func (fields Fields) PutInt32Array(key FieldKey, val []int32) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeInt32Array, Value: val})
	}
}

func (fields Fields) PutInt64Array(key FieldKey, val []int64) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeInt64Array, Value: val})
	}
}

func (fields Fields) PutUint8Array(key FieldKey, val []uint8) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeUint8Array, Value: val})
	}
}

func (fields Fields) PutUint16Array(key FieldKey, val []uint16) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeUint16Array, Value: val})
	}
}

func (fields Fields) PutUint32Array(key FieldKey, val []uint32) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeUint32Array, Value: val})
	}
}

func (fields Fields) PutUint64Array(key FieldKey, val []uint64) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeUint64Array, Value: val})
	}
}

func (fields Fields) PutFloat32Array(key FieldKey, val []float32) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeFloat32Array, Value: val})
	}
}

func (fields Fields) PutFloat64Array(key FieldKey, val []float64) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeFloat64Array, Value: val})
	}
}

func (fields Fields) PutStringArray(key FieldKey, val []string) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeStringArray, Value: val})
	}
}

func (fields Fields) PutFieldsArray(key FieldKey, val []Fields) {
	if val != nil {
		fields.Put(key, &util.Pair[FieldType, interface{}]{Key: FieldTypeStructArray, Value: val})
	}
}
