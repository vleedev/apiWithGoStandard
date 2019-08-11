package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
	"vlee/databases"
	"vlee/handles"
	"vlee/routers"
)
/*
*	The main package
*	Author: Lee Tuan
 */
func middlewareForAll(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" && r.Header.Get("Content-Type") != "application/json" {
			// Define response
			var gRes	handles.GeneralMessage
			gRes.Message = "We only accept Content-Type as application/json"
			gRes.Response(w)
			return
		}
		// Do stuff here
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r) //http.Handler
	})
}
func main() {
	/*
	*	Load environments
	 */
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
	// Create server mux
	r := http.NewServeMux()
	// Define routers
	routers.Go(r)
	// Connect MongoDB and create an instance
	databases.ConnectMongoDB()
	// Listen and serve
	s := &http.Server{
		Addr:           ":" + os.Getenv("APP_PORT"),
		Handler:        middlewareForAll(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("The app is running")
	log.Fatal(s.ListenAndServe())
}
