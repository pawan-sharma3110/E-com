package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {
	connStr := `host=localhost port=5432 user= postgres password=Pawan@2003 dbname=E-com sslmode=disable`
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
