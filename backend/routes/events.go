package routes

import (
	"hms/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EventsSearch(ctx *gin.Context) {
	state := ctx.Query("state")
	var event models.Event
	events := event.GetEventsByState(state)

	ctx.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}
