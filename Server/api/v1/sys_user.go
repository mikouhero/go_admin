package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/Server/global/response"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	"go_admin/Server/service"
	"go_admin/Server/utils"
)

func Register(c *gin.Context) {
	var R request.RegisterStruct

	_ = c.ShouldBindJSON(&R)

	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"Nickname":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)

	if UserVerifyErr != nil {
		response.FailWithMsg(UserVerifyErr.Error(), c)
		return
	}
	user := model.SysUser{
		Username:    R.Username,
		Password:    R.Password,
		NickName:    R.Nickname,
		AuthorityId: R.AuthorityId,
	}
	err, _ := service.Register(user)
	if err != nil {
		response.FailWithDetail(response.ERROR, "", fmt.Sprintf("%v", err), c)
	} else {
		response.Ok(c)
	}

}
