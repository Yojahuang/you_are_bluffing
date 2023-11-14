package Router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var value = []int{
	1,
	2,
}

// getAlbums responds with the list of all albums as JSON.
func getValue(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, value)
}

func RouterFactory() *gin.Engine {
	router := gin.Default()

	router.GET("/value", getValue)

	return router
}
