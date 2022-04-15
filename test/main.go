package main

import (
	"github.com/andphp/go-gin/goby"
	"github.com/andphp/go-gin/test/router"
)

func main() {
	goby.MakeGin().Mount("v1", router.NewIndexRouter()).Run()
}
