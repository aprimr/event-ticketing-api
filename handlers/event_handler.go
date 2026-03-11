package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/aprimr/event-ticketing-api/models"
	"github.com/aprimr/event-ticketing-api/repository"
	"github.com/aprimr/event-ticketing-api/utils"
)

func AddEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	// Decode JSON
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("AddEventHandler -> JSON error: %v", err)
		utils.SendErrorResposnse(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Validate decoded data
	if strings.TrimSpace(event.Title) == "" {
		utils.SendErrorResposnse(w, "Title is required", http.StatusBadRequest)
		return
	}
	if len(event.Title) < 8 {
		utils.SendErrorResposnse(w, "Title must be atleast 8 characters long", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(event.Location) == "" {
		utils.SendErrorResposnse(w, "Location is required", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(event.Category) == "" {
		utils.SendErrorResposnse(w, "Category is required", http.StatusBadRequest)
		return
	}
	if event.Capacity == 0 {
		utils.SendErrorResposnse(w, "Capacity is required", http.StatusBadRequest)
		return
	}
	if event.Price == 0 {
		utils.SendErrorResposnse(w, "Price is required", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(event.EventDate) == "" {
		utils.SendErrorResposnse(w, "Event date is required", http.StatusBadRequest)
		return
	}

	// Call AddEvent
	err = repository.AddEvent(r.Context(), event)
	if err != nil {
		log.Printf("AddEventHandler -> db error: %v", err)
		utils.SendErrorResposnse(w, "Error creating event", http.StatusInternalServerError)
		return
	}

	// Send Success response
	utils.SendSuccessResposnse(w, "Event created successfully", nil, http.StatusCreated)
}
