package main

import (
	"github.com/gin-gonic/gin"
	"github.com/singhJasvinder101/db"
	"github.com/singhJasvinder101/routes"
)

func main()  {
    db.Init()
    // returns a instance of the Gin engine with logger & recovery middleware already attached
    // pre-configured server on top of http
    server := gin.Default()

    server.GET("/", func(c *gin.Context) {
        // gin.H is a shortcut for map[string]any
        // gin.H{} is a slice
        c.JSON(200, gin.H{
            "message": "Hello World!",
        })
    })

    routes.RegisterRoutes(server)

    server.Run(":3000")
}   

