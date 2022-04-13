package goGo

import "github.com/gin-gonic/gin"

type GoGo struct {
	*gin.Engine                  //我们把 engine放到 主类里
	GoGroup     *gin.RouterGroup //这里就是保存 group对象
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
	ain.GoGroup.Handle(method, path, handler)
	return ain
}

func (ain *GoGo) Mount(group string, routers ...RouterInterface) *GoGo {
	ain.GoGroup = ain.Group(group)
	for _, c := range routers {
		c.Build(ain)
	}
	return ain
}
