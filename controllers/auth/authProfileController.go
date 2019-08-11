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
	var gRes	handles.GeneralMessage
	var pRes	handles.ProfileMessage
	var tokenString string
	// Check the header has the field "Authorization"
	if values, ok := r.Header["Authorization"]; ok {
		// Check in slice
		for _, val := range values {
			if strings.Contains(val,"Bearer ") {
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
		pRes.MyProfile = user
		pRes.Response(w)
		return
	} else {
		gRes.Message = "Can not claim the information"
		gRes.Response(w)
		return
	}
}
