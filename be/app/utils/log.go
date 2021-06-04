package utils

import (
	"fmt"
	"github.com/kataras/golog"
	"runtime"
	"strings"
)

func Infof(format string, arg ...interface{}) {
	Logf(golog.InfoLevel, format, arg...)
}

func Errorf(format string, arg ...interface{}) {
	Logf(golog.ErrorLevel, format, arg...)
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
			"neko-question-box-be/app/utils.Logf",
			"neko-question-box-be/app/utils.Errorf",
			"neko-question-box-be/app/utils.Warnf",
			"neko-question-box-be/app/utils.Infof",
		}
		funName = runtime.FuncForPC(fun).Name()
		if !IsArrayContains(funNames, funName) {
			shouldSkip = false
		} else {
			skipTimes++
		}
	}

	fileParts := strings.Split(file, "/")

	file = fileParts[len(fileParts)-1]

	golog.SetPrefix(fmt.Sprintf("[%s:%s:%d] ", funName, file, line))
	golog.Logf(level, format, arg...)
}
