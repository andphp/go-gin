package middleware

import (
	"github.com/andphp/go-gin/goby/echo"
	"github.com/andphp/go-gin/test/model"
	"github.com/gin-gonic/gin"
)

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参数验证
		c.Next()
		var l model.SysUserModel
		echo.Result(c.ShouldBindQuery(&l)).Unwrap(echo.PARAMS_ERROR)
		// do something

	}
}
