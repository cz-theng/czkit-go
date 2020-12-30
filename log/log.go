package log

import (
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type logcat struct {
	zsl  *zap.SugaredLogger
	opts options
}

var _logcat = logcat{
	zsl:  zap.NewExample().Sugar(),
	opts: defaultOptions(),
}

// Init create and init a zap SugarLogger
func Init(opts ...Option) error {

	for _, opt := range opts {
		opt.apply(&_logcat.opts)
	}

	fn := path.Join(_logcat.opts.logPath, _logcat.opts.logName)

	fileSync := lumberjack.Logger{
		Filename: fn,
		MaxSize:  int(_logcat.opts.maxLogSize),
		Compress: _logcat.opts.compress,
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	var ws zapcore.WriteSyncer
	if _logcat.opts.console {
		ws = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&fileSync),
		)
	} else {
		ws = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(&fileSync),
		)
	}
	core := zapcore.NewCore(encoder, ws, zap.NewAtomicLevelAt(_logcat.opts.logLevel))

	// TODO: add caller/tracer/callerskip
	_logcat.zsl = zap.New(core).Sugar()
	return nil

}

// Sync sync all log to disk
func Sync() {
	_logcat.zsl.Sync()
}

// Debug print debug log
func Debug(template string, args ...interface{}) {
	_logcat.zsl.Debugf(template, args...)
}

// Info  print info log
func Info(template string, args ...interface{}) {
	_logcat.zsl.Infof(template, args...)
}

// Warning  print warning log
func Warning(template string, args ...interface{}) {
	_logcat.zsl.Warnf(template, args...)
}

// Error  print error log
func Error(template string, args ...interface{}) {
	_logcat.zsl.Errorf(template, args...)
}

// Fatal  print fatal log
func Fatal(template string, args ...interface{}) {
	_logcat.zsl.Fatalf(template, args...)
}

// Debugw print debug log
func Debugw(msg string, keysAndValues ...interface{}) {
	_logcat.zsl.Debugw(msg, keysAndValues...)
}

// Infow  print info log
func Infow(msg string, keysAndValues ...interface{}) {
	_logcat.zsl.Infow(msg, keysAndValues...)
}

// Warningw  print warning log
func Warningw(msg string, keysAndValues ...interface{}) {
	_logcat.zsl.Warnw(msg, keysAndValues...)
}

// Errorw  print error log
func Errorw(msg string, keysAndValues ...interface{}) {
	_logcat.zsl.Errorw(msg, keysAndValues...)
}

// Fatalw  print fatal log
func Fatalw(msg string, keysAndValues ...interface{}) {
	_logcat.zsl.Fatalw(msg, keysAndValues...)
}
