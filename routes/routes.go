package routes

import (
	"github.com/gin-gonic/gin"
	"modern-delivery-service/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	//events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//users
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
