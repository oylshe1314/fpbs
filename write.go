package fpbs

func Write(fields Fields) []byte {
	var buff = newBuffer1(256)
	writeFields(fields, buff)
	var p = buff.Bytes()
	var code = Verification(p)
	buff.PutByte(code)
	return buff.Bytes()
}

func writeFields(fields Fields, buff *buffer) {
	for key, field := range fields {
		buff.PutByte(byte(key))
		buff.PutByte(byte(field.Key))

		switch field.Key {
		case FieldTypeBool:
			buff.PutBool(field.Value.(bool))
		case FieldTypeInt8:
			buff.PutInt8(field.Value.(int8))
		case FieldTypeInt16:
			buff.PutInt16(field.Value.(int16))
		case FieldTypeInt32:
			buff.PutInt32(field.Value.(int32))
		case FieldTypeInt64:
			buff.PutInt64(field.Value.(int64))
		case FieldTypeUint8:
			buff.PutUint8(field.Value.(uint8))
		case FieldTypeUint16:
			buff.PutUint16(field.Value.(uint16))
		case FieldTypeUint32:
			buff.PutUint32(field.Value.(uint32))
		case FieldTypeUint64:
			buff.PutUint64(field.Value.(uint64))
		case FieldTypeFloat32:
			buff.PutFloat32(field.Value.(float32))
		case FieldTypeFloat64:
			buff.PutFloat64(field.Value.(float64))
		case FieldTypeString:
			buff.PutString(field.Value.(string))
		case FieldTypeStruct:
			writeFields(field.Value.(Fields), buff)
		case FieldTypeBoolArray:
			var s = field.Value.([]bool)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutBool(v)
			}
		case FieldTypeInt8Array:
			var s = field.Value.([]int8)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutInt8(v)
			}
		case FieldTypeInt16Array:
			var s = field.Value.([]int16)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutInt16(v)
			}
		case FieldTypeInt32Array:
			var s = field.Value.([]int32)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutInt32(v)
			}
		case FieldTypeInt64Array:
			var s = field.Value.([]int64)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutInt64(v)
			}
		case FieldTypeUint8Array:
			var s = field.Value.([]uint8)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutUint8(v)
			}
		case FieldTypeUint16Array:
			var s = field.Value.([]uint16)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutUint16(v)
			}
		case FieldTypeUint32Array:
			var s = field.Value.([]uint32)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutUint32(v)
			}
		case FieldTypeUint64Array:
			var s = field.Value.([]uint64)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutUint64(v)
			}
		case FieldTypeFloat32Array:
			var s = field.Value.([]float32)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutFloat32(v)
			}
		case FieldTypeFloat64Array:
			var s = field.Value.([]float64)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				buff.PutFloat64(v)
			}
		case FieldTypeStringArray:
			var ss = field.Value.([]string)
			buff.PutUint32(uint32(len(ss)))
			for i := range ss {
				buff.PutString(ss[i])
			}
		case FieldTypeStructArray:
			var s = field.Value.([]Fields)
			buff.PutUint32(uint32(len(s)))
			for _, v := range s {
				writeFields(v, buff)
			}
		}
	}
	buff.PutByte(byte(0))
	return
}
