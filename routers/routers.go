package routers

import (
	"net/http"
)

// Go function
func Go(r *http.ServeMux) {
	authRouters("/auth", r)
}
