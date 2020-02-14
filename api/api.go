package api

import (
	"capulus/api/manage"
	"capulus/package/crons"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// Routing by grouping
	api := r.Group("/api")
	{
		manage.ApplyRoutes(api)
	}

	// system
	go func() {
		crons.Run()
	}()

	return r
}
