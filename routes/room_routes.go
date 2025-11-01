package routes

import (
	"mcqgame/handler"

	"github.com/gorilla/mux"
)

func RoutesControl(router *mux.Router) {
	router.HandleFunc("/api/getroom/{id}", handler.GetRoomHandler).Methods("GET")
	router.HandleFunc("/api/createroom", handler.CreateRoomHandler).Methods("POST")
	router.HandleFunc("/api/room/{id}/start", handler.StartRoomHandler).Methods("PATCH")
	router.HandleFunc("/api/room/{id}/end", handler.EndRoomHandler).Methods("PATCH")

}
