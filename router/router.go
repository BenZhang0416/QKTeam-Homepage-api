package router

import (
	"api/boot/http"
	"api/controller/applicationForm"
	"api/controller/auth"
	"github.com/gin-gonic/gin"
)

func Init() {
	AddRouter()
	AddStaticRoute()
}

// Add your route here
func AddRouter() {
	http.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "v1.0",
		})
	})
	aFormRouter := http.Router.Group("/aFormRouter")
	{
		aFormRouter.POST("Commit", applicationForm.Commit)
	}
	Auth := http.Router.Group("/Auth")
	{
		Auth.POST("login", auth.Login)
	}
}
