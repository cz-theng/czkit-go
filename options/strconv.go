// Copyright CZ. All rights reserved.
// Author: CZ cz.theng@gmail.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

package options

import (
	"reflect"
	"strconv"
	"strings"
)

func strconvParse(str string, typ reflect.Type) (interface{}, error) {
	switch typ.Kind() {
	case reflect.Bool:
		v, err := strconv.ParseBool(str)
		if err != nil {
			return false, err
		}
		return bool(v), nil

	case reflect.String:
		return str, nil
	case reflect.Int16:
		v, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return int16(0), err
		}
		return int16(v), nil
	case reflect.Uint16:
		v, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return uint16(0), err
		}
		return uint16(v), nil
	case reflect.Int:
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return int(0), err
		}
		return int(v), nil
	case reflect.Uint:
		v, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return uint(0), err
		}
		return uint(v), nil
	case reflect.Uint32:
		v, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return uint32(0), err
		}
		return uint32(v), nil
	case reflect.Int32:
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return int32(0), err
		}
		return int32(v), nil
	case reflect.Int64:
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return int64(0), err
		}
		return int64(v), nil
	case reflect.Uint64:
		v, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return uint64(0), err
		}
		return uint64(v), nil
	case reflect.Float32:
		v, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return float32(0), err
		}
		return float32(v), nil
	case reflect.Float64:
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return float64(0), err
		}
		return float64(v), nil
	case reflect.Array, reflect.Slice:
		return strconvParseSlice(str, typ)
	case reflect.Map:
		return strconvParseMap(str, typ)
	default:
		return str, errUnknownStrconvType
	}
}

func strconvParseSlice(str string, typ reflect.Type) (interface{}, error) {
	s := strings.Split(str, " ")
	rs := reflect.MakeSlice(typ, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch typ.Elem().Kind() {
		case reflect.String:
			v := s[i]
			rs = reflect.Append(rs, reflect.ValueOf(v))
		case reflect.Bool:
			v, _ := strconv.ParseBool(s[i])
			reflect.Append(rs, reflect.ValueOf(v))
		case reflect.Int:
			v, _ := strconv.ParseInt(s[i], 10, 64)
			rs = reflect.Append(rs, reflect.ValueOf(int(v)))
		case reflect.Uint:
			v, _ := strconv.ParseInt(s[i], 10, 64)
			rs = reflect.Append(rs, reflect.ValueOf(uint(v)))
		case reflect.Int16:
			v, _ := strconv.ParseInt(s[i], 10, 16)
			rs = reflect.Append(rs, reflect.ValueOf(int16(v)))
		case reflect.Uint16:
			v, _ := strconv.ParseInt(s[i], 10, 16)
			rs = reflect.Append(rs, reflect.ValueOf(uint16(v)))
		case reflect.Int32:
			v, _ := strconv.ParseInt(s[i], 10, 32)
			rs = reflect.Append(rs, reflect.ValueOf(int32(v)))
		case reflect.Uint32:
			v, _ := strconv.ParseInt(s[i], 10, 32)
			rs = reflect.Append(rs, reflect.ValueOf(uint32(v)))
		case reflect.Int64:
			v, _ := strconv.ParseInt(s[i], 10, 64)
			rs = reflect.Append(rs, reflect.ValueOf(int64(v)))
		case reflect.Uint64:
			v, _ := strconv.ParseInt(s[i], 10, 64)
			rs = reflect.Append(rs, reflect.ValueOf(uint64(v)))
		case reflect.Float32:
			v, _ := strconv.ParseFloat(s[i], 10)
			rs = reflect.Append(rs, reflect.ValueOf(float32(v)))
		case reflect.Float64:
			v, _ := strconv.ParseFloat(s[i], 10)
			rs = reflect.Append(rs, reflect.ValueOf(float64(v)))
		}
	}
	return rs.Interface(), nil
}

func strconvParseMap(str string, typ reflect.Type) (interface{}, error) {
	s := strings.Split(str, " ")
	var kv []string
	for _, e := range s {
		item := strings.Split(e, ":")
		if len(item) != 2 {
			continue
		}
		kv = append(kv, item...)
	}

	rm := reflect.MakeMapWithSize(typ, len(kv)/2)
	var key reflect.Value
	var value reflect.Value
	for i := 0; i < len(kv)/2; i++ {
		keyStr := kv[2*i]
		switch typ.Key().Kind() {
		case reflect.String:
			key = reflect.ValueOf(keyStr)
		case reflect.Int:
			v, _ := strconv.ParseInt(keyStr, 10, 64)
			key = reflect.ValueOf(int(v))
		case reflect.Uint:
			v, _ := strconv.ParseInt(keyStr, 10, 64)
			key = reflect.ValueOf(uint(v))
		case reflect.Int16:
			v, _ := strconv.ParseInt(keyStr, 10, 16)
			key = reflect.ValueOf(int16(v))
		case reflect.Uint16:
			v, _ := strconv.ParseInt(keyStr, 10, 16)
			key = reflect.ValueOf(uint16(v))
		case reflect.Int32:
			v, _ := strconv.ParseInt(keyStr, 10, 32)
			key = reflect.ValueOf(int32(v))
		case reflect.Uint32:
			v, _ := strconv.ParseInt(keyStr, 10, 32)
			key = reflect.ValueOf(uint32(v))
		case reflect.Int64:
			v, _ := strconv.ParseInt(keyStr, 10, 64)
			key = reflect.ValueOf(int64(v))
		case reflect.Uint64:
			v, _ := strconv.ParseInt(keyStr, 10, 64)
			key = reflect.ValueOf(uint64(v))
		case reflect.Float32:
			v, _ := strconv.ParseFloat(keyStr, 10)
			key = reflect.ValueOf(float32(v))
		case reflect.Float64:
			v, _ := strconv.ParseFloat(keyStr, 10)
			key = reflect.ValueOf(float64(v))
		}
	}

	for i := 0; i < len(kv)/2; i++ {
		valStr := kv[2*i+1]
		switch typ.Elem().Kind() {
		case reflect.String:
			value = reflect.ValueOf(valStr)
		case reflect.Int:
			v, _ := strconv.ParseInt(valStr, 10, 64)
			value = reflect.ValueOf(int(v))
		case reflect.Uint:
			v, _ := strconv.ParseInt(valStr, 10, 64)
			value = reflect.ValueOf(uint(v))
		case reflect.Int16:
			v, _ := strconv.ParseInt(valStr, 10, 16)
			value = reflect.ValueOf(int16(v))
		case reflect.Uint16:
			v, _ := strconv.ParseInt(valStr, 10, 16)
			value = reflect.ValueOf(uint16(v))
		case reflect.Int32:
			v, _ := strconv.ParseInt(valStr, 10, 32)
			value = reflect.ValueOf(int32(v))
		case reflect.Uint32:
			v, _ := strconv.ParseInt(valStr, 10, 32)
			value = reflect.ValueOf(uint32(v))
		case reflect.Int64:
			v, _ := strconv.ParseInt(valStr, 10, 64)
			value = reflect.ValueOf(int64(v))
		case reflect.Uint64:
			v, _ := strconv.ParseInt(valStr, 10, 64)
			value = reflect.ValueOf(uint64(v))
		case reflect.Float32:
			v, _ := strconv.ParseFloat(valStr, 10)
			value = reflect.ValueOf(float32(v))
		case reflect.Float64:
			v, _ := strconv.ParseFloat(valStr, 10)
			value = reflect.ValueOf(float64(v))
		}

		rm.SetMapIndex(key, value)
	}
	return rm.Interface(), nil
}
