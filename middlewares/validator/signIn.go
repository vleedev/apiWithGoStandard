package validator

import (
	"context"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"net/http"
	"vlee/models"
)
/*
*	The middleware for signing in
*	Author: Lee Tuan
 */
func SignIn(next http.Handler) http.Handler {
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
				gRes.Message = errr.Error()
				gRes.Response(w)
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