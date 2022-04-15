package router

import (
	"github.com/andphp/go-gin/goby"
	"github.com/andphp/go-gin/test/controller"
)

type IndexRouter struct{}

func NewIndexRouter() *IndexRouter {
	return &IndexRouter{}
}

func (ain *IndexRouter) Build(Goby *goby.Goby) {
	Goby.RouteGroup.Group("va").GET("/", controller.NewIndexCalss().GetIndex())
}
