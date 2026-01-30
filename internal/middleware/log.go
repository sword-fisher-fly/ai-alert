package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// GinZapLogger returns a gin.HandlerFunc that logs requests using zap
func GinZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		message := c.Errors.ByType(gin.ErrorTypePrivate).String()

		token := c.Request.Header.Get("Authorization")
		ctx := logx.ContextWithFields(context.Background(),
			logx.Field("method", method),
			logx.Field("path", path),
			logx.Field("status", status),
			logx.Field("clientIP", clientIP),
			logx.Field("latency", latency),
			logx.Field("token", token),
		)
		logc.Info(ctx, message)
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			logc.Errorf(context.Background(), err.Error())
			c.Abort()
			return
		}
		// !!Important: set the request body back to the request
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		fmt.Println("Body:", string(bodyBytes))
		fmt.Println("Query Params:", c.Request.URL.Query())
		c.Next()
	}
}
func RequestLoggerFormatter(param gin.LogFormatterParams) string {
	level := "info"
	switch {
	case param.StatusCode >= 500:
		level = "error"
	case param.StatusCode >= 400:
		level = "warn"
	case param.StatusCode >= 300:
		level = "debug"
	}

	logData := map[string]interface{}{
		"level":      level,
		"statusCode": param.StatusCode,
		"clientIP":   param.ClientIP,
		"method":     param.Method,
		"path":       param.Path,
		"time":       param.TimeStamp.Format(time.RFC3339),
	}

	jsonData, _ := sonic.Marshal(logData)
	return fmt.Sprintf("%s\n", jsonData)
}
