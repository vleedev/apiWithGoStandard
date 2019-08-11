package validator

import (
	"github.com/asaskevich/govalidator"
	"net/http"
	"vlee/handles"
)
/*
*	The library for validating middleware
*	Author: Lee Tuan
 */
var gRes	handles.GeneralMessage
func checkEmail(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		gRes.Message = "The email field is empty"
		gRes.Response(w)
		return false
	}
	if ! govalidator.IsEmail(*s) {
		gRes.Message = "The email is not valid"
		gRes.Response(w)
		return false
	}
	return true
}
func checkPassword(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		gRes.Message = "The password field is empty"
		gRes.Response(w)
		return false
	}
	return true
}
func checkFirstName(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		gRes.Message = "The first name field is empty"
		gRes.Response(w)
		return false
	}
	return true
}
func checkLastName(w http.ResponseWriter, s *string) bool {
	if *s == "" {
		gRes.Message = "The last name field is empty"
		gRes.Response(w)
		return false
	}
	return true
}