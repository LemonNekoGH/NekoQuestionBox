package main

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/context"
	"neko-question-box-be/app"
	"neko-question-box-be/app/routes"
	"runtime"
)

func main() {
	_, file, line, _ := runtime.Caller(0)

	golog.SetPrefix(fmt.Sprintf("[%s: %d] ", file, line))

	qboxApp := app.NewApp()
	qboxApp.ReadCmdArgs()

	qboxApp.Get("/", routes.RootPath)
	qboxApp.Get("/captcha", routes.Captcha)
	qboxApp.Get("/captcha-image", routes.CaptchaImage)

	qboxApp.Post("/question", routes.SubmitQuestion)
	qboxApp.Get("/question", routes.GetQuestion)

	qboxApp.Iris.Options("/{any:path}", func(context *context.Context) {
		context.StatusCode(200)
	})
	qboxApp.Get("/bing-wallpaper", routes.BingWallpaper)

	qboxApp.Start()
}
