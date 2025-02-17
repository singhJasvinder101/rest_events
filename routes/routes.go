package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/singhJasvinder101/middleware"
)

func RegisterRoutes(server *gin.Engine){
	
	server.POST("/signup", signup)
	server.POST("/login", login)

	// part of same package so all lowercase
	protected := server.Group("/events")
	protected.Use(middleware.Authenticate)

	protected.GET("", getEvents)
	protected.POST("", createEvent)
	protected.DELETE("", deleteAllEvents)
	protected.GET("/:id", getEventById)
	protected.DELETE("/:id", deleteEventById)
	protected.PUT("/:id", updateEventById)
}