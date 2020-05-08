package cron

import (
	"clients/handlers"
	"clients/pkg/cron"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context)  {

	err := cron.AddSystemInfoJob()
	if err != nil {
		handlers.SendResponse(c, err, nil)
		return
	}

	handlers.SendResponse(c, nil, nil)
}
