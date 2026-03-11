package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aprimr/event-ticketing-api/db"
	"github.com/aprimr/event-ticketing-api/handlers"
	"github.com/aprimr/event-ticketing-api/utils"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}
	mux := http.NewServeMux()

	// Connect to db
	db.ConnectDB()

	// routes
	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// POST /events (add event)
		case http.MethodPost:
			handlers.AddEventHandler(w, r)

		// GET /events?page=1&limit=5 (fetch events)
		case http.MethodGet:
			handlers.FetchEventsHandler(w, r)

		// Handle default case
		default:
			utils.SendErrorResposnse(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/events/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET /events/:id (fetch event by id)
		case http.MethodGet:
			handlers.FetchEventByIdHandler(w, r)

		// PUT /events/:id (update event by id)
		case http.MethodPut:
			handlers.UpdateEventByIdHandler(w, r)

		// DELETE /events/:id (delete event by id)
		case http.MethodDelete:
			handlers.DeleteEventByIdHandler(w, r)

		// Handle default case
		default:
			utils.SendErrorResposnse(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// SpinUp server
	port := ":" + os.Getenv("PORT")
	log.Println("Server started on port", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Unexpected error occured")
	}
}
