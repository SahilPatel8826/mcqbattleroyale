package handler

import (
	"encoding/json"
	"fmt"
	"mcqgame/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomCode := params["code"]

	var player models.Player
	json.NewDecoder(r.Body).Decode(&player)

	joined, err := models.JoinRoom(roomCode, player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(joined)
}
func GetPlayersByRoomHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomID, _ := strconv.Atoi(params["room_id"])

	players, err := models.GetPlayersByRoomID(roomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(players)
}
func DeletePlayerByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["player_id"])

	_, err := models.DeletePlayerByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Player with ID %d deleted successfully", id),
	})
}
