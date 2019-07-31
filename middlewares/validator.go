package middlewares

import (
	"context"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"log"
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
			log.Println(err)
		}
		return false
	}
	if ! govalidator.IsEmail(*s) {
		res.Message = "The email is not valid"
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Println(err)
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
			log.Println(err)
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
			log.Println(err)
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
			log.Println(err)
		}
		return false
	}
	return true
}
func ValidateSignIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Define response
			var user models.User
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&user)
			if err != nil {
				res.Message = err.Error()
				err := json.NewEncoder(w).Encode(res)
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
				err := json.NewEncoder(w).Encode(res)
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
func ValidateSignUp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Define response
			var user models.User
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&user)
			if err != nil {
				res.Message = err.Error()
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					log.Println(err)
				}
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
				res.Message = err.Error()
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					log.Println(err)
				}
				return
			}
			user.Email = emailAddress

			ctx := context.WithValue(r.Context(), "signUpInfo", &user)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r) // Next
	})
}