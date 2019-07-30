package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"vlee/handles"
	"vlee/models"
)

func AuthSignUp(w http.ResponseWriter, r *http.Request) {
	// Take the context from middleware
	user := r.Context().Value("signUpInfo").(*models.User)
	Users := models.UsersCollection()
	// Define response
	var res	handles.ResponseResult

	err := Users.FindOne(context.TODO(), bson.D{{"email", user.Email}}).Decode(&user)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
			if err != nil {
				res.Message = "Error While Hashing Password, Try Again"
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					panic(err)
				}
				return
			}
			user.Password = string(hash)

			_, err = Users.InsertOne(context.TODO(), user)
			if err != nil {
				res.Message = "Error While Creating User, Try Again"
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					panic(err)
				}
				return
			}
			res.Message = "Registration Successful"
			err = json.NewEncoder(w).Encode(res)
			if err != nil {
				panic(err)
			}
			return
		}
		res.Message = err.Error()
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err)
		}
		return
	}
	res.Message = "Email already Existed!!"
	w.WriteHeader(409)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}
	return
}