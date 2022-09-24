package logger

import (
	"fmt"
	"neko-question-box-be/pkg/utils"
	"runtime"
	"strings"

	"github.com/kataras/golog"
)

var logger *golog.Logger

func InitLogger() {
	logger = golog.New()
}

func Infof(format string, args ...interface{}) {
	Logf(golog.InfoLevel, format, args...)
}

func Debugf(format string, args ...interface{}) {
	Logf(golog.DebugLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	Logf(golog.ErrorLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	Logf(golog.WarnLevel, format, args...)
}

func Logf(level golog.Level, format string, arg ...interface{}) {
	shouldSkip := true
	skipTimes := 0

	var (
		file    string
		line    int
		fun     uintptr
		funName string
	)

	for shouldSkip {
		fun, file, line, _ = runtime.Caller(skipTimes)
		funNames := []string{
			"neko-question-box-be/pkg/utils.Logf",
			"neko-question-box-be/pkg/utils.Errorf",
			"neko-question-box-be/pkg/utils.Debugf",
			"neko-question-box-be/pkg/utils.Warnf",
			"neko-question-box-be/pkg/utils.Infof",
		}
		funName = runtime.FuncForPC(fun).Name()
		if !utils.IsArrayContains(funNames, funName) {
			shouldSkip = false
		} else {
			skipTimes++
		}
	}

	fileParts := strings.Split(file, "/")

	file = fileParts[len(fileParts)-1]

	logger.SetPrefix(fmt.Sprintf("[%s:%s:%d] ", file, funName, line))
	logger.Logf(level, format, arg...)
}
