package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/proto"


	pb "47HW/proto"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=vakhaboff dbname=shaxboz sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user pb.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := proto.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users (name, age, email, address, phone_numbers, occupation, company, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var id int
	err = db.QueryRow(query, user.Name, user.Age, user.Email, user.Address, user.PhoneNumbers, user.Occupation, user.Company, user.IsActive).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Id = int32(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user pb.User
	query := `SELECT id, name, age, email, address, phone_numbers, occupation, company, is_active FROM users WHERE id = $1`
	row := db.QueryRow(query, id)
	err = row.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Address, &user.PhoneNumbers, &user.Occupation, &user.Company, &user.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, age, email, address, phone_numbers, occupation, company, is_active FROM users`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []*pb.User
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Address, &user.PhoneNumbers, &user.Occupation, &user.Company, &user.IsActive)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, &user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", getUserByIDHandler).Methods("GET")
	r.HandleFunc("/users", getAllUsersHandler).Methods("GET")

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
