package routes

import (
	"mcqgame/handler"

	"github.com/gorilla/mux"
)

func PlayerRoutesControl(router *mux.Router) {
	router.HandleFunc("/api/joinroom/{code}", handler.JoinRoomHandler).Methods("POST")
	router.HandleFunc("/api/allplayers/{room_id}", handler.GetPlayersByRoomHandler).Methods("GET")
	router.HandleFunc("/api/players/{player_id}", handler.DeletePlayerByIDHandler).Methods("DELETE")
}
