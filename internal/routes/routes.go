package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()
	router.GET("/get", getEvents)

	router.Run()
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Get request success"})
}
