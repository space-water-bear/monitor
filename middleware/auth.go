package middleware

import (
	"clients/handlers"
	"clients/pkg/errno"
	"clients/pkg/token"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 校验中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		context, err := token.ParseRequest(c)
		if err != nil {
			handlers.SendResponse(c, errno.ErrToken, nil)
			c.Abort()
			return
		}
		//fmt.Println(context)
		if context.Keys != "game_ops" {
			c.Abort()
			return
		}
		c.Next()
	}
}
