package router

import (
	"github.com/andphp/go-gin/test/controller"
	"github.com/gin-gonic/gin"
)

func Test(r *gin.RouterGroup) {

	r.Group("test").GET("/w", controller.NewIndexCalss().GetIndex())
}
