package routes

import (
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// events routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// registration routes
	authenticated.POST("/events/:id/register", register)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// user routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
