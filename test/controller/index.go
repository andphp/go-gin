package controller

import (
	goGo "github.com/andphp/go-gin/go-go"
	"github.com/gin-gonic/gin"
)

type IndexCalss struct {
}

func NewIndexCalss() *IndexCalss {
	return &IndexCalss{}
}

func (ain *IndexCalss) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "success"})
	}
}

func (ain *IndexCalss) Build(GoGo *goGo.GoGo) {
	GoGo.Handle("GET", "/", ain.GetIndex())
}
