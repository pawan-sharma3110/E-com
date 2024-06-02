package utils

import (
	"database/sql"
	"e-com/services/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	w.Header().Set("Content-Type", "application/json") // Correct the content-type typo
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
func InsertUserInDb(db *sql.DB, w http.ResponseWriter, payload model.RegisterUserPayload) (int, error) {
	query := `CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	email_id TEXT NOT NULL,
	password TEXT NOT NULL,
	created_at TIMESTAMP
)`
	_, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	var emailId string
	err = db.QueryRow(`SELECT email_id FROM users WHERE email_id=$1`, payload.EmailId).Scan(&emailId)

	// Handle the case where the email ID already exists
	if err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return 0, fmt.Errorf("user already exists")
	} else if err != sql.ErrNoRows {
		// Handle any other errors that are not 'no rows'
		return 0, err
	}

	// If we reach here, it means the email ID does not exist

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	var userid int
	query = `INSERT INTO users(first_name,last_name,email_id,password,created_at)VALUES($1,$2,$3,$4,$5)RETURNING id`
	err = db.QueryRow(query, payload.FirstName, payload.LastName, payload.EmailId, string(hashedPassword), time.Now()).Scan(&userid)
	if err != nil {
		return 0, err
	}

	return userid, nil
}

// func IsAlreadyReg(w http.ResponseWriter, payload model.RegisterUserPayload) (int, error) {
// 	db, _ := db.DbConnection()

// 	userId, err := InsertUserInDb(db, payload)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return userId, nil
// }
