package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUSer(User) error
}

// type mockUserStore struct {
// }

func GetUserByEmail(email string) (*User, error) {
	return nil, nil
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	EmailId   string    `json:"email_id"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	EmailId   string `json:"email_id"`
	Password  string `json:"password"`
}

// Credentials represents the user credentials for login
type Credentials struct {
	EmailID  string `json:"email_id"`
	Password string `json:"password"`
}

// Claims represents the JWT claims
type Claims struct {
	GmailID string `json:"gmail_id"`
	jwt.StandardClaims
}
