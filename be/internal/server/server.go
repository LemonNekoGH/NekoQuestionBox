package server

import (
	"neko-question-box-be/internal/api"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	r := gin.Default()
	api.Handlers().Install(r.Group(""))

	return r
}
