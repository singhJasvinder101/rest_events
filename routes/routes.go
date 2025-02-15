package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/singhJasvinder101/models"
)

func getEvents(context *gin.Context){
    events, err := models.GetallEvents()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "message": "Internal Server Error",
        })
        return
    }
    context.JSON(http.StatusOK, gin.H{
        "events": events,
    })
}

func createEvent(c *gin.Context){
    var event models.Event
    // pointer of event variable is passed incoming req must be of same type
    err := c.ShouldBindJSON(&event)
    if err != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Please enter the required fields",
        })
        return
    }

    event.Id = 1
    event.UserId = 1

    event.Save()

    c.JSON(http.StatusCreated, gin.H{
        "message": "Event Created!!", "event": event,
    })
}

func deleteAllEvents(c *gin.Context) {
    err := models.DeleteAllRows()
    if err != nil {
        fmt.Println("Error deleting events:", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Internal Server Error",
            "error":   err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "All events deleted",
    })
}

func getEventById(c *gin.Context){
    id := c.Param("id")
    event, err := models.GetEventById(id)
    if err != nil {
        fmt.Println("Error fetching event:", err)
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "Internal Server Error",
            "error":   err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "event": event,
    })
}


func deleteEventById(c *gin.Context){
	id := c.Param("id")
	err := models.DeleteEventById(id)
	if err != nil {
		fmt.Println("Error deleting event:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted",
	})
}


func updateEventById(c *gin.Context) {
	id := c.Param("id")
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println("Error binding event:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
    parsedId, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        fmt.Println("Error parsing event ID:", err)
        c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
    }
    event.Id = parsedId
	err = event.Update()
	if err != nil {
		fmt.Println("Error updating event:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
		"event":   event,
	})

}

