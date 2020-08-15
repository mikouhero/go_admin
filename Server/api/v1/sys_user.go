package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/Server/model/request"
)

func Register(c *gin.Context) {
	var R request.RegisterStruct
	c.ShouldBindJSON(&R)
	// 数据校验
}
