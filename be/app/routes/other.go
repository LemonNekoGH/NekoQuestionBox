package routes

import (
	"github.com/kataras/iris/v12/context"
	"io"
	"neko-question-box-be/app"
	"neko-question-box-be/app/utils"
	"net/http"
)

// BingWallpaper
// 获取必应每日壁纸
func BingWallpaper(app *app.NekoQuestionBoxApp, ctx *context.Context) {
	utils.ShouldAcceptJson(ctx)
	resp, err := http.Get("https://www.bing.com/HPImageArchive.aspx?format=js&idx=2&n=1")
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		utils.Errorf(err.Error())
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			ctx.StatusCode(http.StatusInternalServerError)
			utils.Errorf(err.Error())
		}
	}(resp.Body)

	p, err := io.ReadAll(resp.Body)

	_, err = ctx.Write(p)
	if err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		utils.Errorf(err.Error())
		return
	}
}
