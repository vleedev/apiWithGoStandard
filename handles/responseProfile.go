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
type ProfileMessage struct {
	Email           string  `json:"email,omitempty"`
	FirstName       string  `json:"firstName,omitempty"`
	LastName        string  `json:"lastName,omitempty"`
	Avatar          string  `json:"avatar,omitempty"`
	Telephone       uint8   `json:"telephone,omitempty"`
	Birthday        int64   `json:"birthday,omitempty"`
	Facebook        string  `json:"facebook,omitempty"`
	Address         string  `json:"address,omitempty"`
	LocationLat     float64 `json:"locationLat,omitempty"`
	LocationLon     float64 `json:"locationLon,omitempty"`
	ResetToken      string  `json:"resetToken,omitempty"`
	Language        string  `json:"language,omitempty"`
	VerifiedByEmail bool    `json:"verifiedByEmail" default:"false"`
	VerifiedByPhone bool    `json:"verifiedByPhone" default:"false"`
	CreatedAt       int64   `json:"createdAt,omitempty"`
	UpdatedAt       int64   `json:"updatedAt,omitempty"`
}

func (m *ProfileMessage) Response(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&m)
	if err != nil {
		log.Println(err)
	}
}
