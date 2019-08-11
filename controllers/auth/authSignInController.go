package authcontrollers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"vlee/handles"
	"vlee/models"
	"vlee/repositories/repoimpls"
)
/*
*	The signing in controller
*	Author: vlee.dev
 */
// Define response
func SignIn(w http.ResponseWriter, r *http.Request) {
	// Define response
	var res	handles.ResponseResult
	// Take the context from middleware
	signInInfo := r.Context().Value("signInInfo").(*models.User)
	// Work with mongodb via repository
	userRepo := repoimpls.NewUserRepo()
	user, err := userRepo.CheckSignInInfo(&signInInfo.Email, &signInInfo.Password)
	if err != nil {
		res.Message = err.Error()
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	// Prepare the information to generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":  user.ID, // Like this 5d3c263b9359848037ff3787
	})
	// Generate the token
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		res.Message = err.Error()
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	res.SignInToken = tokenString
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Println(err)
	}
}
