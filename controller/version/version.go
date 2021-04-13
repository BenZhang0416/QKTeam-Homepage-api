package version

import (
	"api/utils"
	"github.com/gin-gonic/gin"
)

func Version(ctx *gin.Context) {
	utils.Success(ctx, map[string]interface{}{
		"version": "v1.0",
	})
}
