package config

import (
	"crawler/publisher"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	publisher.MountRoutes(v1)

	return router
}
