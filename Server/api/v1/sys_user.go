package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_admin/Server/global/response"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	response2 "go_admin/Server/model/response"
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
		HeaderImg:   R.HeaderImg,
		AuthorityId: R.AuthorityId,
	}
	err, userReturn := service.Register(user)
	if err != nil {
		response.FailWithDetail(response.ERROR, response2.SysUserResponse{userReturn}, fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithData(response2.SysUserResponse{userReturn}, c)
	}

}

func Login(c *gin.Context) {

	var L request.RegisterAndLoginStruct
	c.ShouldBindJSON(&L)

	//fmt.Printf("%#v",L)
	UserVerify := utils.Rules{
		//"CaptchaId": {utils.NotEmpty()},
		//"Captcha":   {utils.NotEmpty()},
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(L, UserVerify)
	fmt.Println(UserVerifyErr)
	if UserVerifyErr != nil {
		response.FailWithMsg(UserVerifyErr.Error(), c)
		return
	}

	// 验证码
	u := model.SysUser{
		Username: L.Username,
		Password: L.Password,
	}
	if err, user := service.Login(&u); err != nil {
		response.FailWithMsg(fmt.Sprintf("用户名或密码错误 %#v", err.Error()), c)
	} else {
		fmt.Println(user)
	}

}

// 签发jwt
func tokenNext(c *gin.Context,u model.SysUser){

}