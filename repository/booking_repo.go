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
	row := tx.QueryRow(ctx, "SELECT e.capacity - COALESCE(SUM(b.seats), 0) AS available_seats FROM events e LEFT JOIN bookings b ON e.id = b.event_id WHERE e.id=$1 GROUP BY e.capacity", event_id)
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
	_, err = tx.Exec(ctx, "INSERT INTO bookings (event_id, name, email, seats) VALUES($1, $2, $3, $4)", event_id, booking.Name, booking.Email, booking.Seats)
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

func GetAllBookingsByEventId(ctx context.Context, EventId int) (*[]models.Booking, error) {
	rows, err := db.Pool.Query(ctx, "SELECT id, event_id, name, email, seats, created_at FROM bookings WHERE event_id=$1", EventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan each row
	bookings := []models.Booking{}
	for rows.Next() {
		booking := models.Booking{}
		err = rows.Scan(&booking.Id, &booking.EventId, &booking.Name, &booking.Email, &booking.Seats, &booking.CreatedAt)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return &bookings, nil
}

func DeleteBookingByEventIdAndBookingId(ctx context.Context, eventId int, bookingId int) error {
	commandTag, err := db.Pool.Exec(ctx, "DELETE FROM bookings WHERE id=$1 AND event_id=$2", bookingId, eventId)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("Booking not found")
	}
	return nil
}
