package validator

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"log"
	"net/http"
	"vlee/handles"
)
/*
*	The library for validating middleware
*	Author: vlee.dev
 */
var res	handles.ResponseResult
func checkEmail(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		res.Message = "The email field is empty"
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return false
	}
	if ! govalidator.IsEmail(*s) {
		res.Message = "The email is not valid"
		err := json.NewEncoder(w).Encode(&res)
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
		err := json.NewEncoder(w).Encode(&res)
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
		err := json.NewEncoder(w).Encode(&res)
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
		err := json.NewEncoder(w).Encode(&res)
		if err != nil {
			log.Println(err)
		}
		return false
	}
	return true
}