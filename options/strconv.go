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
		e := typ.Elem()
		return strconvParseSlice(str, e)
	case reflect.Map:
		k := typ.Key()
		e := typ.Elem()
		return strconvParseMap(str, k, e)
	default:
		return str, errUnknownStrconvType
	}
}

func strconvParseSlice(str string, typ reflect.Type) (interface{}, error) {
	s := strings.Split(str, " ")
	switch typ.Kind() {
	case reflect.Bool:
		ss := make([]bool, len(s))
		for i := 0; i < len(s); i++ {
			ss[i], _ = strconv.ParseBool(s[i])
		}
		return ss, nil

	case reflect.Int:
		ss := make([]int, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseInt(s[i], 10, 64)
			ss[i] = int(v)
		}
		return ss, nil

	case reflect.Uint:
		ss := make([]uint, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseUint(s[i], 10, 64)
			ss[i] = uint(v)
		}
		return ss, nil

	case reflect.Int16:
		ss := make([]int16, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseInt(s[i], 10, 16)
			ss[i] = int16(v)
		}
		return ss, nil

	case reflect.Uint16:
		ss := make([]uint16, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseUint(s[i], 10, 16)
			ss[i] = uint16(v)
		}
		return ss, nil

	case reflect.Int32:
		ss := make([]int32, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseInt(s[i], 10, 32)
			ss[i] = int32(v)
		}
		return ss, nil

	case reflect.Uint32:
		ss := make([]uint32, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseUint(s[i], 10, 32)
			ss[i] = uint32(v)
		}
		return ss, nil

	case reflect.Int64:
		ss := make([]int64, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseInt(s[i], 10, 64)
			ss[i] = int64(v)
		}
		return ss, nil

	case reflect.Uint64:
		ss := make([]uint64, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseUint(s[i], 10, 64)
			ss[i] = uint64(v)
		}
		return ss, nil

	case reflect.Float32:
		ss := make([]float32, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseFloat(s[i], 10)
			ss[i] = float32(v)
		}
		return ss, nil

	case reflect.Float64:
		ss := make([]float64, len(s))
		for i := 0; i < len(s); i++ {
			v, _ := strconv.ParseFloat(s[i], 10)
			ss[i] = float64(v)
		}
		return ss, nil

	case reflect.String:
		ss := make([]string, len(s))
		for i := 0; i < len(s); i++ {
			ss[i] = s[i]
		}
		return ss, nil

	default:
		return str, errUnknownStrconvType
	}
}

