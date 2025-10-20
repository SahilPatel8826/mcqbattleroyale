package handler

import (
	"encoding/json"
	"fmt"
	"mcqgame/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var question models.Question

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("invalid request body: %v", err),
		})
		return
	}

	// Call model function to create room + questions
	createdQuestion, err := models.CreateQuestion(question)
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
	json.NewEncoder(w).Encode(createdQuestion)
}

func GetQuestionsByRoomIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	created, err := models.GetQuestionsByRoomID(id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(created)
}
func UpdateQuestionByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}
	var question models.Question
	json.NewDecoder(r.Body).Decode(&question)
	updated, err := models.UpdateQuestionByID(id, question)
	if err != nil {
		panic(err)

	}

	json.NewEncoder(w).Encode(updated)

}
