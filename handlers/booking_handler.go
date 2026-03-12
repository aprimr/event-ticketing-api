package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aprimr/event-ticketing-api/models"
	"github.com/aprimr/event-ticketing-api/repository"
	"github.com/aprimr/event-ticketing-api/utils"
)

func AddBookingHandler(w http.ResponseWriter, r *http.Request) {
	// Parse URL
	urlStr := strings.TrimPrefix(r.URL.Path, "/events/")
	idStr := strings.TrimSuffix(urlStr, "/bookings")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("AddBookingHandler -> bad request: %v \n", err)
		utils.SendErrorResposnse(w, "Event ID must be a number", http.StatusBadRequest)
		return
	}

	// Decode JSON
	var booking models.Booking
	err = json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		log.Printf("AddBookingHandler -> JSON error: %v \n", err)
		utils.SendErrorResposnse(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate
	if strings.TrimSpace(booking.Name) == "" {
		utils.SendErrorResposnse(w, "Name is required", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(booking.Email) == "" {
		utils.SendErrorResposnse(w, "Email is required", http.StatusBadRequest)
		return
	}
	if booking.Seats <= 0 {
		utils.SendErrorResposnse(w, "Seats must me greater than 0", http.StatusBadRequest)
		return
	}

	// Call AddBooking
	err = repository.AddBooking(r.Context(), id, booking)
	if err != nil {
		log.Printf("AddBookingHandler -> db error: %v \n", err)
		if err.Error() == "Event not found" {
			utils.SendErrorResposnse(w, "Invalid event id", http.StatusNotFound)
			return
		}
		if err.Error() == "Not enough seats" {
			utils.SendErrorResposnse(w, "Not enough seats", http.StatusInternalServerError)
			return
		}
		utils.SendErrorResposnse(w, "Error creating booking", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResposnse(w, "Booking created", nil, http.StatusCreated)
}

func GetAllBookingsByEventIdHandler(w http.ResponseWriter, r *http.Request) {
	// Parse URL
	urlStr := strings.TrimPrefix(r.URL.Path, "/events/")
	idStr := strings.TrimSuffix(urlStr, "/bookings")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetAllBookingsByEventIdHandler -> bad request: %v \n", err)
		utils.SendErrorResposnse(w, "Event ID must be a number", http.StatusBadRequest)
		return
	}

	// Call GetBookingById
	bookings, err := repository.GetAllBookingsByEventId(r.Context(), id)
	if err != nil {
		log.Printf("GetAllBookingsByEventIdHandler -> db error: %v \n", err)
		utils.SendErrorResposnse(w, "Error fetching bookings", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResposnse(w, "Booking fetch successful", bookings, http.StatusOK)
}
