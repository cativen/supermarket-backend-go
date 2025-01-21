package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusBadRequest, 400, data, msg)
}

func Error(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, msg)
}
