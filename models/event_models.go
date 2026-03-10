package models

type Event struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    int    `json:"location"`
	Category    int    `json:"category"`
	Capacity    int    `json:"capacity"`
	Price       int    `json:"price"`
	EventDate   int    `json:"event_date"`
	CreatedAt   int    `json:"created_at"`
}
