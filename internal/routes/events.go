package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Get event request success"})
}

func postEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Post event request success"})
}
