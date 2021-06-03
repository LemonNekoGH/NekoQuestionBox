package test

import (
	"github.com/kataras/iris/v12/httptest"
	nekoApp "neko-question-box-be/app"
	"neko-question-box-be/app/routes"
	"testing"
)

func TestRootPath(t *testing.T) {
	app := nekoApp.NewTestApp()
	app.Get("/", routes.RootPath)

	e := httptest.New(t, app.Iris)
	e.GET("/").Expect().Status(httptest.StatusForbidden)

	e.GET("/").WithHeader("Accept", "application/json").Expect().Status(100)
}
