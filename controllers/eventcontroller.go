package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rahmet/database"
	"rahmet/models"
	"strconv"
)

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	database.Instance.Find(&events, "status = ?", "active")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func GetEventById(w http.ResponseWriter, r *http.Request) {
	eventId := mux.Vars(r)["id"]
	if checkIfEventExists(eventId) == false {
		json.NewEncoder(w).Encode("Event Not Found!")
		return
	}
	var event models.Event
	database.Instance.First(&event, eventId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func GetEventsByUserId(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	var events []models.Event
	database.Instance.Find(&events, "creator_id = ?", userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)
	database.Instance.Create(&event)
	json.NewEncoder(w).Encode(event)
}

func ParticipateEvent(w http.ResponseWriter, r *http.Request) {

	var participate models.Participate
	json.NewDecoder(r.Body).Decode(&participate)
	eventId := strconv.Itoa(int(participate.EventID))
	if checkIfEventExists(eventId) == false {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var event models.Event
	database.Instance.First(&event, eventId)

	if event.Limit < len(event.Participants)+1 {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode("Limit exceeded")
		return
	}

	event.Participants = append(event.Participants, participate.UserID)
	json.NewDecoder(r.Body).Decode(&event)

	database.Instance.Save(&event)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func checkIfEventExists(eventId string) bool {
	var event models.Event
	database.Instance.First(&event, eventId)
	if event.ID == 0 {
		return false
	}
	return true
}
