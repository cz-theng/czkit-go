package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zap.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zap.InfoLevel
	// WarningLevel logs are more important than Info, but don't need individual
	// human review.
	WarningLevel = zap.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zap.ErrorLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zap.FatalLevel
)

type options struct {
	logPath    string
	logName    string
	logLevel   zapcore.Level
	maxLogSize uint32
	compress   bool
	console    bool
}

func defaultOptions() options {
	return options{
		logPath:    "./",
		logName:    "log.log",
		maxLogSize: 100 * 1024 * 1024,
		console:    true,
		compress:   false,
		logLevel:   DebugLevel,
	}
}

// Option configures how we set up the zap logger.
type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// WithLogPath determines which path to create log files
func WithLogPath(path string) Option {
	return newFuncOption(func(o *options) {
		o.logPath = path
	})
}

// WithLogName determines which name to create log files
func WithLogName(name string) Option {
	return newFuncOption(func(o *options) {
		o.logName = name
	})
}

// WithLogLevel determines which level to log
func WithLogLevel(level zapcore.Level) Option {
	return newFuncOption(func(o *options) {
		o.logLevel = level
	})
}

// WithMaxLogSize determines max file size of a log file
func WithMaxLogSize(size uint32) Option {
	return newFuncOption(func(o *options) {
		o.maxLogSize = size
	})
}

// WithCompress determines if compress rolled file
func WithCompress(compress bool) Option {
	return newFuncOption(func(o *options) {
		o.compress = compress
	})
}

// WithConsole determines weather output to console
func WithConsole(console bool) Option {
	return newFuncOption(func(o *options) {
		o.console = console
	})
}
