package monitor

import (
	"clients/handlers"
	"clients/utils"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {

	err := utils.SendInfo()
	if err != nil {
		handlers.SendResponse(c, err, nil)
	}

	handlers.SendResponse(c, nil, nil)
}
