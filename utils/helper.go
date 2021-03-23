package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const ErrnoKey = "_errnoUnique"

func FailureCustomize(ctx *gin.Context, code int, msg string) {
	data := map[string]interface{}{
		"code":    0,
		"message": msg,
	}
	ctx.JSON(code, data)
}

func Failure(ctx *gin.Context, errno ErrCode) {
	errMsg, ok := ErrMsgInfo[errno]
	ctx.Set(ErrnoKey, strconv.Itoa(int(errno)))
	if !ok {
		ctx.JSON(int(errno), gin.H{"code": 0, "message": "unknown"})
	} else {
		ctx.JSON(errMsg.HttpCode, gin.H{"code": 0, "message": errMsg.ErrMsg})
	}
}

func Success(ctx *gin.Context, response map[string]interface{}) {
	data := map[string]interface{}{
		"code":    1,
		"message": "success",
	}
	ctx.JSON(200, MergeMap(data, response))
}

