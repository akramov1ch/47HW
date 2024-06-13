package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"homework47/handler"
	"homework47/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	handler := handler.HandlerREPO(db)

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handler.GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/users", handler.GetAllUsersHandler).Methods("GET")

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
