package main

import (
	"github.com/andphp/go-gin/goby"
	"github.com/andphp/go-gin/test/router"
)

func main() {
	goby.MakeGin().RouterMount("v1")(router.Api, router.Test).Run()
}
