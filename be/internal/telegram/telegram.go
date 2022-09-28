package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"neko-question-box-be/internal/services"
)

var Bot *tgbotapi.BotAPI

func InitTG() {
	Bot, err := tgbotapi.NewBotAPI("") // token to be imported from conf file.
	if err != nil {
		log.Panic(err)
	}

	Bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120

	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.ReplyToMessage == nil {
			continue
		}
		err := services.UpdateAnswer(update.Message.Text, update.Message.ReplyToMessage.Text)
		if err != nil {
			msgText := err.Error()
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			msg.ReplyToMessageID = update.Message.ReplyToMessage.MessageID
			Bot.Send(msg)
		}
	}
}
