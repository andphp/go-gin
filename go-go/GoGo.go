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

func (ain *GoGo) Run() {
	ain.Engine.Run(":8080")
}

func (ain *GoGo) Handle(method, path string, handler gin.HandlerFunc) *GoGo {
	ain.g.Handle(method, path, handler)
	return ain
}

func (ain *GoGo) Mount(group string, controller ...ControllerInterface) *GoGo {
	ain.g = ain.Group(group)
	for _, c := range controller {
		c.Build(ain)
	}
	return ain
}
