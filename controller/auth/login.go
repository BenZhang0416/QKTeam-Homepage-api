package auth

import (
	"api/boot/orm"
	"api/controller"
	"api/model"
	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} auth/login 登陆-Login
 * @apiGroup Auth
 * @apiName Login
 * @apiPermission All
 * @apiParam {string} password 密码(sha1加密)
 * @apiParam {string} username 名字
 * @apiParamExample {json} Request-Example:
 * {
 *      'username': 'administrator',
 *      'password': 'd033e22ae348aeb5660fc2140aec35850c4da997',
 * }
 * @apiParamExample {json} Request-Example2:
 * {
 *      'username': 'administrator',
 *      'password': 'd033e22ae348aeb5660fc2140aec35850c4da997'
 * }
 * @apiSuccess {string} token Api-Token
 * @apiSuccessExample {json} Success-response:
 * {
 *     'token': 'b2336207-3136-47aa-9362-de45f3e49e65'
 * }
 */

//无记住我功能
type LoginValidate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,len=40"`
}

func Login(c *gin.Context) {
	var data LoginValidate
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(422, gin.H{"message": "格式有误"})
		return
	}
	var user model.User
	db := orm.GetDB()
	//未获取到输入内容
	if db.Where("username = ?", data.Username).First(&user).RecordNotFound() {
		c.JSON(401, gin.H{"message": "用户名或密码错误"})
		return
	}
	//用户名和密码不匹配
	if !controller.Sha256Check(user.Password, data.Password) {
		c.JSON(401, gin.H{"message": "用户名或密码错误"})
		return
	}

	//成功，反给token
	var apiToken = model.ApiToken{UserId: user.ID}
	db.Create(&apiToken)
	c.JSON(200, gin.H{"token": apiToken.Token})
}
