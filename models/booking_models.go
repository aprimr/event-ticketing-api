package models

import "time"

type Booking struct {
	Id        int       `json:"id"`
	EventId   int       `json:"event_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Seats     int       `json:"seats"`
	CreatedAt time.Time `json:"created_at"`
}
