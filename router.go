package main

import (
	"github.com/gin-gonic/gin"
)

const apiVersion = "/v1"

func SetupRouter() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	group := router.Group(apiVersion)

	return router, group
}
