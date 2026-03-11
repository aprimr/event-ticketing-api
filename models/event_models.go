package models

import "time"

type Event struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Category    string    `json:"category"`
	Capacity    int       `json:"capacity"`
	Price       float64   `json:"price"`
	EventDate   time.Time `json:"event_date"`
	CreatedAt   time.Time `json:"created_at"`
}

type PaginatedEvent struct {
	Events     []Event `json:"events"`
	Page       int     `json:"page"`
	Limit      int     `json:"limit"`
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
}
