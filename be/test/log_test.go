package test

import (
	"github.com/kataras/golog"
	"neko-question-box-be/app/utils"
	"testing"
)

func TestLogf(t *testing.T) {
	golog.Logf(golog.DebugLevel, "this is a log text")
	utils.Logf(golog.DebugLevel, "this is a log text.")
}
