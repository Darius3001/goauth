package main

import (
	// "encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func createAuthRoute() *mux.Router {
	
	router := mux.NewRouter()

	router.HandleFunc("/login", func (w http.ResponseWriter, r *http.Request) {
		
	})

	router.HandleFunc("/register", func (w http.ResponseWriter, r *http.Request) {
		
	})

	return router

}