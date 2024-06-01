package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSon(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing reqest body")
	}
	return json.NewDecoder(r.Body).Decode(&payload)
}
func WriteJSon(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "aplication/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSon(w, status, map[string]string{"error": ""})
}
