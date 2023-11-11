package contract

import "time"

type CreateEventRequest struct {
	Name                string    `json:"name" binding:"required"`
	IsOneTimeOccurrence bool      `json:"isOneTimeOccurrence" binding:"required"`
	StartTime           time.Time `json:"startTime" binding:"required"`
	EndTime             time.Time `json:"endTime"`
}

type CreateEventResponse struct {
	ID                  int32      `json:"id"`
	Name                string     `json:"name"`
	IsOneTimeOccurrence bool       `json:"isOneTimeOccurrence"`
	StartTime           time.Time  `json:"startTime"`
	EndTime             *time.Time `json:"endTime,omitempty"`
}

type GetEventRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

type GetEventsRequest struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}
