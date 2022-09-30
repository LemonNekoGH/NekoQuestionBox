package services

import (
	"errors"
	"neko-question-box-be/internal/config"
	"neko-question-box-be/internal/database"
	"neko-question-box-be/internal/database/types"
	"neko-question-box-be/internal/logger"
	"neko-question-box-be/internal/telegram"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

// 把问题通过 Telegram 发送给被提问的人
func SendToTgChat(question string) error {
	msg := tgbotapi.NewMessage(config.Conf.Telegram.ChatID, question) // ChatID will be imported from a configuration file.
	_, err := telegram.Bot.Send(msg)
	if err != nil {
		logger.Errorf("send message error: ", err.Error())
		return err
	}
	return nil
}

// 开始接收问题的答案
func ReceiveQuestionAnswer() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120

	updates := telegram.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.ReplyToMessage == nil {
			continue
		}
		err := UpdateAnswer(update.Message.Text, update.Message.ReplyToMessage.Text)
		// 如果出错，而且允许发送错误，将发送错误信息到目标聊天 ID
		if err != nil && config.Conf.Telegram.SendErrors {
			msgText := err.Error()
			if database.IsNoRecordFoundError(err) {
				msgText = "问题不存在"
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			msg.ReplyToMessageID = update.Message.ReplyToMessage.MessageID
			telegram.Bot.Send(msg)
		}
	}
}
