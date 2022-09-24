package api

import (
	"fmt"
	"neko-question-box-be/internal/config"
	"neko-question-box-be/internal/database"
	"neko-question-box-be/internal/database/types"
	"neko-question-box-be/internal/logger"
	"neko-question-box-be/pkg/handler"
	"net/http"
	"testing"

	"github.com/dchest/captcha"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	logger.InitLogger()
	config.InitConfig(true)
	database.InitDBTest()

	m.Run()
}

func TestGetPing(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)
		_, c := handler.CreateTestContext(http.MethodGet, "/ping", nil)
		resp, err := getPing(c)
		r.Empty(resp)
		r.Empty(err)
	})
}

func TestGetBing(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)
		_, c := handler.CreateTestContext(http.MethodGet, "/bing-wallpaper", nil)
		resp, err := getBingWallpaper(c)
		r.NotEmpty(resp)
		r.Empty(err)
	})
}

func TestGetQuestion(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)
		// 存入问题
		q := types.Question{
			Question: "test",
		}
		err := database.DB.Create(&q).Error
		r.Empty(err)

		// 收尾，移除数据
		defer func() {
			err := database.DB.Model(&q).Where("question = ?", q.Question).Delete(q).Error
			r.Empty(err)
		}()

		_, c := handler.CreateTestContext(http.MethodGet, "/question", nil)
		resp, err := getQuestion(c)
		respB, ok := resp.([]types.Question)

		r.True(ok)
		r.Len(respB, 1)
		r.Equal(q.Question, respB[0].Question)
		r.NotEmpty(respB[0].CreatedAt)
		r.NotEmpty(respB[0].ID)
		r.Empty(respB[0].Answer)
		r.Empty(respB[0].AnsweredAt)
		r.Empty(err)
	})

	t.Run("empty", func(t *testing.T) {
		r := require.New(t)
		_, c := handler.CreateTestContext(http.MethodGet, "/question", nil)
		resp, err := getQuestion(c)
		r.Empty(resp)
		r.Empty(err)
	})
}

func TestPostQuestion(t *testing.T) {
	t.Run("ErrParams", func(t *testing.T) {
		r := require.New(t)
		_, c := handler.CreateTestContext(http.MethodPost, "/question", &postQuestionReq{})
		resp, err := postQuestion(c)
		r.Empty(resp)
		r.Equal(handler.ErrParams, err)
	})

	t.Run("ErrCaptcha", func(t *testing.T) {
		r := require.New(t)
		_, c := handler.CreateTestContext(http.MethodPost, "/question", &postQuestionReq{
			Id:       "test",
			Value:    "test2",
			Question: "q",
		})
		resp, err := postQuestion(c)
		r.Empty(resp)
		r.Equal(handler.ErrCaptcha, err)
	})
}

func TestGetCpatchaImage(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)
		// 新建 Captcha
		id := captcha.New()

		_, c := handler.CreateTestContext(http.MethodGet, fmt.Sprintf("/captcha-image?id=%s", id), nil)
		resp, err := getCaptchaImage(c)
		r.NotEmpty(resp)
		r.Empty(err)
	})

	t.Run("captcha not exists", func(t *testing.T) {
		r := require.New(t)
		_, c := handler.CreateTestContext(http.MethodGet, "/captcha-image?id=123456", nil)
		resp, err := getCaptchaImage(c)

		convertErr, ok := err.(handler.HandlerError)
		r.True(ok)
		r.Empty(resp)
		r.Equal(http.StatusInternalServerError, convertErr.Status)
	})

	t.Run("ErrParams", func(t *testing.T) {
		r := require.New(t)
		_, c := handler.CreateTestContext(http.MethodGet, "/captcha-image", nil)
		resp, err := getCaptchaImage(c)
		r.Empty(resp)
		r.Equal(handler.ErrParams, err)
	})
}

func TestGetCaptcha(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)

		_, c := handler.CreateTestContext(http.MethodGet, "/captcha", nil)
		resp, err := getCaptcha(c)
		r.NotEmpty(resp)
		r.Empty(err)
	})
}

func TestHandlers(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)

		h := Handlers()
		// 提取出路径
		paths := []string{}

		for path := range h.Group {
			paths = append(paths, path)
		}
		r.ElementsMatch(paths, []string{
			"captcha",
			"captcha-image",
			"bing",
			"question",
			"ping",
		})
	})
}
