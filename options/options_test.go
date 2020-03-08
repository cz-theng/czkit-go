// Copyright CZ. All rights reserved.
// Author: CZ cz.theng@gmail.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

package options

import (
	"os"
	"testing"
)

type testOptions struct {
	Host string `toml:"host" flag:"host" sflag:"h" default:"127.0.0.1" help:"TCP host for listen." `
	Port uint32 `toml:"port" flag:"port" sflag:"p" default:"9053" help:"TCP port for listen."`
}

type testDefaultOptions struct {
	Host            string `toml:"host" flag:"host" sflag:"h" default:"127.0.0.1" help:"TCP host for listen." `
	Port            int32  `toml:"port" flag:"port" sflag:"p" default:"9053" help:"TCP port for listen."`
	NoDefaultString string `toml:"no_default_string" flag:"no-default-string" sflag:"s" help:"no default string"`
	NoDefaultInt    string `toml:"no_default_int" flag:"no-default" sflag:"i" help:"no default int"`
}

type testArrayMapOptions struct {
	Host       string         `toml:"host" flag:"host" sflag:"h" default:"127.0.0.1" help:"TCP host for listen." `
	Port       uint32         `toml:"port" flag:"port" sflag:"p" default:"9053" help:"TCP port for listen."`
	OrderCMD   []uint32       `toml:"cmd" default:"1 2 3" `
	Hosts      map[string]int `toml:"hosts" default:"127.0.0.1:1 127.0.0.2:2"`
	Dispatcher map[int]string `toml:"dispatcher" default:"1:127.0.0.1 2:127.0.0.2"`
}

type testArrayMapFlagOptions struct {
	Host       string         `toml:"host" flag:"host" sflag:"h" default:"127.0.0.1" help:"TCP host for listen." `
	Port       uint32         `toml:"port" flag:"port" sflag:"p" default:"9053" help:"TCP port for listen."`
	OrderCMD   []uint32       `toml:"cmd" default:"1 2 3"  flag:"ocmd" help:"order cmd"`
	Hosts      map[string]int `toml:"hosts" default:"127.0.0.1:1 127.0.0.2:2" flag:"hosts" help:"hosts"`
	Dispatcher map[int]string `toml:"dispatcher" default:"1:127.0.0.1 2:127.0.0.2" flag:"dispatcher" help:"dispatcher"`
}

func TestParse(t *testing.T) {
	//os.Args = []string{"--help"}
	os.Args = []string{"czkit"}
	opt := &testOptions{}
	Parse(opt)
}

func TestParseConfig(t *testing.T) {
	os.Args = []string{"czkit", "-c", "./testdata/config.toml"}
	opt := &testOptions{}
	Parse(opt)
}

func TestParseDefault(t *testing.T) {
	os.Args = []string{"czkit"}
	opt := &testDefaultOptions{}
	Parse(opt)
}

func TestParseFlag(t *testing.T) {
	os.Args = []string{"czkit", "-c", "./testdata/config.toml", "-p", "9111"}
	opt := &testOptions{}
	Parse(opt)
}

func TestParseArrayMap(t *testing.T) {
	os.Args = []string{"czkit", "-c", "./testdata/config.toml", "-p", "9111"}
	opt := &testArrayMapOptions{}
	Parse(opt)
}

func TestParseArrayMapFlag(t *testing.T) {
	os.Args = []string{"czkit", "-c", "./testdata/config.toml", "-p", "9111", "-ocmd", "4 5 6", "-hosts", "127.0.0.5:5 127.0.0.6:6", "-dispatcher", "7:127.0.0.7 8:127.0.0.8"}
	opt := &testArrayMapFlagOptions{}
	Parse(opt)
}
