package authcontrollers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"strings"
	"vlee/handles"
	"vlee/repositories/repoimpls"
)

/*
*	The profile controller
*	Author: Lee Tuan
 */
func Profile(w http.ResponseWriter, r *http.Request) {
	// Define response
	var gRes handles.GeneralMessage
	var pRes handles.ProfileMessage
	var tokenString string
	// Check the header has the field "Authorization"
	if values, ok := r.Header["Authorization"]; ok {
		// Check in slice
		for _, val := range values {
			if strings.Contains(val, "Bearer ") {
				tokenString = strings.Split(val, "Bearer ")[1]
			}
		}
	}
	if tokenString == "" {
		gRes.Message = "Token is not found!"
		gRes.Response(w)
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		gRes.Message = err.Error()
		gRes.Response(w)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Take id string
		ID := claims["_id"].(string)
		// Model Users
		userRepo := repoimpls.NewUserRepo()
		// Convert ID string to ID Object
		objID, err := primitive.ObjectIDFromHex(ID)
		if err != nil {
			gRes.Message = err.Error()
			gRes.Response(w)
			return
		}
		// Use objID to search in database
		user, err := userRepo.GetUserByObjectID(&objID)
		if err != nil {
			gRes.Message = err.Error()
			gRes.Response(w)
			return
		}
		// Hide some fields before showing
		user.Password = ""
		// Response to the client
		pRes.Email = user.Email
		pRes.FirstName = user.FirstName
		pRes.LastName = user.LastName
		pRes.Address = user.Address
		pRes.Avatar = user.Avatar
		pRes.Birthday = user.Birthday
		pRes.Facebook = user.Facebook
		pRes.Language = user.Language
		pRes.LocationLat = user.LocationLat
		pRes.LocationLon = user.LocationLon
		pRes.Telephone = user.Telephone
		pRes.CreatedAt = user.CreatedAt
		pRes.UpdatedAt = user.UpdatedAt
		pRes.VerifiedByEmail = user.VerifiedByEmail
		pRes.VerifiedByPhone = user.VerifiedByPhone
		pRes.Response(w)
		return
	} else {
		gRes.Message = "Can not claim the information"
		gRes.Response(w)
		return
	}
}
