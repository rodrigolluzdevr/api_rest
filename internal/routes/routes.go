package routes

import (
	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()
	router.GET("/get/events", getEvents)
	router.GET("/get/user", getUser)
	router.POST("/post/events", postEvents)
	router.POST("/post/user", postUser)

	router.Run()
}
