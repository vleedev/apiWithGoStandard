package validator

import (
	"context"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"net/http"
	"vlee/models"
)
/*
*	The middleware for signing up
*	Author: Lee Tuan
 */
func SignUp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Define response
			var user models.User
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&user)
			if err != nil {
				gRes.Message = err.Error()
				gRes.Response(w)
				return
			}
			// Check first name
			if ! checkFirstName(w, &user.FirstName) {
				return
			}
			// Check last name
			if ! checkLastName(w, &user.LastName) {
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
			emailAddress, err := govalidator.NormalizeEmail(user.Email)
			if err != nil {
				gRes.Message = err.Error()
				gRes.Response(w)
				return
			}
			user.Email = emailAddress

			ctx := context.WithValue(r.Context(), "signUpInfo", &user)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r) // Next
	})
}