func strconvParseMap(str string, ktyp reflect.Type, vtyp reflect.Type) (interface{}, error) {
	s := strings.Split(str, " ")
	var kv []string
	for _, e := range s {
		item := strings.Split(e, ":")
		if len(item) != 2 {
			continue
		}
		kv = append(kv, item...)
	}

	/** key is string **/
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.String {
		m := make(map[string]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v := kv[2*i+1]
			m[k] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Bool {
		m := make(map[string]bool)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Int {
		m := make(map[string]int)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Uint {
		m := make(map[string]uint)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Int16 {
		m := make(map[string]int16)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Uint16 {
		m := make(map[string]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Int32 {
		m := make(map[string]int32)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Uint32 {
		m := make(map[string]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Int64 {
		m := make(map[string]int64)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Uint64 {
		m := make(map[string]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Float32 {
		m := make(map[string]float32)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.String && vtyp.Kind() == reflect.Float64 {
		m := make(map[string]float64)
		for i := 0; i < len(kv)/2; i++ {
			k := kv[2*i]
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k] = v.(float64)
		}
		return m, nil
	}

	/** key is int **/
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.String {
		m := make(map[int]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(int)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Bool {
		m := make(map[int]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Int {
		m := make(map[int]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Uint {
		m := make(map[int]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Int16 {
		m := make(map[int]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Uint16 {
		m := make(map[int]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Int32 {
		m := make(map[int]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Uint32 {
		m := make(map[int]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Int64 {
		m := make(map[int]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Uint64 {
		m := make(map[int]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Float32 {
		m := make(map[int]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Float64 {
		m := make(map[int]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int)] = v.(float64)
		}
		return m, nil
	}

	/** key is uint **/
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.String {
		m := make(map[uint]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(uint)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Bool {
		m := make(map[uint]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Int {
		m := make(map[uint]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Uint {
		m := make(map[uint]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Int16 {
		m := make(map[uint]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Uint16 {
		m := make(map[uint]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Int32 {
		m := make(map[uint]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Uint32 {
		m := make(map[uint]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Int64 {
		m := make(map[uint]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Uint64 {
		m := make(map[uint]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Float32 {
		m := make(map[uint]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Int && vtyp.Kind() == reflect.Float64 {
		m := make(map[uint]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint)] = v.(float64)
		}
		return m, nil
	}

	/** key is int16 **/
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.String {
		m := make(map[int16]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(int16)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Bool {
		m := make(map[int16]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Int {
		m := make(map[int16]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Uint {
		m := make(map[int16]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Int16 {
		m := make(map[int16]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Uint16 {
		m := make(map[int16]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Int32 {
		m := make(map[int16]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Uint32 {
		m := make(map[int16]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Int64 {
		m := make(map[int16]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Uint64 {
		m := make(map[int16]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Float32 {
		m := make(map[int16]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Int16 && vtyp.Kind() == reflect.Float64 {
		m := make(map[int16]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int16)] = v.(float64)
		}
		return m, nil
	}

	/** key is uint16 **/
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.String {
		m := make(map[uint16]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(uint16)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Bool {
		m := make(map[uint16]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Int {
		m := make(map[uint16]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Uint {
		m := make(map[uint16]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Uint16 {
		m := make(map[uint16]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Int16 {
		m := make(map[uint16]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Int32 {
		m := make(map[uint16]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Uint32 {
		m := make(map[uint16]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Int64 {
		m := make(map[uint16]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Uint64 {
		m := make(map[uint16]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Float32 {
		m := make(map[uint16]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Uint16 && vtyp.Kind() == reflect.Float64 {
		m := make(map[uint16]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint16)] = v.(float64)
		}
		return m, nil
	}

	/** key is int32 **/
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.String {
		m := make(map[int32]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(int32)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Bool {
		m := make(map[int32]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Int {
		m := make(map[int32]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Uint {
		m := make(map[int32]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Int16 {
		m := make(map[int32]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Uint16 {
		m := make(map[int32]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Int32 {
		m := make(map[int32]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Uint32 {
		m := make(map[int32]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Int64 {
		m := make(map[int32]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Uint64 {
		m := make(map[int32]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Float32 {
		m := make(map[int32]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Int32 && vtyp.Kind() == reflect.Float64 {
		m := make(map[int32]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int32)] = v.(float64)
		}
		return m, nil
	}

	/** key is uint32 **/
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.String {
		m := make(map[uint32]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(uint32)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Bool {
		m := make(map[uint32]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Int {
		m := make(map[uint32]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Uint {
		m := make(map[uint32]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Int16 {
		m := make(map[uint32]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Uint16 {
		m := make(map[uint32]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Int32 {
		m := make(map[uint32]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Uint32 {
		m := make(map[uint32]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Int64 {
		m := make(map[uint32]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Uint64 {
		m := make(map[uint32]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Float32 {
		m := make(map[uint32]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Uint32 && vtyp.Kind() == reflect.Float64 {
		m := make(map[uint32]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint32)] = v.(float64)
		}
		return m, nil
	}

	/** key is int64 **/
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.String {
		m := make(map[int64]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(int64)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Bool {
		m := make(map[int64]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Int {
		m := make(map[int64]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Uint {
		m := make(map[int64]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Int16 {
		m := make(map[int64]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Uint16 {
		m := make(map[int64]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Int32 {
		m := make(map[int64]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Uint32 {
		m := make(map[int64]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Int64 {
		m := make(map[int64]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Uint64 {
		m := make(map[int64]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Float32 {
		m := make(map[int64]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Int64 && vtyp.Kind() == reflect.Float64 {
		m := make(map[int64]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(int64)] = v.(float64)
		}
		return m, nil
	}

	/** key is uint64 **/
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.String {
		m := make(map[uint64]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(uint64)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Bool {
		m := make(map[uint64]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Int {
		m := make(map[uint64]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Uint {
		m := make(map[uint64]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Int64 {
		m := make(map[uint64]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Uint16 {
		m := make(map[uint64]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Int32 {
		m := make(map[uint64]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Uint32 {
		m := make(map[uint64]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Int64 {
		m := make(map[uint64]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Uint64 {
		m := make(map[uint64]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Float32 {
		m := make(map[uint64]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Uint64 && vtyp.Kind() == reflect.Float64 {
		m := make(map[uint64]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(uint64)] = v.(float64)
		}
		return m, nil
	}

	/** key is float32 **/
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.String {
		m := make(map[float32]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(float32)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Bool {
		m := make(map[float32]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Int {
		m := make(map[float32]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Uint {
		m := make(map[float32]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Int16 {
		m := make(map[float32]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Uint16 {
		m := make(map[float32]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Int32 {
		m := make(map[float32]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Uint32 {
		m := make(map[float32]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Int64 {
		m := make(map[float32]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Uint64 {
		m := make(map[float32]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Float32 {
		m := make(map[float32]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Float32 && vtyp.Kind() == reflect.Float64 {
		m := make(map[float32]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float32)] = v.(float64)
		}
		return m, nil
	}

	/** key is float64 **/
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.String {
		m := make(map[float64]string, len(s))
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v := kv[2*i+1]
			if err != nil {
				continue
			}
			m[k.(float64)] = v
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Bool {
		m := make(map[float64]bool)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(bool)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Int {
		m := make(map[float64]int)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(int)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Uint {
		m := make(map[float64]uint)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(uint)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Int16 {
		m := make(map[float64]int16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(int16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Uint16 {
		m := make(map[float64]uint16)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(uint16)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Int32 {
		m := make(map[float64]int32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(int32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Uint32 {
		m := make(map[float64]uint32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(uint32)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Int64 {
		m := make(map[float64]int64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(int64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Uint64 {
		m := make(map[float64]uint64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(uint64)
		}
		return m, nil
	}
	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Float32 {
		m := make(map[float64]float32)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(float32)
		}
		return m, nil
	}

	if ktyp.Kind() == reflect.Float64 && vtyp.Kind() == reflect.Float64 {
		m := make(map[float64]float64)
		for i := 0; i < len(kv)/2; i++ {
			k, err := strconvParse(kv[2*i], ktyp)
			v, err := strconvParse(kv[2*i+1], vtyp)
			if err != nil {
				continue
			}
			m[k.(float64)] = v.(float64)
		}
		return m, nil
	}
	return str, errUnknownStrconvType
}
