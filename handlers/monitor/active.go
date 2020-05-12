package monitor

import (
	"clients/handlers"
	"clients/utils"
	"github.com/gin-gonic/gin"
)

func Active(c *gin.Context) {

	err := utils.SendMonitor()
	if err != nil {
		handlers.SendResponse(c, err, nil)
	}

	handlers.SendResponse(c, nil, nil)
}
