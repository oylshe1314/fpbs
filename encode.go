package fpbs

import (
	"fpbs/errors"
	"fpbs/util"
	"reflect"
	"strconv"
)

func Marshal(v interface{}) ([]byte, error) {
	var vv = reflect.ValueOf(v)

	if vv.Kind() != reflect.Pointer {
		return nil, errors.Error("marshal non-pointer")
	}

	if vv.IsNil() {
		return nil, errors.Error("marshal nil-pointer")
	}

	vv = vv.Elem()

	if vv.Kind() != reflect.Struct {
		return nil, errors.Error("marshal non-structure")
	}

	fields, err := writeStruct(vv)
	if err != nil {
		return nil, err
	}

	return Write(fields), nil
}

func writeStruct(v reflect.Value) (Fields, error) {
	var vt = v.Type()
	var fn = v.NumField()
	var fields = Fields{}
	for i := 0; i < fn; i++ {
		var sf = vt.Field(i)
		if !sf.IsExported() {
			continue
		}

		var strKey = sf.Tag.Get("key")
		fieldKey, err := strconv.ParseUint(strKey, 10, 64)
		if err != nil {
			return nil, err
		}

		if fieldKey == 0 || fieldKey > 255 {
			return nil, errors.Error("the range of field tag 'key' is 1-255")
		}

		var fv = v.Field(i)
		switch sf.Type.Kind() {
		case reflect.Bool:
			fields.PutBool(FieldKey(fieldKey), fv.Bool())
		case reflect.Int:
			fields.PutInt64(FieldKey(fieldKey), fv.Int())
		case reflect.Int8:
			fields.PutInt8(FieldKey(fieldKey), int8(fv.Int()))
		case reflect.Int16:
			fields.PutInt16(FieldKey(fieldKey), int16(fv.Int()))
		case reflect.Int32:
			fields.PutInt32(FieldKey(fieldKey), int32(fv.Int()))
		case reflect.Int64:
			fields.PutInt64(FieldKey(fieldKey), fv.Int())
		case reflect.Uint:
			fields.PutUint64(FieldKey(fieldKey), fv.Uint())
		case reflect.Uint8:
			fields.PutUint8(FieldKey(fieldKey), uint8(fv.Uint()))
		case reflect.Uint16:
			fields.PutUint16(FieldKey(fieldKey), uint16(fv.Uint()))
		case reflect.Uint32:
			fields.PutUint32(FieldKey(fieldKey), uint32(fv.Uint()))
		case reflect.Uint64:
			fields.PutUint64(FieldKey(fieldKey), fv.Uint())
		case reflect.Float32:
			fields.PutFloat32(FieldKey(fieldKey), float32(fv.Float()))
		case reflect.Float64:
			fields.PutFloat64(FieldKey(fieldKey), fv.Float())
		case reflect.String:
			fields.PutString(FieldKey(fieldKey), fv.String())
		case reflect.Pointer:
			if sf.Type.Elem().Kind() != reflect.Struct {
				return nil, errors.Error("write field non-structure-pointer")
			}

			subFields, err := writeStruct(fv.Elem())
			if err != nil {
				return nil, err
			}

			fields.PutFields(FieldKey(fieldKey), subFields)
		case reflect.Slice:
			var eft = sf.Type.Elem()
			switch eft.Kind() {
			case reflect.Bool:
				fields.PutBoolArray(FieldKey(fieldKey), fv.Interface().([]bool))
			case reflect.Int:
				var array []int64
				util.NumbersConvert1(fv.Interface().([]int), &array)
				fields.PutInt64Array(FieldKey(fieldKey), array)
			case reflect.Int8:
				fields.PutInt8Array(FieldKey(fieldKey), fv.Interface().([]int8))
			case reflect.Int16:
				fields.PutInt16Array(FieldKey(fieldKey), fv.Interface().([]int16))
			case reflect.Int32:
				fields.PutInt32Array(FieldKey(fieldKey), fv.Interface().([]int32))
			case reflect.Int64:
				fields.PutInt64Array(FieldKey(fieldKey), fv.Interface().([]int64))
			case reflect.Uint:
				var array []uint64
				util.NumbersConvert1(fv.Interface().([]uint), &array)
				fields.PutUint64Array(FieldKey(fieldKey), array)
			case reflect.Uint8:
				fields.PutUint8Array(FieldKey(fieldKey), fv.Interface().([]uint8))
			case reflect.Uint16:
				fields.PutUint16Array(FieldKey(fieldKey), fv.Interface().([]uint16))
			case reflect.Uint32:
				fields.PutUint32Array(FieldKey(fieldKey), fv.Interface().([]uint32))
			case reflect.Uint64:
				fields.PutUint64Array(FieldKey(fieldKey), fv.Interface().([]uint64))
			case reflect.Float32:
				fields.PutFloat32Array(FieldKey(fieldKey), fv.Interface().([]float32))
			case reflect.Float64:
				fields.PutFloat64Array(FieldKey(fieldKey), fv.Interface().([]float64))
			case reflect.String:
				fields.PutStringArray(FieldKey(fieldKey), fv.Interface().([]string))
			case reflect.Pointer:
				if eft.Elem().Kind() != reflect.Struct {
					return nil, errors.Error("write field non-structure-pointer-slice")
				}

				var fieldArray = make([]Fields, fv.Len())
				for ai := 0; ai < fv.Len(); ai++ {
					var ipv = fv.Index(ai)
					af, err := writeStruct(ipv.Elem())
					if err != nil {
						return nil, err
					}

					fieldArray[ai] = af
				}
				fields.PutFieldsArray(FieldKey(fieldKey), fieldArray)
			case reflect.Slice:
				sf.Type.Elem().Kind()
			default:
				return nil, errors.Error("write struct unsupported element type of slice field")
			}
		default:
			return nil, errors.Error("write struct unsupported field type")
		}
	}
	return fields, nil
}
