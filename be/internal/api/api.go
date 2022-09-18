package api

import (
	"neko-question-box-be/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 验证 Captcha
func postCaptcha(ctx *gin.Context) (handler.HandlerResponse, error) {
	return nil, nil
}

// 获取新的 captcha 图片
func getCaptcha(ctx *gin.Context) (handler.HandlerResponse, error) {
	return nil, nil
}

// 获取 bing 每日壁纸
func getBingWallpaper(ctx *gin.Context) (handler.HandlerResponse, error) {
	return nil, nil
}

// 获取已经有的问题和答案
func getQuestion(ctx *gin.Context) (handler.HandlerResponse, error) {
	return nil, nil
}

// 提交新的问题
func postQuestion(ctx *gin.Context) (handler.HandlerResponse, error) {
	return nil, nil
}

// 检查服务器状态
func getPing(ctx *gin.Context) (handler.HandlerResponse, error) {
	return nil, nil
}

func Handlers() handler.HandlerGroup {
	return handler.HandlerGroup{
		Name: "",
		Group: map[string][]handler.Handler{
			"captcha": {
				handler.NewHandler(http.MethodPost, postCaptcha),
				handler.NewHandler(http.MethodGet, getCaptcha),
			},
			"bing": {handler.NewHandler(http.MethodGet, getBingWallpaper)},
			"question": {
				handler.NewHandler(http.MethodGet, getQuestion),
				handler.NewHandler(http.MethodPost, postQuestion),
			},
			"ping": {
				handler.NewHandler(http.MethodGet, getPing),
			},
		},
	}
}
