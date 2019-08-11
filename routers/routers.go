package routers

import (
	"net/http"
)
/*
*	The main router
*	Author: Lee Tuan
 */
// Go function
func Go(r *http.ServeMux) {
	authRouters("/auth", r)
}
