package validator

import (
	"context"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"log"
	"net/http"
	"vlee/models"
)

func SignIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Define response
			var user models.User
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&user)
			if err != nil {
				res.Message = err.Error()
				err := json.NewEncoder(w).Encode(&res)
				if err != nil {
					log.Println(err)
				}
				return
			}
			// Check email
			if ! checkEmail(w, &user.Email) {
				return
			}
			// Check password
			if ! checkPassword(w, &user.Password) {
				return
			}
			// Normalize email
			emailAddress, errr := govalidator.NormalizeEmail(user.Email)
			if errr != nil {
				res.Message = errr.Error()
				err := json.NewEncoder(w).Encode(&res)
				if err != nil {
					log.Println(err)
				}
				return
			}
			user.Email = emailAddress
			// Pass these to context with value
			ctx := context.WithValue(r.Context(), "signInInfo", &user)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r) // Next
	})
}