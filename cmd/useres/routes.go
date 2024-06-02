package useres

import (
	"e-com/cmd/db"
	"e-com/services/model"
	"e-com/services/utils"
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
	db, err := db.DbConnection()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close() // Ensure the database connection is closed after use

	var payload model.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Check if user is already registered
	userId, err := utils.IsAlreadyReg(db, payload)
	if err != nil {
		if err.Error() == "user already exists" {
			utils.WriteError(w, http.StatusConflict, err)
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}

	utils.WriteJson(w, http.StatusOK, userId)
}
