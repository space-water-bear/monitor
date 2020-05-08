package work

import (
	"clients/handlers"
	"clients/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context)  {

	var r createRequest

	if err := c.BindJSON(&r); err != nil {
		//fmt.Println(err)
		handlers.SendResponse(c, errno.ErrBind, nil)
		return
	}

	fmt.Println(r.ID)
	fmt.Println(r.Data)

	handlers.SendResponse(c, nil, nil)
}