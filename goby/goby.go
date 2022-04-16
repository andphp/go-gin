package goby

import (
	"github.com/gin-gonic/gin"
)

type Goby struct {
	*gin.Engine                  //我们把 engine放到 主类里
	RouteGroup  *gin.RouterGroup //这里就是保存 group对象
}
type RouteGroupOption struct {
	apply func(*gin.RouterGroup)
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

func (ain *Goby) RouterMount(groupName string, middlebrows ...gin.HandlerFunc) func(opts ...func(*gin.RouterGroup)) *Goby {
	ain.RouteGroup = ain.Group(groupName)
	return func(opts ...func(*gin.RouterGroup)) *Goby {
		for _, middlewareOne := range middlebrows {
			ain.RouteGroup.Use(middlewareOne)
		}
		for _, option := range opts {
			routeGroupOption := &RouteGroupOption{apply: option}
			routeGroupOption.apply(ain.RouteGroup)
		}
		return ain
	}
}
