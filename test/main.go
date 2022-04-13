package main

import (
	goGo "github.com/andphp/go-gin/go-go"
	"github.com/andphp/go-gin/test/router"
)

func main() {
	goGo.NewGoGo().Mount("v1", router.NewIndexRouter()).Run()
}
