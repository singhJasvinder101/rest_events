package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){

	// part of same package so all lowercase
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.DELETE("/events", deleteAllEvents)
	server.GET("/events/:id", getEventById)
	server.DELETE("/events/:id", deleteEventById)
	server.PUT("/events/:id", updateEventById)

	server.POST("/signup", signup)
	server.POST("/login", login)
}