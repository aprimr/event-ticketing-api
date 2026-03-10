package repository

import (
	"context"

	"github.com/aprimr/event-ticketing-api/db"
	"github.com/aprimr/event-ticketing-api/models"
)

func AddEvent(ctx context.Context, event models.Event) error {
	_, err := db.Pool.Exec(ctx, "INSERT INTO events (title, description, location, category, capacity, price, event_date) VALUES($1, $2, $3, $4, $5, $6, $7)", event.Title, event.Description, event.Location, event.Category, event.Capacity, event.Price, event.EventDate)
	return err
}
