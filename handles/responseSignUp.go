package handles

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
*	The handle functions
*	Author: Lee Tuan
 */
type SignUpMessage struct {
	Token	string	`json:"token,omitempty"`
}
func (m *SignUpMessage) Response(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&m)
	if err != nil {
		log.Println(err)
	}
}