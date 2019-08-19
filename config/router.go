package config

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(mountRoutes func(*gin.RouterGroup)) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	mountRoutes(v1)

	return router
}
