package utils

import (
	"database/sql"
	"e-com/cmd/db"
	"e-com/services/model"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	defer r.Body.Close() // Ensure the body is closed after parsing

	return json.NewDecoder(r.Body).Decode(payload)
}
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("content-type", "aplication/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
func insertUserInDb(db *sql.DB,payload model.RegisterUserPayload) (int, error) {
	query := `CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT NOT NULL,
	created_at TIMESTAMP
)`
	_, err := db.Exec(query)
	if err != nil {
		return 0, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	var userid int
	query = `INSERT INTO users(first_name,last_name,email,password,created_at)VALUES($1,$2,$3,$4,$5,$6)RETURNING id`
	err = db.QueryRow(query, payload.FirstName, payload.LastName, payload.Email, string(hashedPassword)).Scan(&userid)
	if err != nil {
		return 0, err
	}

	return userid, nil
}

func IsAlreadyReg(w http.ResponseWriter, payload model.RegisterUserPayload) (int, error) {
	var exists bool
	db, err := db.DbConnection()
	if err != nil {
		return 0, err
	}
	err = db.QueryRow(`SELECT email_id FROM users WHERE email=$1`, payload.Email).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if exists {
		http.Error(w, "User already exists", http.StatusConflict)
	}
	userId, err := insertUserInDb(db, payload)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
