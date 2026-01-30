package middleware

import (
	"github.com/sword-fisher-fly/ai-alert/internal/ctx"
	"github.com/sword-fisher-fly/ai-alert/pkg/response"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(context)
			context.Abort()
			return
		}

		code, ok := IsTokenValid(ctx.DO(), tokenStr)
		if !ok {
			if code == 401 {
				response.TokenFail(context)
				context.Abort()
				return
			}
		}

	}
}

// Mock here. Any token is valid.
func IsTokenValid(ctx *ctx.Context, tokenStr string) (int64, bool) {
	// 1) check token when first business request is coming
	// Parse token and set token with expiration into redis if valid.
	// 2) check token whether it exists in redis
	// if exists, return true
	return 200, true
}
