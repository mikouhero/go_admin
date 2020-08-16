package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回体结构
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "ok", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "ok", c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "fail", c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "fail", c)
}

func FailWithDetail(code int, data interface{}, msg string, c *gin.Context) {
	Result(code, data, msg, c)
}
