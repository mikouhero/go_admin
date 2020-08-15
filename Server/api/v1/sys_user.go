package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/Server/model/request"
	"go_admin/Server/utils"
)

func Register(c *gin.Context) {
	var R request.RegisterStruct
	c.ShouldBindJSON(&R)

	UserVerify := utils.Rules{
		"Username": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	fmt.Println(UserVerifyErr)
}
