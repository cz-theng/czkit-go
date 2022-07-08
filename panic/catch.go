package panic

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
)

var panicWriter io.Writer = os.Stderr

// CatchPanic catch panic and logs panics.
func CatchPanic(err *error) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	r := recover()
	fmt.Fprintf(panicWriter, "cauth panic: %v\n%s\n", r, debug.Stack())
	*err = fmt.Errorf("panic at [%s:%d]: %s", file, line, r)
}
