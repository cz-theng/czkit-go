package panic

import (
	"testing"
)

func Panic() (err error) {
	defer CatchPanic(&err)
	panic("TestCatchPanic")
}

func TestCatchPanic(t *testing.T) {
	err := Panic()
	t.Log(err)
}
