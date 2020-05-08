package monitor

import (
	"clients/handlers"
	"clients/utils"
	"github.com/gin-gonic/gin"
)

func Active(c *gin.Context) {

	data := utils.SystemMonitor()

	handlers.SendResponse(c, nil, data)
}
