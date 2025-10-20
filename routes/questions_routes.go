package routes

import (
	"mcqgame/handler"

	"github.com/gorilla/mux"
)

func QuestionRoutesControl(router *mux.Router) {
	router.HandleFunc("/api/question", handler.CreateQuestionHandler).Methods("POST")
	router.HandleFunc("/api/question/{id}", handler.GetQuestionsByRoomIDHandler).Methods("GET")
	router.HandleFunc("/api/question/{id}", handler.UpdateQuestionByIDHandler).Methods("PUT")
}
