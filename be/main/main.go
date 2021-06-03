package main

import (
	"neko-question-box-be/app"
	"neko-question-box-be/app/routes"
)

func main() {
	qboxApp := app.NewApp()
	qboxApp.ReadCmdArgs()

	qboxApp.Get("/", routes.RootPath)
	qboxApp.Get("/captcha", routes.Captcha)
	qboxApp.Get("/captcha-image", routes.CaptchaImage)

	qboxApp.Start()
}
