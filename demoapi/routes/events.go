package routes

import (
	"demoapi/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Could not fetch events, try again later!",
			},
		)
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse event id!",
			},
		)
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not find event with id!",
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"event": event,
		},
	)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse this request.",
			"stack":   err,
		})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Could not create event, try again later!",
			},
		)
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created!",
		"event":   event,
	})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse event id!",
			},
		)
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Could not fetch event!",
			},
		)
		return
	}

	if event.UserID != userId {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "You are not allowed to do that.",
			},
		)
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Could not update event, try again later!",
			},
		)
		return
	}

	updatedEvent.ID = int(eventId)
	err = updatedEvent.Update()
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Could not update event, try again later!",
			},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not parse event id!",
			},
		)
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Could not fetch event!",
			},
		)
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "You are not allowed to do that.",
			},
		)
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Could not delete event!",
			},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Event deleted.",
		},
	)
}
