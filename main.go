package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/noodlecak-e/timelines/db/sqlc"
	"github.com/noodlecak-e/timelines/internal/resource"
)

func main() {
	connectionStr := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	queries := sqlc.New(conn)

	r := gin.Default()

	event := resource.NewEventResource(queries)

	r.POST("/events", event.CreateEvent)
	r.GET("/events/:id", event.GetEvent)
	r.GET("/events", event.GetEvents)

	if err = r.Run(":8080"); err != nil {
		panic(err)
	}
}
