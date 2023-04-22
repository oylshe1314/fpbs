package fpbs

import (
	"fpbs/errors"
	"reflect"
	"strconv"
)

func Unmarshal(p []byte, v interface{}) error {
	var vv = reflect.ValueOf(v)
	if vv.Kind() != reflect.Pointer {
		return errors.Error("unmarshal non-pointer")
	}

	if vv.IsNil() {
		return errors.Error("unmarshal null-pointer")
	}

	vv = vv.Elem()

	if vv.Kind() != reflect.Struct {
		return errors.Error("unmarshal non-structure")
	}

	fields, err := Parse(p)
	if err != nil {
		return err
	}

	return readStruct(fields, vv)
}

func readStruct(fields Fields, v reflect.Value) error {
	var vt = v.Type()
	var fn = v.NumField()
	for i := 0; i < fn; i++ {
		var sf = vt.Field(i)
		if !sf.IsExported() {
			continue
		}

		var strKey = sf.Tag.Get("key")
		fieldKey, err := strconv.ParseUint(strKey, 10, 64)
		if err != nil {
			return err
		}

		if fieldKey == 0 || fieldKey > 255 {
			return errors.Error("the range of field tag 'key' is 1-255")
		}

		var pair = fields.Get(FieldKey(fieldKey))
		if pair == nil || pair.Value == nil {
			continue
		}

		var fv = v.Field(i)
		switch pair.Key {
		case FieldTypeBool:
			if sf.Type.Kind() != reflect.Bool {
				return errors.Error("read field non-bool")
			}
			fv.SetBool(pair.Value.(bool))
		case FieldTypeInt8:
			if sf.Type.Kind() != reflect.Int8 {
				return errors.Error("read field non-int8")
			}
			fv.SetInt(int64(pair.Value.(int8)))
		case FieldTypeInt16:
			if sf.Type.Kind() != reflect.Int16 {
				return errors.Error("read field non-int16")
			}
			fv.SetInt(int64(pair.Value.(int16)))
		case FieldTypeInt32:
			if sf.Type.Kind() != reflect.Int32 {
				return errors.Error("read field non-int32")
			}
			fv.SetInt(int64(pair.Value.(int32)))
		case FieldTypeInt64:
			if sf.Type.Kind() != reflect.Int64 {
				return errors.Error("read field non-int64")
			}
			fv.SetInt(pair.Value.(int64))
		case FieldTypeUint8:
			if sf.Type.Kind() != reflect.Uint8 {
				return errors.Error("read field non-uint8")
			}
			fv.SetUint(uint64(pair.Value.(uint8)))
		case FieldTypeUint16:
			if sf.Type.Kind() != reflect.Uint16 {
				return errors.Error("read field non-uint16")
			}
			fv.SetUint(uint64(pair.Value.(uint16)))
		case FieldTypeUint32:
			if sf.Type.Kind() != reflect.Uint32 {
				return errors.Error("read field non-uint32")
			}
			fv.SetUint(uint64(pair.Value.(uint32)))
		case FieldTypeUint64:
			if sf.Type.Kind() != reflect.Uint64 {
				return errors.Error("read field non-uint64")
			}
			fv.SetUint(pair.Value.(uint64))
		case FieldTypeFloat32:
			if sf.Type.Kind() != reflect.Float32 {
				return errors.Error("read field non-float32")
			}
			fv.SetFloat(float64(pair.Value.(float32)))
		case FieldTypeFloat64:
			if sf.Type.Kind() != reflect.Float64 {
				return errors.Error("read field non-float64")
			}
			fv.SetFloat(pair.Value.(float64))
		case FieldTypeString:
			if sf.Type.Kind() != reflect.String {
				return errors.Error("read field non-string")
			}
			fv.SetString(pair.Value.(string))
		case FieldTypeStruct:
			if sf.Type.Kind() != reflect.Pointer && sf.Type.Elem().Kind() != reflect.Struct {
				return errors.Error("read field non-structure-pointer")
			}

			var pv = reflect.New(sf.Type.Elem())
			err = readStruct(pair.Value.(Fields), pv.Elem())
			if err != nil {
				return err
			}
			fv.Set(pv)
		case FieldTypeBoolArray:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Bool {
				return errors.Error("read field non-bool-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]bool)))
		case FieldTypeInt8Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Int8 {
				return errors.Error("read field non-int8-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]int8)))
		case FieldTypeInt16Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Int16 {
				return errors.Error("read field non-int16-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]int16)))
		case FieldTypeInt32Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Int32 {
				return errors.Error("read field non-int32-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]int32)))
		case FieldTypeInt64Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Int64 {
				return errors.Error("read field non-int64-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]int64)))
		case FieldTypeUint8Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Uint8 {
				return errors.Error("read field non-uint8-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]uint8)))
		case FieldTypeUint16Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Uint16 {
				return errors.Error("read field non-uint16-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]uint16)))
		case FieldTypeUint32Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Uint32 {
				return errors.Error("read field non-uint32-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]uint32)))
		case FieldTypeUint64Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Uint64 {
				return errors.Error("read field non-uint64-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]uint64)))
		case FieldTypeFloat32Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Float32 {
				return errors.Error("read field non-float32-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]float32)))
		case FieldTypeFloat64Array:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.Float64 {
				return errors.Error("read field non-float64-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]float64)))
		case FieldTypeStringArray:
			if sf.Type.Kind() != reflect.Slice && sf.Type.Elem().Kind() != reflect.String {
				return errors.Error("read field non-string-slice")
			}
			fv.Set(reflect.ValueOf(pair.Value.([]string)))
		case FieldTypeStructArray:
			if sf.Type.Kind() != reflect.Slice {
				return errors.Error("read field non-slice")
			}
			var eft = sf.Type.Elem()
			if eft.Kind() != reflect.Pointer && eft.Elem().Kind() != reflect.Struct {
				return errors.Error("read field non-structure-pointer-slice")
			}

			var fieldsArray = pair.Value.([]Fields)
			var sv = reflect.MakeSlice(sf.Type.Elem(), len(fieldsArray), len(fieldsArray))
			for ai, af := range fieldsArray {
				var ipv = reflect.New(eft.Elem())
				err = readStruct(af, ipv.Elem())
				if err != nil {
					return err
				}

				sv.Index(ai).Set(ipv)
			}
			fv.Set(sv)
		}
	}
	return nil
}
