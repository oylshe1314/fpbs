package fpbs

import (
	"encoding/binary"
	"fpbs/errors"
	"io"
	"math"
)

type buffer struct {
	buf []byte
	len uint32
	pos uint32

	end bool
}

func newBuffer1(len uint32) *buffer {
	if len < 64 {
		len = 64
	}
	return &buffer{buf: make([]byte, len), len: len, pos: 0}
}

func newBuffer2(buf []byte) *buffer {
	return &buffer{buf: buf, len: uint32(len(buf)), pos: 0}
}

func (this *buffer) remake(al uint32) {
	if this.pos+al > this.len {
		var nl = this.len * (((this.pos + al) / this.len) + 1)
		var nb = make([]byte, nl)
		copy(nb, this.buf)
		this.buf = nb
		this.len = nl
	}
}

func (this *buffer) Bytes() []byte {
	if this.pos >= this.len {
		return this.buf
	} else {
		return this.buf[:this.pos]
	}
}

func (this *buffer) Len() uint32 {
	return this.len
}

func (this *buffer) Pos() uint32 {
	return this.pos
}

func (this *buffer) PutByte(v byte) {
	this.remake(1)
	this.buf[this.pos] = v
	this.pos += 1
}

func (this *buffer) PutBool(v bool) {
	if v {
		this.PutByte(1)
	} else {
		this.PutByte(0)
	}
}

func (this *buffer) PutInt(v int) {
	this.PutInt64(int64(v))
}

func (this *buffer) PutInt8(v int8) {
	this.PutByte(byte(v))
}

func (this *buffer) PutInt16(v int16) {
	this.PutInt64(int64(v))
}

func (this *buffer) PutInt32(v int32) {
	this.PutInt64(int64(v))
}

func (this *buffer) PutInt64(v int64) {
	this.remake(binary.MaxVarintLen64)
	this.pos += uint32(binary.PutVarint(this.buf[this.pos:], v))
}

func (this *buffer) PutUint(v uint) {
	this.PutUint64(uint64(v))
}

func (this *buffer) PutUint8(v uint8) {
	this.PutByte(v)
}

func (this *buffer) PutUint16(v uint16) {
	this.PutUint64(uint64(v))
}

func (this *buffer) PutUint32(v uint32) {
	this.PutUint64(uint64(v))
}

func (this *buffer) PutUint64(v uint64) {
	this.remake(binary.MaxVarintLen64)
	this.pos += uint32(binary.PutUvarint(this.buf[this.pos:], v))
}

func (this *buffer) PutFloat32(v float32) {
	this.PutUint32(math.Float32bits(v))
}

func (this *buffer) PutFloat64(v float64) {
	this.PutUint64(math.Float64bits(v))
}

func (this *buffer) PutString(v string) {
	this.PutBytes([]byte(v))
}

func (this *buffer) PutBytes(v []byte) {
	var pl = uint32(len(v))
	this.PutUint32(pl)

	if pl == 0 {
		return
	}

	this.remake(pl)
	copy(this.buf[this.pos:], v)
	this.pos += pl
}

func (this *buffer) GetByte() (v byte, err error) {
	if this.pos+1 > this.len {
		return 0, io.EOF
	}
	v = this.buf[this.pos]
	this.pos += 1
	return
}

func (this *buffer) GetBool() (v bool, err error) {
	b, err := this.GetByte()
	return b != 0, err
}

func (this *buffer) GetInt() (int, error) {
	v, err := this.GetInt64()
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (this *buffer) GetInt8() (v int8, err error) {
	b, err := this.GetByte()
	return int8(b), err
}

func (this *buffer) GetInt16() (int16, error) {
	v, err := this.GetInt64()
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

func (this *buffer) GetInt32() (int32, error) {
	v, err := this.GetInt64()
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func (this *buffer) GetInt64() (int64, error) {
	if this.pos+1 > this.len {
		return 0, io.EOF
	}
	v, l := binary.Varint(this.buf[this.pos:])
	if l < 0 {
		return 0, errors.Error("overflow")
	}
	this.pos += uint32(l)
	return v, nil
}

func (this *buffer) GetUint() (uint, error) {
	v, err := this.GetUint64()
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func (this *buffer) GetUint8() (uint8, error) {
	return this.GetByte()
}

func (this *buffer) GetUint16() (uint16, error) {
	v, err := this.GetUint64()
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

func (this *buffer) GetUint32() (uint32, error) {
	v, err := this.GetUint64()
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func (this *buffer) GetUint64() (uint64, error) {
	if this.pos+1 > this.len {
		return 0, io.EOF
	}
	v, l := binary.Uvarint(this.buf[this.pos:])
	if l < 0 {
		return 0, errors.Error("overflow")
	}
	this.pos += uint32(l)
	return v, nil
}

func (this *buffer) GetFloat32() (float32, error) {
	v, err := this.GetUint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(v), nil
}

func (this *buffer) GetFloat64() (float64, error) {
	v, err := this.GetUint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(v), nil
}

func (this *buffer) GetString() (v string, err error) {
	bs, err := this.GetBytes()
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (this *buffer) GetBytes() (v []byte, err error) {
	var l uint32
	l, err = this.GetUint32()
	if err != nil {
		return nil, err
	}

	if l == 0 {
		return nil, nil
	}

	if this.pos+l > this.len {
		return nil, io.EOF
	}

	v = this.buf[this.pos : this.pos+l]
	this.pos += l
	return
}
