package monitor

import (
	"clients/handlers"
	"clients/utils"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {

	data := utils.SystemInfo()

	handlers.SendResponse(c, nil, data)
}