package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestDefaultOptions(t *testing.T) {
	opts := defaultOptions()
	if opts.console != true ||
		opts.compress != false ||
		opts.logLevel != DebugLevel {
		t.Fatal("default options not default")
	}
}

var tests = []struct {
	path     string
	name     string
	level    zapcore.Level
	max      uint32
	compress bool
	console  bool
}{
	{
		"/",
		"root.log",
		zap.DebugLevel,
		1024,
		true,
		true,
	},
	{
		"/abc/",
		"abc.log",
		zap.InfoLevel,
		10240,
		false,
		true,
	},
}

func TestWithOptions(t *testing.T) {
	for _, tt := range tests {
		opts := defaultOptions()
		opt := WithLogPath(tt.path)
		opt.apply(&opts)
		opt = WithLogName(tt.name)
		opt.apply(&opts)
		opt = WithMaxLogSize(tt.max)
		opt.apply(&opts)
		opt = WithLogLevel(tt.level)
		opt.apply(&opts)
		opt = WithCompress(tt.compress)
		opt.apply(&opts)
		opt = WithConsole(tt.console)
		opt.apply(&opts)

		assert.Equal(t, opts.logPath, tt.path)
		assert.Equal(t, opts.logName, tt.name)
		assert.Equal(t, opts.logLevel, tt.level)
		assert.Equal(t, opts.maxLogSize, tt.max)
		assert.Equal(t, opts.compress, tt.compress)
		assert.Equal(t, opts.console, tt.console)
	}
}
