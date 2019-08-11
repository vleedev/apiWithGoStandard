package handles

import (
	"encoding/json"
	"log"
	"net/http"
	"vlee/models"
)

/*
*	The handle functions
*	Author: Lee Tuan
 */
type ProfileMessage struct {
	MyProfile			*models.User	`json:"myProfile,omitempty"`
}
func (m *ProfileMessage) Response(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&m)
	if err != nil {
		log.Println(err)
	}
}