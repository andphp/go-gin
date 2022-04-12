package goGo

import "github.com/gin-gonic/gin"

type GoGo struct {
	*gin.Engine                  //我们把 engine放到 主类里
	g           *gin.RouterGroup //这里就是保存 group对象
}

func NewGoGo() *GoGo {
	return &GoGo{
		Engine: gin.New(),
	}
}

func (g *GoGo) Run() {
	g.Engine.Run(":8080")
}

func (g *GoGo) Mount(controller ...ControllerInterface) *GoGo {
	for _, c := range controller {
		c.Build(g)
	}
	return g
}
