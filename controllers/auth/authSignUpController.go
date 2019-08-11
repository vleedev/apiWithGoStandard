package authcontrollers

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"vlee/handles"
	"vlee/models"
	"vlee/repositories/repoimpls"
)
/*
*	The signing up controller
*	Author: Lee Tuan
 */
func SignUp(w http.ResponseWriter, r *http.Request) {
	// Define response
	var gRes	handles.GeneralMessage
	var suRes	handles.SignUpMessage
	// Take the context from middleware
	signUpInfo := r.Context().Value("signUpInfo").(*models.User)
	userRepo := repoimpls.NewUserRepo()
	status, err := userRepo.CheckUserExistence(&signUpInfo.Email)
	if err != nil {
		gRes.Message = err.Error()
		gRes.Response(w)
		return
	}
	if status == true {
		gRes.Message = "User exists"
		gRes.Response(w)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(signUpInfo.Password), 5)
	if err != nil {
		gRes.Message = "Error While Hashing Password, Try Again"
		gRes.Response(w)
		return
	}
	signUpInfo.Password = string(hash)
	err = userRepo.InsertOneUser(signUpInfo)
	if err != nil {
		gRes.Message = "Error While Creating User, Try Again"
		gRes.Response(w)
		return
	}
	suRes.Token = "Registration Successful"
	suRes.Response(w)
	return
}