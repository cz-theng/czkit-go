// Copyright CZ. All rights reserved.
// Author: CZ cz.theng@gmail.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file

// Package czkit  is cz's code collection for golang developping

package czkit

import (
	"fmt"
)

const (
	major = 0
	minor = 1
	patch = 0
)

// Version return czkit's version
func Version() string {
	return fmt.Sprintf("czkit[%d.%d.%d]", major, minor, patch)
}
