package main

import (
	"neko-question-box-be/internal/config"
	"neko-question-box-be/internal/database"
	"neko-question-box-be/internal/logger"
	"neko-question-box-be/internal/server"
)

func main() {
	logger.InitLogger()
	config.InitConfig(false)
	database.InitDB()

	r := server.InitServer()
	r.Run()
}
