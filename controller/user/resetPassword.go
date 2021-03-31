package user

import (
	"api/boot/orm"
	"api/controller"
	"api/model"
	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} user/security 修改密码-ResetPassword
 * @apiGroup User
 * @apiName ResetPassword
 * @apiPermission User/Admin
 * @apiParam {string} password_old 旧密码
 * @apiParam {string} password_new 新密码
 * @apiParamExample {json} Request-Example:
 * {
 *      'password_old': 'ce8749330e860006d92f26528071b26bc93d234d',
 *      'password_new': 'dc89fbb8f0bdb28a3755743032f8ab05f0e0b77d'
 * }
 */

type ResetPasswordValidate struct {
	PasswordOld string `json:"password_old" binding:"required,len=40"`
	PasswordNew string `json:"password_new" binding:"required,len=40"`
}

func ResetPassword(c *gin.Context)  {
	_user, _ := c.Get("user")
	if _user == nil {
		c.AbortWithStatus(401)
		return
	}
	user := _user.(model.user)
	var data ResetPasswordValidate
	if err :=c.ShouldBindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message":"格式不正确哦"})
		return
	}
	db := orm.GetDB()
	if !controller.Sha256Check(user.Password, data.PasswordOld) {
		c.JSON(403, gin.H{"message":"密码错误"})
		return
	}
	db.Model(&user).UpdateColumn("Password", controller.Sha256Get(data.PasswordNew))
}