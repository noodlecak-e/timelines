package resource

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/noodlecak-e/timelines/contract"
	"github.com/noodlecak-e/timelines/db/sqlc"
)

type EventResource struct {
	queries *sqlc.Queries
}

func NewEventResource(queries *sqlc.Queries) *EventResource {
	return &EventResource{
		queries: queries,
	}
}

func (e EventResource) CreateEvent(c *gin.Context) {
	var req contract.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event, err := e.queries.CreateEvent(c, sqlc.CreateEventParams{
		Name:      req.Name,
		StartDate: req.StartTime,
		OneTime:   req.IsOneTimeOccurrence,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := contract.CreateEventResponse{
		ID:                  event.ID,
		Name:                event.Name,
		IsOneTimeOccurrence: event.OneTime,
		StartTime:           event.StartDate,
	}

	if !req.IsOneTimeOccurrence {
		resp.EndTime = &event.EndDate.Time
	}

	c.JSON(http.StatusOK, resp)
}

func (e EventResource) GetEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	event, err := e.queries.GetEvent(c, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (e EventResource) GetEvents(c *gin.Context) {
	var req contract.GetEventsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	events, err := e.queries.GetEvents(
		c,
		sqlc.GetEventsParams{
			Offset: req.Offset,
			Limit:  req.Limit,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, events)
}
