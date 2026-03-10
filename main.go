package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aprimr/event-ticketing-api/db"
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

	// SpinUp server
	port := ":" + os.Getenv("PORT")
	log.Println("Server started on port", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Unexpected error occured")
	}
}
