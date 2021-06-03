package test

import (
	"github.com/kataras/iris/v12/httptest"
	nekoApp "neko-question-box-be/app"
	"neko-question-box-be/app/routes"
	"testing"
)

func TestCaptcha(t *testing.T) {
	app := nekoApp.NewTestApp()
	app.Get("/captcha", routes.Captcha)
	app.Post("/captcha", routes.Captcha)

	e := httptest.New(t, app.Iris)
	println(e.GET("/captcha").WithHeader("Accept", "application/json").Expect().Body().Raw())
}
