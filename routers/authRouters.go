package routers

import (
	"net/http"
	c "vlee/controllers"
	m "vlee/middlewares"
)

func authRouters(prefix string, r *http.ServeMux) {
	r.Handle(prefix + "/signIn", m.ValidateSignIn(http.HandlerFunc(c.AuthSignIn)))
	r.Handle(prefix + "/signUp", m.ValidateSignUp(http.HandlerFunc(c.AuthSignUp)))
	r.Handle(prefix + "/profile", http.HandlerFunc(c.AuthProfile))
}
