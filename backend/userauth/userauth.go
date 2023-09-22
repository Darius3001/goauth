package userauth

import (
	"net/http"
	"github.com/gorilla/mux"
	"strings"
)

func CreateAuthRoutes() *mux.Router {
	
	router := mux.NewRouter()

	router.Use(jsonMiddleware)
	
	router.
		HandleFunc("/login", handleLogin).
		Methods("POST")

	router.
		HandleFunc("/register", handleRegistration).
		Methods("POST")

	return router
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	})
}