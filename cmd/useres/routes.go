package useres

import (
	"e-com/cmd/db"
	"e-com/services/model"
	"e-com/services/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

// func NewHandler() *Handler {
// 	return &Handler{}
// }
// func (h *Handler) RegisterRoutes(router *mux.Router) {
// 	router.HandleFunc("/login", h.handlelogin).Methods("POST")
// 	router.HandleFunc("/login", h.handleRegister).Methods("POST")

// }

func EcomRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", handlelogin).Methods("POST")
	r.HandleFunc("/register", handleRegister).Methods("POST")
	return r
}
func handlelogin(w http.ResponseWriter, r *http.Request) {

}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	db, _ := db.DbConnection()
	var payload model.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	json.NewDecoder(r.Body).Decode(&payload)
	userId, err := utils.InsertUserInDb(db, w, payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusOK, userId)
}
