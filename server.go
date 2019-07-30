package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
	"vlee/handles"
	"vlee/routers"
)

func middlewareForAll(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" && r.Header.Get("Content-Type") != "application/json" {
			// Define response
			var res	handles.ResponseResult
			res.Message = "We only accept Content-Type as application/json"
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				panic(err)
			}
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
