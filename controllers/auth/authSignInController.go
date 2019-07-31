package authcontrollers

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"vlee/handles"
	"vlee/models"
)
// Define response
func SignIn(w http.ResponseWriter, r *http.Request) {
	// Define response
	var res	handles.ResponseResult
	// Take the context from middleware
	user := r.Context().Value("signInInfo").(*models.User)
	passwordInput := user.Password
	// Get mongodb collection
	Users := models.UsersCollection()
	err := Users.FindOne(context.TODO(), bson.D{{"email", user.Email}}).Decode(&user)
	if err != nil {
		res.Message = err.Error()
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordInput))
	if err != nil {
		res.Message = err.Error()
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":  user.ID, // Like this 5d3c263b9359848037ff3787
	})
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
