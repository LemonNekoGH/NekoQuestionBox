package telegram

import (
	"neko-question-box-be/internal/config"
	"neko-question-box-be/internal/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func InitTG() {
	var err error
	Bot, err = tgbotapi.NewBotAPI(config.Conf.Telegram.ApiToken) // token to be imported from conf file.
	if err != nil {
		panic(err)
	}
	// 发送一条消息至目标聊天以验证配置正确
	_, err = Bot.Send(tgbotapi.NewMessage(config.Conf.Telegram.ChatID, "NekoQuestionBox Bot Started"))
	if err != nil {
		panic(err)
	}
	logger.Infof("TG Bot 启动成功")
}
