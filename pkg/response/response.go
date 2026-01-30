package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var CodeInfo = map[int64]string{
	200: "OK",
	400: "request params error",
	401: "token authentication failed",
	403: "authorization failed",
}

func Response(c *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data interface{}, msg string) {
	code := 400
	Response(ctx, code, code, data, msg)
}

func TokenFail(ctx *gin.Context) {
	code := 401
	Response(ctx, code, code, CodeInfo[int64(code)], "failed")
}

func PermissionFail(ctx *gin.Context) {
	code := 403
	Response(ctx, code, code, CodeInfo[int64(code)], "failed")
}
