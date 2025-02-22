package routes

import (
	"rodrigolluzdevr/api_rest/internal/handlers"

	"github.com/gin-gonic/gin"
)

func EventRoutes(router *gin.Engine) {
	eventRoutes := router.Group("events/add")
	{
		eventRoutes.GET("/", handlers.GetEvents)
		eventRoutes.POST("/", handlers.CreateEvent)
	}
}
