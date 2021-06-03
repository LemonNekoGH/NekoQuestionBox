package routes

import (
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
		ctx.StatusCode(http.StatusInternalServerError)
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

	_, _ = ctx.JSON(question)
}
