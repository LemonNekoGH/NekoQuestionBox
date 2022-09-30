package main

import (
	"fmt"
	"neko-question-box-be/internal/config"
	"neko-question-box-be/internal/database"
	"neko-question-box-be/internal/logger"
	"neko-question-box-be/internal/server"
	"neko-question-box-be/internal/services"
	"neko-question-box-be/internal/telegram"
)

func main() {
	logger.InitLogger()
	config.InitConfig(false)
	database.InitDB()

	if config.Conf.Telegram.Enabled {
		telegram.InitTG()
		go services.ReceiveQuestionAnswer()
	} else {
		logger.Infof("配置中没有启用 TG Bot")
	}

	r := server.InitServer()
	r.Run(fmt.Sprintf("0.0.0.0:%d", config.Conf.Port))
}
