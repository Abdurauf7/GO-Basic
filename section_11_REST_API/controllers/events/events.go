package events

import (
	"net/http"
	"practice/Rest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't  fetch events"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "err": err})
		return
	}

	user_id := ctx.GetInt64("user_id")

	event.UserId = user_id

	err = event.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event.", "err": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func GetEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get id of event"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get event"})
		return
	}

	ctx.JSON(http.StatusOK, event)

}

func UpdateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get id of event"})
		return
	}

	user_id := ctx.GetInt64("user_id")
	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get event"})
		return
	}

	if event.UserId != user_id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}
	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get id of event"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func DeleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get id of event"})
		return
	}
	user_id := ctx.GetInt64("user_id")
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get event"})
		return
	}
	if event.UserId != user_id {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't get event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})

}
