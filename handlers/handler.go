package handlers

import (
	"github.com/gin-gonic/gin"
	"clients/pkg/errno"
	"net/http"
)

type Response struct {	// 包装返回结构
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
