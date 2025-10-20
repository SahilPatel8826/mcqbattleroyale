package handler

import (
	"encoding/json"
	"fmt"
	"mcqgame/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateRoomHandler handles creating a new room with multiple questions
func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	var room models.Room

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("invalid request body: %v", err),
		})
		return
	}

	// Call model function to create room + questions
	createdRoom, err := models.CreateRoom(room)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("unable to create room: %v", err),
		})
		return
	}

	// Send the created room back as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdRoom)
}
func GetRoomHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	get, err := models.GetRoomWithQuestions(id)
	if err != nil {
		fmt.Println("GetRoom in model is not working", err)
	}

	json.NewEncoder(w).Encode(get)
}
