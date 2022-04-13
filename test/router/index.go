package router

import (
	goGo "github.com/andphp/go-gin/go-go"
	"github.com/andphp/go-gin/test/controller"
)

type IndexRouter struct{}

func NewIndexRouter() *IndexRouter {
	return &IndexRouter{}
}

func (ain *IndexRouter) Build(GoGo *goGo.GoGo) {
	GoGo.GoGroup.Group("va").Handle("GET", "/", controller.NewIndexCalss().GetIndex())
}
