// Copyright CZ. All rights reserved.
// Author: CZ cz.theng@gmail.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

package options

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/BurntSushi/toml"
)

var (
	errUnknownStrconvType = errors.New("no a basic type of strconv")
)

const (
	configHelp string = "path for config file"
	helpHelp   string = " for help document"
)

// Parse will parse flag and read config file to option with default values
func Parse(options interface{}) (err error) {

	err = parseDefaultValues(options)
	if err != nil {
		err = fmt.Errorf("parse default value:%w", err)
		fmt.Printf("[ERROR]:%s", err.Error())
		return err
	}

	err = parseConfigFile(options)
	if err != nil {
		err = fmt.Errorf("parse config file: %w", err)
		fmt.Printf("[ERROR]:%s", err.Error())
		return err
	}

	err = parseFlags(options)
	if err != nil {
		err = fmt.Errorf("parse flags:%w", err)
		fmt.Printf("[ERROR]:%s", err.Error())
		return err
	}

	return nil
}

func parseDefaultValues(options interface{}) (err error) {
	val := reflect.ValueOf(options).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		dval := field.Tag.Get("default")
		if dval == "" {
			continue
		}
		fieldVal := val.FieldByName(field.Name)
		v, err := strconvParse(dval, fieldVal.Type().Name())
		if err != nil {
			continue
		}
		fieldVal.Set(reflect.ValueOf(v))
	}
	return nil
}

func strconvParse(str string, typ string) (interface{}, error) {
	switch typ {
	case "string":
		return str, nil
	case "int16":
		v, err := strconv.ParseInt(str, 10, 16)
		if err != nil {
			return int16(0), err
		}
		return int16(v), nil
	case "uint16":
		v, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return uint16(0), err
		}
		return uint16(v), nil
	case "int":
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return int(0), err
		}
		return int(v), nil
	case "uint":
		v, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return uint(0), err
		}
		return uint(v), nil
	case "uint32":
		v, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return uint32(0), err
		}
		return uint32(v), nil
	case "int32":
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return int32(0), err
		}
		return int32(v), nil
	case "int64":
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return int64(0), err
		}
		return int64(v), nil
	case "uint64":
		v, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return uint64(0), err
		}
		return uint64(v), nil
	case "float32":
		v, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return float32(0), err
		}
		return float32(v), nil
	case "float64":
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return float64(0), err
		}
		return float64(v), nil
	default:
		return str, errUnknownStrconvType
	}
	return str, errUnknownStrconvType
}

func parseConfigFile(options interface{}) (err error) {
	configFilePath := ""
	if len(os.Args) < 3 {
		return nil
	}
	for i := 1; i < len(os.Args[1:]); i++ {
		if os.Args[i] == "-c" || os.Args[i] == "--config" || os.Args[i] == "-config" {
			i++
			configFilePath = os.Args[i]
		}
	}

	if configFilePath == "" {
		return nil
	}
	_, err = toml.DecodeFile(configFilePath, options)
	if err != nil {
		return err
	}
	return nil
}

func parseFlags(options interface{}) (err error) {
	val := reflect.ValueOf(options).Elem()
	typ := val.Type()

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var args []string = make([]string, len(os.Args))

	if len(os.Args) > 2 {
		j := 0
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i] == "-c" || os.Args[i] == "--config" || os.Args[i] == "-config" {
				i++
				continue
			}
			args[j] = os.Args[i]
			j++
		}
	} else {
		copy(args, os.Args)
	}

	// for fake parse to build all the FlagSet.formal        map[string]*Flag
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		flagHelp := field.Tag.Get("help")
		dval := field.Tag.Get("default")

		flagKey := field.Tag.Get("flag")
		if flagKey != "" {
			fs.String(flagKey, dval, flagHelp)
		}
		sflagKey := field.Tag.Get("sflag")
		if sflagKey != "" {
			fs.String(sflagKey, dval, flagHelp)
		}
	}
	fs.Parse(args)

	// build option from Flag in FlagSet.formal
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		flagKey := field.Tag.Get("flag")
		sflagKey := field.Tag.Get("sflag")
		if flagKey == "" && sflagKey == "" {
			continue
		}

		fieldVal := val.FieldByName(field.Name)
		sf := fs.Lookup(sflagKey)
		if sf != nil {
			if sf.Value.String() != sf.DefValue {
				sv, _ := strconvParse(sf.Value.String(), fieldVal.Type().Name())
				fieldVal.Set(reflect.ValueOf(sv))
			}
		}
		f := fs.Lookup(flagKey)
		if f.Value.String() != f.DefValue {
			v, _ := strconvParse(f.Value.String(), fieldVal.Type().Name())
			fieldVal.Set(reflect.ValueOf(v))
		}
	}

	return nil
}
