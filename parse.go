package fpbs

import (
	"fpbs/errors"
	"fpbs/util"
	"io"
)

func Parse(p []byte) (Fields, error) {
	if !Verify(p) {
		return nil, errors.Error("verify failed")
	}

	var buff = newBuffer2(p)
	var fields, err = parseFields(buff)
	if err != nil {
		return nil, err
	}

	return fields, nil
}

func parseFields(buff *buffer) (Fields, error) {
	var fields Fields
	for {
		key, err := buff.GetByte()
		if err == io.EOF {
			break
		}

		if key == 0 {
			return fields, nil
		}

		if fields == nil {
			fields = Fields{}
		} else {
			if _, ok := fields[FieldKey(key)]; ok {
				return nil, errors.Error("repeated field key")
			}
		}

		tipe, err := buff.GetByte()
		if err != nil {
			return nil, err
		}

		var value interface{}
		switch FieldType(tipe) {
		case FieldTypeBool:
			value, err = buff.GetBool()
			if err != nil {
				return nil, err
			}
		case FieldTypeInt8:
			value, err = buff.GetInt8()
			if err != nil {
				return nil, err
			}
		case FieldTypeInt16:
			value, err = buff.GetInt16()
			if err != nil {
				return nil, err
			}
		case FieldTypeInt32:
			value, err = buff.GetInt32()
			if err != nil {
				return nil, err
			}
		case FieldTypeInt64:
			value, err = buff.GetInt64()
			if err != nil {
				return nil, err
			}
		case FieldTypeUint8:
			value, err = buff.GetUint8()
			if err != nil {
				return nil, err
			}
		case FieldTypeUint16:
			value, err = buff.GetUint16()
			if err != nil {
				return nil, err
			}
		case FieldTypeUint32:
			value, err = buff.GetUint32()
			if err != nil {
				return nil, err
			}
		case FieldTypeUint64:
			value, err = buff.GetUint64()
			if err != nil {
				return nil, err
			}
		case FieldTypeFloat32:
			value, err = buff.GetFloat32()
			if err != nil {
				return nil, err
			}
		case FieldTypeFloat64:
			value, err = buff.GetFloat64()
			if err != nil {
				return nil, err
			}
		case FieldTypeString:
			value, err = buff.GetString()
			if err != nil {
				return nil, err
			}
		case FieldTypeStruct:
			value, err = parseFields(buff)
			if err != nil {
				return nil, err
			}
		case FieldTypeBoolArray:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []bool
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetBool()
				if err != nil {
					s = append(s, v)
					return nil, err
				}
			}
			value = s
		case FieldTypeInt8Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []int8
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetInt8()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeInt16Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []int16
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetInt16()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeInt32Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []int32
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetInt32()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeInt64Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []int64
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetInt64()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeUint8Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []uint8
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetUint8()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeUint16Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []uint16
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetUint16()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeUint32Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []uint32
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetUint32()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeUint64Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []uint64
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetUint64()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeFloat32Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []float32
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetFloat32()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeFloat64Array:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []float64
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetFloat64()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeStringArray:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []string
			for i := uint32(0); i < sl; i++ {
				v, err := buff.GetString()
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		case FieldTypeStructArray:
			sl, err := buff.GetUint32()
			if err != nil {
				return nil, err
			}
			var s []Fields
			for i := uint32(0); i < sl; i++ {
				v, err := parseFields(buff)
				if err != nil {
					return nil, err
				}
				s = append(s, v)
			}
			value = s
		default:
			return nil, errors.Error("unknown field type")
		}
		fields.Put(FieldKey(key), &util.Pair[FieldType, interface{}]{Key: FieldType(tipe), Value: value})
	}

	return fields, nil
}
