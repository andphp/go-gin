package router

import (
	"github.com/andphp/go-gin/test/controller"
	"github.com/gin-gonic/gin"
)

func Api(r *gin.RouterGroup) {
	r.Group("vaq").GET("/w", controller.NewIndexCalss().GetIndex())
}
