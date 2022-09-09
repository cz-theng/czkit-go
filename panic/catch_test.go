package panic

import (
	"testing"
)

func causePanic() (err error) {
	defer CatchPanic(&err, "cause panic")
	panic("TestCatchPanic")
}

func TestCatchPanic(t *testing.T) {
	err := causePanic()
	t.Log(err)
}
