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
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "success"})
	}
}

func (ain *IndexCalss) Build(G *goby.Goby) {
	G.Handle("GET", "/", ain.GetIndex())
}
