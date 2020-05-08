package cron

import (
	"clients/handlers"
	"clients/pkg/cron"
	"github.com/gin-gonic/gin"
)

func Reload(c *gin.Context)  {

	cron.Close()

	go func() {
		cron.Init()
		defer cron.Close()
	}()
	handlers.SendResponse(c, nil, nil)
}
