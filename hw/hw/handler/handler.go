package handler

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strconv"

	pb "homework47/user"

	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Handler struct {
	db *sql.DB
}

func HandlerREPO(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user pb.User
	readBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	protojson.Unmarshal(readBody, &user)

	_, err = proto.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := `INSERT INTO users (name, age, email, address, phone_numbers, occupation, company, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var id int
	err = h.db.QueryRow(query, user.Name, user.Age, user.Email, user.Address, user.PhoneNumbers, user.Occupation, user.Company, user.IsActive).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Id = int32(id)
	w.Header().Set("Content-Type", "application/x-protobuf")
	data, err := protojson.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

func (h *Handler) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user pb.User
	query := `SELECT id, name, age, email, address, phone_numbers, occupation, company, is_active FROM users WHERE id = $1`
	row := h.db.QueryRow(query, id)
	err = row.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Address, &user.PhoneNumbers, &user.Occupation, &user.Company, &user.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/x-protobuf")
	data, err := protojson.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

func (h *Handler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, age, email, address, phone_numbers, occupation, company, is_active FROM users`
	rows, err := h.db.Query(query)
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

	w.Header().Set("Content-Type", "application/x-protobuf")
	userMSG := &pb.Users{Users: users}
	jsonBytes, err := protojson.Marshal(userMSG)
	if err != nil{
		log.Fatal(err)
	}
	
	w.Write(jsonBytes)
}
