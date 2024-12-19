package routes

import (
	"database/sql"
	"hms/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EventsIndex(ctx *gin.Context) {
	// state := ctx.Query("state")

	// fetch popular hotels in state=lagos
	var event models.Event
	popular_events, err := event.GetPopularEventsIn("lagos")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error fetching popular events in " + "lagos",
		})
		return
	}
	// get online, muisc and business events
	online_events := event.GetEventsByFormat("online")
	music_events := event.GetEventsByCategory("music")
	business_events := event.GetEventsByCategory("business")

	ctx.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"popular_events":  popular_events,
			"online_events":   online_events,
			"music_events":    music_events,
			"business_events": business_events,
		},
	})

}

func EventsSearch(ctx *gin.Context) {
	state := ctx.Query("state")
	var event models.Event
	events := event.GetEventsByState(state)

	ctx.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

// event details
func EventDetail(ctx *gin.Context) {
	event_id := ctx.Params.ByName("event_id")
	if event_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid hotel_id",
		})
		return
	}
	var event models.Event

	err := event.GetEventByID(event_id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": "no event with such id",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	// get events in the same category
	similarEvents := event.GetEventsByCategory(event.Category)
	limit := 5
	if len(similarEvents) < limit {
		limit = len(similarEvents)
	}
	similar := similarEvents[:limit]
	ctx.JSON(http.StatusOK, gin.H{
		"event":   event,
		"similar": similar,
	})
}
