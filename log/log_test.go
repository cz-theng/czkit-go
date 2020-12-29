package log

import (
	"testing"
)

func Test(t *testing.T) {
	Init()
	Infow("failed to connect to ",
		"URL", "https://www.baidu.com",
		"attempt", 3)
}
