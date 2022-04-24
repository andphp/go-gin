package middleware

import (
	"github.com/andphp/go-gin/goby/echo"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				var code int = 999
				var msg string = "Failed"
				var errorData interface{} = e
				if err, ok := e.(echo.ErrorStruct); ok {
					code = err.Code
					msg = err.Message
					errorData = err.Error
				}
				context.JSON(200, gin.H{"msg": msg, "code": code, "data": errorData})
			}
		}()
		context.Next()
	}
}
