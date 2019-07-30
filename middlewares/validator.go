package middlewares

import (
	"context"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"net/http"
	"vlee/handles"
	"vlee/models"
)
var res	handles.ResponseResult
func checkEmail(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		res.Message = "The email field is empty"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err)
		}
		return false
	}
	if ! govalidator.IsEmail(*s) {
		res.Message = "The email is not valid"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err)
		}
		return false
	}
	return true
}
func checkPassword(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		res.Message = "The password field is empty"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err)
		}
		return false
	}
	return true
}
func checkFirstName(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		res.Message = "The first name field is empty"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err)
		}
		return false
	}
	return true
}
func checkLastName(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		res.Message = "The last name field is empty"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			panic(err)
		}
		return false
	}
	return true
}
func ValidateSignIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Define response
			var i models.User
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&i)
			if err != nil {
				res.Message = err.Error()
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					panic(err)
				}
				return
			}
			// Check email
			if ! checkEmail(w, &i.Email) {
				return
			}
			// Check password
			if ! checkPassword(w, &i.Password) {
				return
			}
			// Normalize email
			emailAddress, errr := govalidator.NormalizeEmail(i.Email)
			if errr != nil {
				res.Message = errr.Error()
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					panic(err)
				}
				return
			}
			i.Email = emailAddress
			// Pass these to context with value
			ctx := context.WithValue(r.Context(), "signInInfo", &i)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r) // Next
	})
}
func ValidateSignUp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Define response
			var i models.User
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&i)
			if err != nil {
				res.Message = err.Error()
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					panic(err)
				}
				return
			}
			// Check first name
			if ! checkFirstName(w, &i.FirstName) {
				return
			}
			// Check last name
			if ! checkLastName(w, &i.LastName) {
				return
			}
			// Check email
			if ! checkEmail(w, &i.Email) {
				return
			}
			// Check password
			if ! checkPassword(w, &i.Password) {
				return
			}
			// Normalize email
			emailAddress, err := govalidator.NormalizeEmail(i.Email)
			if err != nil {
				res.Message = err.Error()
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					panic(err)
				}
				return
			}
			i.Email = emailAddress

			ctx := context.WithValue(r.Context(), "signUpInfo", &i)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r) // Next
	})
}