package test

import (
	"github.com/kataras/iris/v12/httptest"
	nekoApp "neko-question-box-be/app"
	"neko-question-box-be/app/routes"
	"net/http"
	"testing"
)

func TestSaveQuestion(t *testing.T) {
	app := nekoApp.NewApp()
	app.Post("/save-question", routes.SaveQuestion)

	e := httptest.New(t, app.Iris)
	e.POST("/save-question").WithHeader("Accept", "application/json").WithText("miao").Expect().Status(http.StatusOK)
}
