package repository

import (
	"context"
	"fmt"

	"github.com/aprimr/event-ticketing-api/db"
	"github.com/aprimr/event-ticketing-api/models"
	"github.com/jackc/pgx/v5"
)

func AddBooking(ctx context.Context, event_id int, booking models.Booking) error {
	// Start transaction
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Check available seats
	var availableSeats int
	row := tx.QueryRow(ctx, "SELECT e.capacity - COALESCE(SUM(b.seats), 0) AS available_seats FROM events e LEFT JOIN bookings b ON e.id = b.event_id WHERE id=$1 GROUP BY e.capacity")
	err = row.Scan(&availableSeats)
	if err == pgx.ErrNoRows {
		return fmt.Errorf("Event not found")
	}
	if err != nil {
		return err
	}

	// Check if enough seats are available
	if booking.Seats > availableSeats {
		return fmt.Errorf("Not enough seats")
	}

	// Add booking data into db
	_, err = tx.Exec(ctx, "INSERT INTO bookings (event_id, name, email, seats) VALUES($1, $2, $3, $4)", booking.EventId, booking.Name, booking.Email, booking.Seats)
	if err != nil {
		return err
	}

	// commit transaction
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
