package goby

import (
	"bytes"
	"log"
	"os"

	"github.com/andphp/go-gin/goby/echo"
	"github.com/gin-gonic/gin"

	"runtime"
)

const (
	HTTP_STATUS = "HTTP_STATUS"
)

func panicTrace(kb int) string {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return string(stack)
}
func printError(err interface{}) {
	if os.Getenv("GIN_MODE") == "release" {
		return
	}
	log.Println(err)
}
func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				if os.Getenv("GIN_MODE") != "release" {
					log.Println("============================", panicTrace(10))
				}
				status := 400 //default status==400
				var code int = 999
				var msg string = "Failed"
				var errorData interface{} = e

				if strE, ok := e.(string); ok {
					printError(strE)
					errorData = strE
				} else if err, ok := e.(echo.ErrorStruct); ok {
					status = 200
					code = err.Code
					msg = err.Message
					errorData = err.Error
				} else if pe, ok := e.(error); ok {
					errorData = pe.Error()
				} else {
					errorData = e
				}
				if value, exists := context.Get(HTTP_STATUS); exists {
					if v, ok := value.(int); ok {
						status = v
					}
				}
				context.AbortWithStatusJSON(status, gin.H{"msg": msg, "code": code, "data": errorData})
			}
		}()
		context.Next()
	}
}
func Throw(err string, code int, context *gin.Context) {
	context.Set(HTTP_STATUS, code)
	panic(err)
}
func Error(err error, msg ...string) {
	if err == nil {
		return
	} else {
		errMsg := err.Error()
		if len(msg) > 0 {
			errMsg = msg[0]
		}
		panic(errMsg)
	}
}
