package useres

import (
	"database/sql"
	"e-com/services/model"
	"fmt"
	"log"
)

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Db: db,
	}
}
func (s *Store) GetUserByEmail(email string) (*model.User, error) {
	rows, err := s.Db.Query(`SELECT * FROM users WHERE email=$1`, email)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	u := new(model.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}

	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}
func scanRowIntoUser(rows *sql.Rows) (*model.User, error) {
	user := new(model.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.EmailId,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
