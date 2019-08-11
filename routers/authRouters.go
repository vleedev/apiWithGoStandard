package routers

import (
	"net/http"
	auth "vlee/controllers/auth"
	"vlee/middlewares/validator"
)
/*
*	The main router for authentication
*	Author: Lee Tuan
 */
func authRouters(prefix string, r *http.ServeMux) {
	r.Handle(prefix + "/signIn", validator.SignIn(http.HandlerFunc(auth.SignIn)))
	r.Handle(prefix + "/signUp", validator.SignUp(http.HandlerFunc(auth.SignUp)))
	r.Handle(prefix + "/profile", http.HandlerFunc(auth.Profile))
}
