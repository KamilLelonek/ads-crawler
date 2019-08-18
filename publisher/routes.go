package publisher

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MountRoutes(group *gin.RouterGroup) {
	group.GET("/publisher/:name", fetchPublisher)
}

func fetchPublisher(c *gin.Context) {
	publisher := c.Param("name")
	rows, err := FetchList(publisher)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"publisher": rows})
	}
}
