package v1

import (
	"github.com/gin-gonic/gin"
	"go_admin/Server/global/response"
	"go_admin/Server/model/request"
	"go_admin/Server/utils"
)

func UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	 c.ShouldBindJSON(&cmr)

	err := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])

	if err != nil {
		response.FailWithMsg(err.Error(),c)
		return
	}
	//service.Update
}
