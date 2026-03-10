package models

type Event struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Category    string  `json:"category"`
	Capacity    int     `json:"capacity"`
	Price       float64 `json:"price"`
	EventDate   string  `json:"event_date"`
	CreatedAt   string  `json:"created_at"`
}
