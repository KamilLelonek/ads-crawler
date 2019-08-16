package publisher

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MountRoutes(group *gin.RouterGroup) {
	group.GET("/publisher/:name", FetchPublisher)
}

func FetchPublisher(c *gin.Context) {
	c.JSON(http.StatusOK, renderResult(c))
}

func renderResult(c *gin.Context) gin.H {
	return gin.H{"publisher": c.Param("name")}
}
