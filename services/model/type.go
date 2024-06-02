package model

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUSer(User) error
}
type mockUserStore struct {
}

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

// func handleRegister(w http.ResponseWriter, r *http.Request) {
// 	db, _ := db.DbConnection()
// 	var payload model.RegisterUserPayload
// 	json.NewDecoder(r.Body).Decode(&payload)

// 	id, err := utils.InsertUserInDb(db, payload)
// 	if err != nil {
// 		utils.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	utils.WriteJson(w, http.StatusOK, id)
// }
