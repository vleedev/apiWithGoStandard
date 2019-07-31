package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"os"
	"strings"
	"vlee/handles"
	"vlee/models"
)
func AuthProfile(w http.ResponseWriter, r *http.Request) {
	// Define response
	var res	handles.ResponseResult
	tokenString := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		res.Message = err.Error()
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	var result models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Take id string
		ID := claims["_id"].(string)
		// Model Users
		Users := models.UsersCollection()
		// Convert ID string to ID Object
		objID, err := primitive.ObjectIDFromHex(ID)
		if err != nil {
			res.Message = err.Error()
			err := json.NewEncoder(w).Encode(&res)
			if err != nil {
				log.Println(err)
			}
			return
		}
		// Use objID to search in database
		err = Users.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&result)
		if err != nil {
			res.Message = err.Error()
			err := json.NewEncoder(w).Encode(&res)
			if err != nil {
				log.Println(err)
			}
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Println(err)
		}
		return
	} else {
		res.Message = "Can not claim the information"
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
}
