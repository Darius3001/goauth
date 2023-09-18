package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {

		responseString := "HIIIIII"

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		_, err := w.Write([]byte(responseString))

		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
