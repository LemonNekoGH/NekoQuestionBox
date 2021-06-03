package routes

import (
	"github.com/kataras/iris/v12/context"
	"neko-question-box-be/app"
)

// RootPath
// 用于给前端检测服务器状态
func RootPath(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	if app.DevMode {
		ctx.Header("Access-Control-Allow-Origin", "*")
	} else {
		ctx.Header("Access-Control-Allow-Origin", "qbox.lemonneko.moe")
	}

	accept := ctx.GetHeader("Accept")
	if accept == "application/json" {
		ctx.StatusCode(200)
	} else {
		// 需要先发送响应码再进行其它操作
		ctx.StatusCode(403)
		ctx.Application().Logger().Warnf("%s %s", ctx.Path(), "有人来了不该来的地方")
		_, err := ctx.WriteString("You should not be here.")
		if err != nil {
			return
		}
	}
}
