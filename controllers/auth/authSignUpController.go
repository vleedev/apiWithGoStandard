package authcontrollers

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
	"vlee/handles"
	"vlee/models"
	"vlee/repositories/repoimpls"
)
/*
*	The signing up controller
*	Author: vlee.dev
 */
func SignUp(w http.ResponseWriter, r *http.Request) {
	// Define response
	var res	handles.ResponseResult
	// Take the context from middleware
	signUpInfo := r.Context().Value("signUpInfo").(*models.User)
	userRepo := repoimpls.NewUserRepo()
	status, err := userRepo.CheckUserExistence(&signUpInfo.Email)
	if err != nil {
		res.Message = err.Error()
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	if status == true {
		res.Message = "User exists"
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(signUpInfo.Password), 5)
	if err != nil {
		res.Message = "Error While Hashing Password, Try Again"
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	signUpInfo.Password = string(hash)
	signUpInfo.CreatedAt = time.Now().UnixNano()
	signUpInfo.UpdatedAt = time.Now().UnixNano()
	err = userRepo.InsertOneUser(signUpInfo)
	if err != nil {
		res.Message = "Error While Creating User, Try Again"
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	res.Message = "Registration Successful"
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Println(err)
	}
	return
}