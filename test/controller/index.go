package controller

import (
	"github.com/andphp/go-gin/goby"
	"github.com/gin-gonic/gin"
)

type IndexCalss struct {
}

func NewIndexCalss() *IndexCalss {
	return &IndexCalss{}
}

func (ain *IndexCalss) GetIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var l model.SysUserModel
		//echo.Result(c.ShouldBind(&l)).Unwrap(echo.PARAMS_ERROR)
		//panic("这就是个错误")
		c.JSON(200, gin.H{"result": "success"})
	}
}

func (ain *IndexCalss) Build(G *goby.Goby) {
	G.Handle("GET", "/", ain.GetIndex())
}
