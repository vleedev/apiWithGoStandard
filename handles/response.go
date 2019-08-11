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
type GeneralMessage struct {
	Message	string	`json:"message,omitempty"`
	OriginalMessage	string	`json:"originalMessage,omitempty"`
}
func (m *GeneralMessage) Response(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&m)
	if err != nil {
		log.Println(err)
	}
}