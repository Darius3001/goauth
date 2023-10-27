package userauth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
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

	subrouter := router.NewRoute().Subrouter()

	subrouter.Use(JWTMiddleware)

	subrouter.
		HandleFunc("/testjwt", jwtTestRoute).
		Methods("GET")

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

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "No JWT in request", http.StatusUnauthorized)
			return
		}

		extracedToken := GetOrNil(strings.Split(token, " "), 1)

		if extracedToken == nil {
			http.Error(w, "Authorization Header has wrong format", http.StatusBadRequest)
			return
		}

		userId, err := GetUserIdAndValidateToken(*extracedToken)

		if err != nil {
			http.Error(w, "JWT not valid", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func jwtTestRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	println(r.Context().Value("userId"))
}

func GetOrNil(arr []string, index int) *string {
	if index >= 0 && index < len(arr) {
		element := arr[index]
		return &element
	}
	return nil
}
