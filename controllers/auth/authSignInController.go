package authcontrollers

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"vlee/handles"
	"vlee/models"
	"vlee/repositories/repoimpls"
)
/*
*	The signing in controller
*	Author: Lee Tuan
 */
// Define response
func SignIn(w http.ResponseWriter, r *http.Request) {
	// Define response
	var gRes	handles.GeneralMessage
	var siRes	handles.SignInMessage
	// Take the context from middleware
	signInInfo := r.Context().Value("signInInfo").(*models.User)
	// Work with mongodb via repository
	userRepo := repoimpls.NewUserRepo()
	user, err := userRepo.CheckSignInInfo(&signInInfo.Email, &signInInfo.Password)
	if err != nil {
		gRes.OriginalMessage = err.Error()
		gRes.Response(w)
		return
	}
	// Prepare the information to generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":  user.ID, // Like this 5d3c263b9359848037ff3787
	})
	// Generate the token
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		gRes.OriginalMessage = err.Error()
		gRes.Response(w)
		return
	}
	siRes.Token = tokenString
	siRes.Response(w)
}
