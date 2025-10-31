package handler

import (
	"encoding/json"
	"fmt"
	"mcqgame/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SubmitAnswerHandler(w http.ResponseWriter, r *http.Request) {
	var answer models.Answer
	json.NewDecoder(r.Body).Decode(&answer)

	createdAnswer, err := models.SubmitAnswer(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdAnswer)
}

func GetAnswersByPlayerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// extract route params
	playerIDStr := params["player_id"]
	roomIDStr := params["room_id"]

	// convert string to int
	playerID, err := strconv.Atoi(playerIDStr)
	if err != nil {
		http.Error(w, "Invalid player_id", http.StatusBadRequest)
		return
	}

	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		http.Error(w, "Invalid room_id", http.StatusBadRequest)
		return
	}

	// call model function
	answers, err := models.GetAnswersByPlayerID(playerID, roomID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching answers: %v", err), http.StatusInternalServerError)
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answers)
}
