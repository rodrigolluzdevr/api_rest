package handlers

import (
	"net/http"
	"rodrigolluzdevr/api_rest/internal/domain"
	"rodrigolluzdevr/api_rest/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var event domain.Event

	// verify if JSON is correct
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// implemented the create services logic
	eventService := services.NewEventService()
	if err := eventService.CreateEvent(&event); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	// event created success
	context.JSON(http.StatusOK, gin.H{"message": "Event created successfully", "event": event})
}

func GetEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Get event request success"})
}
