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
	database.Instance.Find(&events)
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
	eventId := strconv.Itoa(participate.EventID)
	if checkIfEventExists(eventId) == false {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var event models.Event
	database.Instance.First(&event, eventId)

	var user models.User
	userId := participate.UserID

	if checkIfUserExists(strconv.Itoa(userId)) == false {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}

	database.Instance.First(&user, userId)
	user.Bonus += 2 // participation bonus

	event.Participants = append(event.Participants, user)
	if event.Limit == len(event.Participants) {
		event.Status = "in progress"
	}
	//event.RahmetParticipants = append(event.RahmetParticipants, int32(user.ID))
	json.NewDecoder(r.Body).Decode(&event)

	database.Instance.Save(&event)
	database.Instance.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func RahmetEvent(w http.ResponseWriter, r *http.Request) {

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

	event.Bonus += 5

	if len(event.RahmetParticipants) > 0 {
		for i, _ := range event.RahmetParticipants {
			if event.RahmetParticipants[i] == int32(participate.UserID) {
				w.WriteHeader(http.StatusBadGateway)
				json.NewEncoder(w).Encode("Cannot rahmet more than once!")
				return
			}
		}
	}
	event.RahmetParticipants = append(event.RahmetParticipants, int32(participate.UserID))

	if len(event.Participants)-len(event.RahmetParticipants) == 0 {
		event.Status = "inactive"
	}

	json.NewDecoder(r.Body).Decode(&event)

	var creator models.User
	userId := event.CreatorID

	if checkIfUserExists(strconv.Itoa(int(userId))) == false {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}

	database.Instance.First(&creator, userId)

	creator.Bonus += 5

	database.Instance.Save(&event)
	database.Instance.Save(&creator)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func EndEventByCreator(w http.ResponseWriter, r *http.Request) {

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
	event.Status = "finished"

	database.Instance.Save(&event)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func StartEventByCreator(w http.ResponseWriter, r *http.Request) {

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
	event.Status = "in progress"

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
