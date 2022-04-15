package goby

import "github.com/gin-gonic/gin"

type Goby struct {
	*gin.Engine                  //我们把 engine放到 主类里
	RouteGroup  *gin.RouterGroup //这里就是保存 group对象
}

func MakeGin() *Goby {
	return &Goby{
		Engine: gin.New(),
	}
}

func (ain *Goby) Run() {
	ain.Engine.Run(":8080")
}

func (ain *Goby) Handle(method, path string, handler gin.HandlerFunc) *Goby {
	ain.RouteGroup.Handle(method, path, handler)
	return ain
}

func (ain *Goby) Mount(group string, routers ...RouterInterface) *Goby {
	ain.RouteGroup = ain.Group(group)
	for _, c := range routers {
		c.Build(ain)
	}
	return ain
}
