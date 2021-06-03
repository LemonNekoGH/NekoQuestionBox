package app

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"neko-question-box-be/app/utils"
	"os"
	"strconv"
)

type NekoQuestionBoxApp struct {
	// 端口号，默认为 80
	// --port <int>
	port int64
	// 用于区分是否为开发模式
	// --dev
	DevMode bool
	Iris    *iris.Application
	// 验证失败的次数
	CaptchaFailedTimes  int
	CaptchaSuccessTimes int
}

type NekoHandler func(app *NekoQuestionBoxApp, ctx *context.Context)

// NewApp
// 新建自己的 APP 实例
func NewApp() *NekoQuestionBoxApp {
	irisApp := iris.New()
	return &NekoQuestionBoxApp{
		Iris:                irisApp,
		CaptchaFailedTimes:  0,
		CaptchaSuccessTimes: 0,
	}
}

// NewTestApp
// 新建用于测试的 APP 实例
func NewTestApp() *NekoQuestionBoxApp {
	irisApp := iris.New()
	return &NekoQuestionBoxApp{
		Iris:                irisApp,
		DevMode:             true,
		CaptchaFailedTimes:  0,
		CaptchaSuccessTimes: 0,
	}
}

// ReadCmdArgs
// 从命令行读取参数
func (app *NekoQuestionBoxApp) ReadCmdArgs() {
	if utils.IsArrayContains(os.Args, "--dev") {
		app.DevMode = true
	}

	allowOrigin := "http://qbox.lemonneko.moe"

	if app.DevMode {
		allowOrigin = "*"
	}

	app.Iris.Use(cors.New(cors.Options{
		AllowedOrigins: []string{allowOrigin},
	}))

	var (
		portArgIndex int
		port         string
	)

	portArgIndex = utils.IndexOf(os.Args, "--port")

	if portArgIndex != -1 {
		port = os.Args[portArgIndex+1]
		portNumber, err := strconv.ParseInt(port, 10, 16)
		if err != nil {
			golog.Error(err.Error())
			return
		}
		app.port = portNumber
	} else {
		app.port = 80
	}
}

// Get
// 重新封装一层 Get 以注入自己的变量
func (app *NekoQuestionBoxApp) Get(path string, handler NekoHandler) {
	app.Iris.Get(path, func(context *context.Context) {
		handler(app, context)
	})
}

// Post
// 重新封装一层 Post 以注入自己的变量
func (app *NekoQuestionBoxApp) Post(path string, handler NekoHandler) {
	app.Iris.Post(path, func(context *context.Context) {
		handler(app, context)
	})
}

// Start
// 启动服务器
func (app *NekoQuestionBoxApp) Start() {
	app.Iris.Listen(fmt.Sprintf(":%d", app.port))
}
