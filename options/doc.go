// Copyright CZ. All rights reserved.
// Author: CZ cz.theng@gmail.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

/**
* Package options implement parse config from a config file and then parse flag
* arguments from command line. flags name according to GNU's format which
* user "-" with an alphabet  for short format, and "--" with a word for long format.
*
* Users should defeine their own option,like
*
* type UsersOptions struct {
* 	Host `toml:"host" flag:"host" sflag:"h" default:"127.0.0.1" help:"TCP host for listen." `
*   Port `toml:"port" flag:"port" sflag:"p" default:"9053" help:"TCP port for listen."`
* }
*
* czoptions will parse Host and Port from your config file, and then parse them from the
* command line with flags(-h for short and --host for long) if they are there. a New method will return  an UserOptions
* with default value of tag "default" .  tag "help" will see in flag's help message
*
* Note: czoptions will take a flag name "-c" "--config"
*
 */
package options
