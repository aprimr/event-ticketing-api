package repository

import (
	"context"
	"fmt"
	"math"

	"github.com/aprimr/event-ticketing-api/db"
	"github.com/aprimr/event-ticketing-api/models"
	"github.com/jackc/pgx/v5"
)

func AddEvent(ctx context.Context, event models.Event) error {
	_, err := db.Pool.Exec(ctx, "INSERT INTO events (title, description, location, category, capacity, price, event_date) VALUES($1, $2, $3, $4, $5, $6, $7)", event.Title, event.Description, event.Location, event.Category, event.Capacity, event.Price, event.EventDate)
	return err
}

func FetchEvents(ctx context.Context, page int, limit int, category string, event_date string) (*models.PaginatedEvent, error) {
	// Calculate offset
	offset := (page - 1) * limit

	// Get total no of rows
	var totalRows int
	err := db.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM events").Scan(&totalRows)
	if err != nil {
		return nil, err
	}

	// Base query
	query := "SELECT id, title, description, location, category, capacity, price, event_date, created_at FROM events WHERE 1=1"
	args := []any{}
	argsCount := 1

	// Add `category and date` filters if exists
	if category != "" {
		query += fmt.Sprintf(" AND category=$%d", argsCount)
		args = append(args, category)
		argsCount++
	}
	if event_date != "" {
		query += fmt.Sprintf(" AND event_date=$%d", argsCount)
		args = append(args, event_date)
		argsCount++
	}

	// Add pagination at end
	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", argsCount, argsCount+1)
	args = append(args, limit, offset)

	// Query DB
	rows, err := db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through rows and scan all rows
	events := []models.Event{}

	for rows.Next() {
		var event models.Event

		err = rows.Scan(&event.Id, &event.Title, &event.Description, &event.Location, &event.Category, &event.Capacity, &event.Price, &event.EventDate, &event.CreatedAt)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))
	paginaginatedEvents := models.PaginatedEvent{
		Events:     events,
		Page:       page,
		Limit:      limit,
		TotalCount: totalRows,
		TotalPages: totalPages,
	}

	return &paginaginatedEvents, nil
}

func FetchEventById(ctx context.Context, id int) (*models.Event, error) {
	event := models.Event{}

	// Query db
	row := db.Pool.QueryRow(ctx, "SELECT id, title, description, location, category, capacity, price, event_date, created_at FROM events WHERE id=$1", id)
	err := row.Scan(&event.Id, &event.Title, &event.Description, &event.Location, &event.Category, &event.Capacity, &event.Price, &event.EventDate, &event.CreatedAt)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("No rows with the provided id matched")
	}
	return &event, nil
}

func UpdateEventById(ctx context.Context, id int, event models.Event) (*models.Event, error) {
	// Store updated event
	updatedEvent := models.Event{}

	row := db.Pool.QueryRow(ctx, "UPDATE events SET title=$1, description=$2, location=$3, category=$4, capacity=$5, price=$6, event_date=$7 WHERE id=$8 RETURNING id, title, description, location, category, capacity, price, event_date, created_at", event.Title, event.Description, event.Location, event.Category, event.Capacity, event.Price, event.EventDate, id)
	err := row.Scan(&updatedEvent.Id, &updatedEvent.Title, &updatedEvent.Description, &updatedEvent.Location, &updatedEvent.Category, &updatedEvent.Capacity, &updatedEvent.Price, &updatedEvent.EventDate, &updatedEvent.CreatedAt)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("No events matched with provided id")
	}
	if err != nil {
		return nil, err
	}
	return &updatedEvent, nil
}
