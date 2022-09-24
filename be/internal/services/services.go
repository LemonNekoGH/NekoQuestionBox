package services

import (
	"errors"
	"neko-question-box-be/internal/database"
	"neko-question-box-be/internal/database/types"
	"strings"
	"time"
)

var (
	ErrQuestionAnswered = errors.New("问题已经被回答过了")
	ErrQuestionExists   = errors.New("问题已存在")
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
	err := database.DB.Create(&types.Question{
		Question: question,
	}).Error
	// 已存在
	if err != nil && strings.Contains(err.Error(), "23505") {
		return ErrQuestionExists
	}
	return err
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
	// 不存在，更新回答
	t := time.Now()
	return database.DB.Model(q).Updates(&types.Question{
		Question:   question,
		Answer:     &answer,
		AnsweredAt: &t,
	}).Error
}
