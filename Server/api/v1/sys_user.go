package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go_admin/Server/global"
	"go_admin/Server/global/response"
	"go_admin/Server/middleware"
	"go_admin/Server/model"
	"go_admin/Server/model/request"
	response2 "go_admin/Server/model/response"
	"go_admin/Server/service"
	"go_admin/Server/utils"
	"time"
)

func SetUserAuthority(c *gin.Context) {
	var setUsetAuth request.SetUserAuth
	_ = c.ShouldBindJSON(&setUsetAuth)
	UserVerify := utils.Rules{
		"UUID":        {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	err := utils.Verify(setUsetAuth, UserVerify)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
		return
	}

	err = service.SetUserAuthority(setUsetAuth.UUID, setUsetAuth.AuthorityId)
	if err != nil {
		response.FailWithMsg("更新失败"+err.Error(), c)
	} else {
		response.OkWithMsg("ok", c)
	}

}

func DeleteUser(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])

	if IdVerifyErr != nil {
		response.FailWithMsg(IdVerifyErr.Error(), c)
		return
	}
	err := service.DeleteUser(reqId.Id)
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.FailWithMsg("删除成功", c)
	}
}

func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	pageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if pageVerifyErr != nil {
		response.FailWithMsg(pageVerifyErr.Error(), c)
		return
	}

	err, list, totle := service.GetUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMsg(err.Error(), c)
	} else {
		response.OkWithData(response2.PageResult{
			List:     list,
			Totle:    totle,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}

}

func UploadHeaderImg(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse, _ := claims.(*request.CustomClaims)
	uuid := waitUse.UUID
	_, header, err := c.Request.FormFile("headerImg")
	if err != nil {
		response.FailWithMsg("上传文件失败，"+err.Error(), c)
	} else {
		err, filePath, _ := utils.Upload(header)
		if err != nil {
			response.FailWithMsg("上传文件失败"+err.Error(), c)
		} else {

			err, inter := service.UploadHeaderImg(uuid, filePath)
			if err != nil {
				response.FailWithMsg("图片更新失败"+err.Error(), c)
			} else {
				response.OkWithData(response2.SysUserResponse{User: *inter}, c)
			}
		}

	}
}

func ChangePassword(c *gin.Context) {
	var params request.ChangePasswordStruct
	_ = c.ShouldBindJSON(&params)
	UserVerity := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"NewPassword": {utils.NotEmpty()},
	}
	UserverityErr := utils.Verify(params, UserVerity)
	if UserverityErr != nil {
		response.FailWithMsg(UserverityErr.Error(), c)
		return
	}
	user := &model.SysUser{
		Username: params.Username,
		Password: params.Password,
	}
	if err, _ := service.ChangePassword(user, params.NewPassword); err != nil {
		response.FailWithMsg("修改失败，检查用户名或密码"+err.Error(), c)
	} else {
		response.OkWithMsg("修改成功", c)
	}

}

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
	_ = c.ShouldBindJSON(&L)

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
		tokenNext(c, *user)
	}

}

// 签发jwt
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{[]byte(global.GVA_CONFIG.JWT.SigningKey)}
	clams := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,    // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, //过期时间
			Issuer:    "amdin",                     // 签名的发行者
		},
	}

	token, err := j.CreateToken(clams)

	if err != nil {
		response.FailWithMsg("获取token 失败"+err.Error(), c)
	}

	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithData(response2.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt,
		}, c)
		return
	}

	var loginJwt model.JwtBlacklist
	loginJwt.Jwt = token

	err, jwtStr := service.GetRedisJWT(user.Username)

	err = nil

	if err == redis.Nil {
		if err := service.SetRedisJWT(loginJwt, user.Username); err != nil {
			response.FailWithMsg("设置登录状态失败"+err.Error(), c)
			return
		}
		response.OkWithData(response2.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt,
		}, c)
	} else if err != nil {
		response.FailWithMsg(fmt.Sprintf("%v", err), c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		fmt.Println("--------")
		fmt.Printf("%#v", blackJWT)
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMsg("jwt作废失败", c)
			return
		}

		if err := service.SetRedisJWT(loginJwt, user.Username); err != nil {
			response.FailWithMsg("设置登录状态失败", c)
			return
		}
		response.OkWithData(response2.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt,
		}, c)

	}

}
