package main

import (
	"e-com/cmd/db"
	"e-com/cmd/useres"

	"log"
	"net/http"
)

func main() {
	db, err := db.DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	server := useres.EcomRouter()
	http.ListenAndServe(":8080", server)
	// server := api.NewAPIServer(":8080", db)
	// if err := server.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
