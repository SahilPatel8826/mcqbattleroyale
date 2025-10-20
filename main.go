package main

import (
	"fmt"
	"log"
	middleware "mcqgame/db"
	"mcqgame/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := middleware.CreateConnection()
	defer db.Close()
	r := mux.NewRouter()
	routes.RoutesControl(r)
	routes.QuestionRoutesControl(r)

	fmt.Println("Database connection established successfully")
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Server started at port 8000")
}
