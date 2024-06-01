package useres

import (
	"e-com/services/model"
	"e-com/services/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store *model.UserStore
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
	r.HandleFunc("/login", handleRegister).Methods("POST")
	return r
}
func handlelogin(w http.ResponseWriter, r *http.Request) {

}
func handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload model.RegisterUserPayload
	if err := utils.ParseJSon(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
}
