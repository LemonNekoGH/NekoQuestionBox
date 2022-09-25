package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

var (
	ErrParams         = NewHandlerError(http.StatusBadRequest, 40001, "params error")
	ErrCaptcha        = NewHandlerError(http.StatusNotAcceptable, 40601, "captcha error")
	ErrQuestionExists = NewHandlerError(http.StatusConflict, 40901, "question exists")
)

// 创建测试用上下文
func CreateTestContext(method string, path string, body any) (*httptest.ResponseRecorder, *gin.Context) {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	// 把请求参数写入上下文中
	switch t := body.(type) {
	case io.Reader:
		c.Request, _ = http.NewRequest(method, path, t)
	case string:
		// 是查询参数
		params, err := url.ParseQuery(t)
		if err != nil {
			panic(err)
		}
		c.Request, _ = http.NewRequest(method, path, nil)
		// 告诉上下文这个请求的格式
		c.Request.Header.Set("content-type", "application/x-www-form-urlencoded")
		// 设置查询参数
		c.Request.PostForm = params
	default:
		// 其它情况，把请求体转换成 JSON
		bd, _ := json.Marshal(body)
		c.Request, _ = http.NewRequest(method, path, bytes.NewBuffer(bd))
		// 告诉上下文这个请求体的格式
		c.Request.Header.Set("content-type", "application/json")
	}

	return recorder, c
}

// 控制器要返回的错误
type HandlerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
}

// 控制器的正常返回
type HandlerResponse interface{}

// 实现错误接口
func (h HandlerError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", h.Code, h.Message)
}

// 新建一个控制器错误
func NewHandlerError(status int, code int, message string) HandlerError {
	return HandlerError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}

type Handler struct {
	Func   gin.HandlerFunc // 实际函数
	Mehtod string          // 方法
}

type HandlerFunc func(ctx *gin.Context) (HandlerResponse, error)

// 新建控制器
func NewHandler(method string, fun HandlerFunc) Handler {
	return Handler{
		Mehtod: method,
		Func: func(ctx *gin.Context) {
			resp, err := fun(ctx)
			// 错误处理
			if err != nil {
				switch e := err.(type) {
				case HandlerError:
					ctx.AbortWithStatusJSON(e.Status, e)
					return
				// 其它错误，返回 500
				default:
					ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"code":    50001,
						"message": e.Error(),
					})
					return
				}
			}
			// 正常
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": resp,
			})
		},
	}
}

type HandlerGroup struct {
	Name        string               // 控制器组名
	Group       map[string][]Handler // 控制器
	SubHandlers []HandlerGroup       // 子控制器组
}

// 挂载控制器组
func (g HandlerGroup) Install(parent *gin.RouterGroup) {
	group := parent.Group(g.Name)

	// 挂载控制器
	for name, handlers := range g.Group {
		for _, handler := range handlers {
			group.Handle(handler.Mehtod, name, handler.Func)
		}
	}

	// 挂载子控制器组
	for _, subGroup := range g.SubHandlers {
		subGroup.Install(group)
	}
}
