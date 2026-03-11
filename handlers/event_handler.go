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

func AddEventHandler(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}

	// Decode JSON
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Printf("AddEventHandler -> JSON error: %v \n", err)
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

	// Call AddEvent
	err = repository.AddEvent(r.Context(), event)
	if err != nil {
		log.Printf("AddEventHandler -> db error: %v \n", err)
		utils.SendErrorResposnse(w, "Error creating event", http.StatusInternalServerError)
		return
	}

	// Send Success response
	utils.SendSuccessResposnse(w, "Event created successfully", nil, http.StatusCreated)
}

func FetchEventsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse URL query
	page := utils.ParseQueryInt(r, "page", 1)
	limit := utils.ParseQueryInt(r, "limit", 10)
	category := utils.ParseQueryString(r, "category", "")
	date := utils.ParseQueryString(r, "date", "")

	// Validate
	if page <= 0 {
		utils.SendErrorResposnse(w, "Page cannot be less than zero", http.StatusBadRequest)
		return
	}
	if limit <= 0 || limit >= 200 {
		utils.SendErrorResposnse(w, "Limit cannot be less than zero or more than 200", http.StatusBadRequest)
		return
	}

	// Call FetchEvents
	events, err := repository.FetchEvents(r.Context(), page, limit, category, date)
	if err != nil {
		log.Printf("FetchEventHandler -> db error: %v \n", err)
		utils.SendErrorResposnse(w, "Error fetching events", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResposnse(w, "Events fetched successfully", events, http.StatusOK)
}

func FetchEventByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Parse URL
	idStr := strings.TrimPrefix(r.URL.Path, "/events/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("FetchEventByIdHandler -> bad request: %v \n", err)
		utils.SendErrorResposnse(w, "Event ID must be a number", http.StatusBadRequest)
		return
	}

	// Call FetchEventById
	event, err := repository.FetchEventById(r.Context(), id)
	if err != nil {
		log.Printf("FetchEventByIdHandler -> db error: %v \n", err)
		utils.SendErrorResposnse(w, "Event not found", http.StatusBadRequest)
		return
	}

	utils.SendSuccessResposnse(w, "Event fetched successfully", event, http.StatusOK)
}
