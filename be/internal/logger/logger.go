package logger

import (
	"fmt"
	"runtime"

	"github.com/kataras/golog"
)

var logger *golog.Logger

func InitLogger() {
	logger = golog.New()

	_, file, line, ok := runtime.Caller(0)

	if !ok {
		logger.SetPrefix(fmt.Sprintf("(%s: %d) ", "unknown", 0))
		return
	}
	logger.SetPrefix(fmt.Sprintf("(%s: %d) ", file, line))
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}
