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

	submitData := routes.SubmitData{
		Question: "miao",
	}
	e.POST("/save-question").WithHeader("Accept", "application/json").WithJSON(submitData).Expect().Status(http.StatusOK)
}

func TestGetQuestion(t *testing.T) {
	app := nekoApp.NewApp()
	app.Get("/question", routes.GetQuestion)

	e := httptest.New(t, app.Iris)
	e.GET("/question").WithHeader("Accept", "application/json").Expect().Status(httptest.StatusOK)
}
