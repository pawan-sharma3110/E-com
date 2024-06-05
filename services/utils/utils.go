package utils

import (
	"database/sql"
	"e-com/services/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
func InsertUserInDb(db *sql.DB, payload model.RegisterUserPayload) (int, error) {
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

	hashedPassword, err := gernateHashedPassword(payload.Password)
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

func IsAlreadyReg(db *sql.DB, payload model.RegisterUserPayload) (int, error) {
	var emailId string
	err := db.QueryRow(`SELECT email_id FROM users WHERE email_id=$1`, payload.EmailId).Scan(&emailId)

	if err == nil {
		return 0, fmt.Errorf("user already exists")
	} else if err != sql.ErrNoRows {
		return 0, err
	}

	userId, err := InsertUserInDb(db, payload)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
func gernateHashedPassword(password string) (pass string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func comparePasswords(hashedPassword string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), plainPassword)
	return err == nil
}

func UserDetailsInDb(w http.ResponseWriter, db *sql.DB, payload model.Credentials) {
	var userCredentials model.Credentials
	jwtSecret := "secret_key"
	query := `SELECT email_id,password FROM users WHERE email_id=$1`
	err := db.QueryRow(query, payload.EmailID).Scan(&userCredentials.EmailID, &userCredentials.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Unkwon Users", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	validPAssword := comparePasswords(userCredentials.Password, []byte(payload.Password))
	if !validPAssword {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Wrong password"})
		return
	}
	println(validPAssword)
	// Generate a new JWT for the user
	expirationTime := time.Now().Add(5 * time.Minute) // Token valid for 5 minutes
	claims := &model.Claims{
		GmailID: payload.EmailID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Set token expiration time
		},
	}

	// Create the JWT using the specified signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		fmt.Println("3")
		return
	}

	// Return the generated JWT to the client
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}
