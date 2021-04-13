package user

import (
	"api/boot/orm"
	"github.com/gin-gonic/gin"
)

type ViewValidate struct {
	View bool `json:"view" binding:"required `
}

//中间件，验证登录

func View(c *gin.Context) {
	var data = ViewValidate
	err := c.ShouldBindQuery(&data)
	if err != nil {
		c.JSON(422, gin.H{"message": "发生错误"})
	}

	db := orm.GetDB()
}
