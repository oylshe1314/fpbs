package util

import (
	"fpbs/errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type UrlValue string

func (urlValue UrlValue) SetValue(v reflect.Value) error {
	var value = string(urlValue)
	switch v.Kind() {
	case reflect.String:
		v.Set(reflect.ValueOf(value))
	case reflect.Bool:
		ev, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(ev)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		ev, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(ev)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		ev, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetUint(ev)
	case reflect.Float32, reflect.Float64:
		ev, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		v.SetFloat(ev)
	default:
		return errors.Errorf("unsupported type '%s'", v.Type().String())
	}
	return nil
}

type UrlValues url.Values

func (urlValues UrlValues) Read(v interface{}) error {
	var vt = reflect.TypeOf(v)
	if vt.Kind() != reflect.Pointer {
		return errors.Error("read get query: non-pointer")
	}

	vt = vt.Elem()
	if vt.Kind() != reflect.Struct {
		return errors.Error("read get query: non-struct")
	}

	var vv = reflect.ValueOf(v).Elem()
	var fn = vt.NumField()
	for i := 0; i < fn; i++ {
		var f = vt.Field(i)
		var name = f.Tag.Get("json")
		if name == "" {
			name = f.Name
		}

		value, ok := urlValues[name]
		if !ok || len(value) == 0 {
			continue
		}

		var vl = len(value)
		var fv = vv.Field(i)
		switch f.Type.Kind() {
		case reflect.Slice:
			fv.Set(reflect.MakeSlice(f.Type, vl, vl))
			fallthrough
		case reflect.Array:
			if vl > fv.Len() {
				vl = fv.Len()
			}
			for fi := 0; fi < vl; fi++ {
				var ev = fv.Index(fi)
				var err = UrlValue(value[fi]).SetValue(ev)
				if err != nil {
					return errors.Errorf("can not set the value '%s' to index %d of the array field '%s', %v", value[fi], fi, f.Name, err)
				}
			}
		default:
			var err = UrlValue(value[0]).SetValue(fv)
			if err != nil {
				return errors.Errorf("can not set the value '%s' to the array field '%s', %v", value[0], f.Name, err)
			}
		}
	}

	return nil
}

func Filename(url string) string {
	var filename = url
	var i = strings.LastIndex(filename, "/")
	if i < 0 {
		i = strings.LastIndex(filename, "\\")
	}
	if i > 0 {
		filename = filename[i+1:]
	}

	i = strings.Index(filename, ".")
	if i > 0 {
		filename = filename[:i]
	}
	return filename
}
