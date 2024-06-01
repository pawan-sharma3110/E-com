package utils

import (
	"database/sql"
	"e-com/cmd/db"
	"e-com/services/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSon(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing reqest body")
	}
	return json.NewDecoder(r.Body).Decode(&payload)
}
func WriteJSon(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "aplication/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSon(w, status, map[string]string{"error": ""})
}
func insertUserInDb(db *sql.DB, w http.ResponseWriter, payload model.RegisterUserPayload) int {
	query := `CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT NOT NULL,
	created_at TIMESTAMP
)`
	_, err := db.Query(query)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	var userid int
	query = `INSERT INTO users(first_name,last_name,email,password,created_at)VALUES($1,$2,$3,$4,$5,$6)RETURNING id`
	err = db.QueryRow(query, payload.FirstName, payload.LastName, payload.Email, payload.Password).Scan(&userid)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	return userid
}
func IsAlreadyReg(w http.ResponseWriter, payload model.RegisterUserPayload) int {
	var exists bool
	db, err := db.DbConnection()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = db.QueryRow(`SELECT email_id FROM users WHERE email=$1`, payload.Email).Scan(&exists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if exists {
		http.Error(w, "User already exists", http.StatusConflict)
	}
	userId := insertUserInDb(db, w, payload)
	return userId
}
