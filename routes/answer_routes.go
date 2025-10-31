package routes

import (
	"mcqgame/handler"

	"github.com/gorilla/mux"
)

func AnswerRoutesControl(router *mux.Router) {
	router.HandleFunc("/api/submitanswer", handler.SubmitAnswerHandler).Methods("POST")
	router.HandleFunc("/api/player/{player_id}/room/{room_id}/answers", handler.GetAnswersByPlayerHandler).Methods("GET")

}
