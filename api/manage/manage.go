package manage

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	manage := r.Group("/manage")
	{
		manage.GET("/", GetHosts)
		manage.POST("/", AddHost)
		manage.DELETE("/:id", DeleteHost)
	}
}
