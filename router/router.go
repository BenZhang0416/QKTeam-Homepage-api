package router

import (
	"api/boot/http"
	"api/controller/auth"
	"api/controller/user"
	"api/controller/version"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/qiankaihua/ginDemo/Boot/Http"
)

func Init() {
	AddRouter()
	AddStaticRoute()
}

// Add your route here
func AddRouter() {
	http.Router.GET("/", func(c *gin.Context) {
		c.JSON(200 ,gin.H{
			"version":"v1.0",
		})
	})
	Auth := Http.Router.Group("/Auth")
	{
		Auth.POST("login", auth.login)
	}
}

