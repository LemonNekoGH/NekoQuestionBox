package services

import (
	"errors"
	"neko-question-box-be/internal/database"
	"neko-question-box-be/internal/database/types"
)

var (
	ErrQuestionAnswered = errors.New("问题已经被回答过了")
)

// 获取所有的问题
func GetAllQuestions() ([]types.Question, error) {
	questions := make([]types.Question, 0)
	err := database.DB.Find(&questions).Error
	if err != nil && !database.IsNoRecordFoundError(err) {
		return []types.Question{}, err
	}
	return questions, err
}

// 添加新问题
func CreateNewQuestion(question string) error {
	return database.DB.Save(&types.Question{
		Question: question,
	}).Error
}

// 更新回答
func UpdateAnswer(answer, question string) error {
	q := &types.Question{}
	err := database.DB.Model(q).Where("question = ?", question).First(q).Error
	if err != nil {
		return err
	}
	// 回答已经存在
	if q.Answer != nil {
		return ErrQuestionAnswered
	}
	// 一切正常
	return nil
}
