package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-blog-service/global"
	"go-blog-service/pkg/app"
	"go-blog-service/pkg/errcode"
	"go-blog-service/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err!= nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		// 格式化时间输出
		beginTime := time.Now().Format("2006/01/02 15:04:05")
		c.Next()
		endTime := time.Now().Format("2006/01/02 15:04:05")

		fields := logger.Fields{
			"request": c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}

		global.Logger.WithFields(fields).Infof(
			"access log: method %s, status_code:% d, begin_time: %s, end_time: %s",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
			)
	}
}

func Recovery() gin.HandlerFunc {
	return func (c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}