package services

import (
	"neko-question-box-be/internal/config"
	"neko-question-box-be/internal/database"
	"neko-question-box-be/internal/database/types"
	"neko-question-box-be/internal/logger"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	logger.InitLogger()
	config.InitConfig(true)
	database.InitDBTest()

	m.Run()
}

func TestGetQuestions(t *testing.T) {
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

		respB, err := GetAllQuestions()

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
		resp, err := GetAllQuestions()
		r.Empty(resp)
		r.Empty(err)
	})
}

func TestCreateNewQuestion(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		r := require.New(t)
		// 存入问题
		err := CreateNewQuestion("test_question")
		r.Empty(err)

		// 收尾，移除数据
		defer func() {
			err := database.DB.Model(&types.Question{}).Where("question = ?", "test_question").Delete(&types.Question{}).Error
			r.Empty(err)
		}()

		// 读取问题
		questions := make([]types.Question, 0)
		err = database.DB.Find(&questions).Error
		r.Empty(err)
	})

	t.Run("exists", func(t *testing.T) {
		r := require.New(t)
		// 存入问题
		err := CreateNewQuestion("test_question")
		r.Empty(err)
		// 再次存入
		err = CreateNewQuestion("test_question")
		r.ErrorIs(err, ErrQuestionExists)

		// 收尾，移除数据
		defer func() {
			err := database.DB.Model(&types.Question{}).Where("question = ?", "test_question").Delete(&types.Question{}).Error
			r.Empty(err)
		}()
	})
}

func TestAnswerQuestion(t *testing.T) {
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

		// 回答问题
		err = UpdateAnswer("answer", "test")
		r.Empty(err)
	})

	t.Run("answered", func(t *testing.T) {
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

		// 回答问题
		err = UpdateAnswer("answer", "test")
		r.Empty(err)
		// 再次回答
		err = UpdateAnswer("answer", "test")
		r.ErrorIs(err, ErrQuestionAnswered)
	})

	t.Run("not exists", func(t *testing.T) {
		r := require.New(t)

		// 回答问题
		err := UpdateAnswer("answer", "test")
		r.ErrorIs(err, gorm.ErrRecordNotFound)
	})
}
