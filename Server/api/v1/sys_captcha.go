package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go_admin/Server/global"
	"go_admin/Server/global/response"
	response2 "go_admin/Server/model/response"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMsg(fmt.Sprintf("获取数据失败 ，%v", err), c)
	} else {
		response.OkWithData(response2.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, c)
	}

}
