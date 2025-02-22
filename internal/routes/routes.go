package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Add routes
	EventRoutes(router)

	//return function
	return router
}
