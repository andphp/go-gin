package goby

import "github.com/gin-gonic/gin"

type RouteGroupOption struct {
	apply func(*gin.RouterGroup)
}
