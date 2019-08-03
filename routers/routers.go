package routers

import (
	"net/http"
)
/*
*	The main router
*	Author: vlee.dev
 */
// Go function
func Go(r *http.ServeMux) {
	authRouters("/auth", r)
}
