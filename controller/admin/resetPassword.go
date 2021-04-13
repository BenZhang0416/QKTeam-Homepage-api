package admin

import (
	"api/boot/orm"
	"api/controller"
	"api/model"
	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} user/security 管理员重置密码-ResetPassword
 * @apiGroup User
 * @apiName ResetPassword
 * @apiPermission User/Admin
 * @apiParam {string} password_new 新密码
 * @apiParamExample {json} Request-Example:
 * {
 *      'password_old': 'ce8749330e860006d92f26528071b26bc93d234d',
 *      'password_new': 'dc89fbb8f0bdb28a3755743032f8ab05f0e0b77d'
 * }
 */

type ResetPasswordValidate struct {
	PasswordNew string `json:"password_new" binding:"required,len=40"`
}

func ResetPassword(c *gin.Context) {
	_user, _ := c.Get("user")
	if _user == nil {
		c.AbortWithStatus(401)
		return
	}
	user := _user.(model.User)
	var data ResetPasswordValidate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": "格式不正确"})
		return
	}
	db := orm.GetDB()
	if !controller.Sha256Check(user.Password, data.PasswordOld) {
		c.JSON(403, gin.H{"message": "密码错误"})
		return
	}
	//重置新密码
	db.Model(&user).UpdateColumn("Password", controller.Sha256Get(data.PasswordNew))
}
