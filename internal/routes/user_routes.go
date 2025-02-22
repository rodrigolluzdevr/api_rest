package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Get user request success"})
}

func postUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Post user request success"})
}
