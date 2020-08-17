package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_admin/Server/global"
	"go_admin/Server/global/response"
	"go_admin/Server/model"
	"go_admin/Server/service"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录

		token := c.Request.Header.Get("x-token")

		if token == "" {
			response.Result(400, gin.H{"reload": true}, "未登录或者非法访问", c)
			c.Abort()
			return
		}
		modelToken := model.JwtBlacklist{
			Jwt: token,
		}
		if service.IsBlackList(token, modelToken) {
			response.Result(400, gin.H{"reload": true}, "token失效", c)
			c.Abort()
			return
		}

	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}
}

// 创建token

// 删除token

// 更新token
