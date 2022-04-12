package main

import (
	goGo "github.com/andphp/go-gin/go-go"
	"github.com/andphp/go-gin/test/controller"
)

func main() {
	goGo.NewGoGo().Mount("v1", controller.NewIndexCalss()).Run()
}