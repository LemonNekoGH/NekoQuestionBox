package routes

import (
	"github.com/dchest/captcha"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/context"
	"neko-question-box-be/app"
	"neko-question-box-be/app/utils"
	"net/http"
)

type CaptchaData struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

// Captcha
// GET 获取新的 Captcha ID
// POST 进行验证
func Captcha(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	if app.DevMode {
		ctx.Header("Access-Control-Allow-Origin", "*")
	} else {
		ctx.Header("Access-Control-Allow-Origin", "qbox.lemonneko.moe")
	}
	if !utils.ShouldAcceptJson(ctx) {
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(struct {
		Id string `json:"id"`
	}{
		Id: captcha.New(),
	})
}

func CaptchaImage(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	id := ctx.URLParamDefault("id", "null")
	if id == "null" {
		ctx.StatusCode(http.StatusBadRequest)
	}
	if app.DevMode {
		ctx.Header("Access-Control-Allow-Origin", "*")
	} else {
		ctx.Header("Access-Control-Allow-Origin", "qbox.lemonneko.moe")
	}
	ctx.Header("Content-Type", "image/png")
	ctx.Header("Cache-Control", "no-store")
	err := captcha.WriteImage(ctx.ResponseWriter(), id, 200, 100)
	if err != nil {
		golog.Error(err.Error())
		ctx.StatusCode(http.StatusNotFound)
	}
}
