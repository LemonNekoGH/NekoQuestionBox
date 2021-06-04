package routes

import (
	"database/sql"
	"github.com/dchest/captcha"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/context"
	"neko-question-box-be/app"
	"neko-question-box-be/app/utils"
	"net/http"
)

type SubmitData struct {
	Id       string `json:"captchaId"`
	Value    string `json:"captchaValue"`
	Question string `json:"question"`
}

func SubmitQuestion(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	if !utils.ShouldAcceptJson(ctx) {
		return
	}
	var question = &SubmitData{}
	err := ctx.ReadJSON(question)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		golog.Error(err.Error())
		return
	}

	success := captcha.VerifyString(question.Id, question.Value)
	if !success {
		ctx.StatusCode(http.StatusNotAcceptable)
		app.CaptchaFailedTimes++
		golog.Warnf("验证失败，已累计 %d 次", app.CaptchaFailedTimes)
		return
	}
	app.CaptchaSuccessTimes++
	golog.Warnf("验证成功，已累计 %d 次", app.CaptchaSuccessTimes)

	ctx.Values().Set("question", question.Question)

	SaveQuestion(app, ctx)
}

func SaveQuestion(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	question := ""
	if ctx.Path() == "/save-question" {
		err := ctx.ReadBody(question)
		if err != nil {
			golog.Error(err.Error())
			return
		}
	}
	question = ctx.Values().GetString("question")

	db, err := sql.Open("mysql", "root:LemonNeko@tcp(localhost:3306)")
	if err != nil {
		golog.Error(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("")
	if err != nil {
		golog.Error(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}

	err = db.Close()
	if err != nil {
		golog.Error(err.Error())
		ctx.StatusCode(http.StatusInternalServerError)
		return
	}
}
