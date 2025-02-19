package registration

import (
	"net/http"
	"practice/Rest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctx *gin.Context) {
	user_id := ctx.GetInt64("user_id")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get id of event"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch event"})
		return
	}

	err = event.Register(user_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't register user for event"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event Registered"})

}

func CancelRegistration(ctx *gin.Context) {
	user_id := ctx.GetInt64("user_id")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't get id of event"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegister(user_id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't cancel registration"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registration Cancelled"})

}
