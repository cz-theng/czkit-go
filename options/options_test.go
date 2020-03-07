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
