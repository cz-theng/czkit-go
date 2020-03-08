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
	fmt.Printf("parseDefaultValues:%v\n", options)

	err = parseConfigFile(options)
	if err != nil {
		err = fmt.Errorf("parse config file: %w", err)
		fmt.Printf("[ERROR]:%s", err.Error())
		return err
	}
	fmt.Printf("parseConfigFile:%v\n", options)

	err = parseFlags(options)
	if err != nil {
		err = fmt.Errorf("parse flags:%w", err)
		fmt.Printf("[ERROR]:%s", err.Error())
		return err
	}
	fmt.Printf("parseFlags:%v\n", options)

	return nil
}

func parseDefaultValues(options interface{}) (err error) {
	val := reflect.ValueOf(options).Elem()
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		dval := field.Tag.Get("default")
		fieldVal := val.FieldByName(field.Name)
		v, err := strconvParse(dval, fieldVal.Type())
		if err != nil {
			//  Recursively resolve embedded types.
			var fieldPtr reflect.Value
			fmt.Printf("fieldVal:%v  Elem:\n", fieldVal.Type())
			if fieldVal.Type().Kind() == reflect.Struct {
				fieldPtr = fieldVal.Addr()
				if !fieldPtr.IsNil() {
					err = parseDefaultValues(fieldPtr.Interface())
				}
			}

			continue
		}
		fieldVal.Set(reflect.ValueOf(v))
	}
	return nil
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
		fs.Visit(func(f *flag.Flag) {
			if f.Name == flagKey || f.Name == sflagKey {
				sv, _ := strconvParse(f.Value.String(), fieldVal.Type())
				fieldVal.Set(reflect.ValueOf(sv))
			}
		})
	}

	return nil
}
