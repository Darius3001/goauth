package main

import (
	"fmt"
	"net/http"
	"database/sql"
)

func main() {

	err := executeQuery(func(db *sql.DB) error {
		
		rows, err := db.Query("SHOW TABLES")

		if err != nil {
			return err
		}			
		defer rows.Close()

		fmt.Println("Tables in Database")
		for rows.Next() {
			var tableName string
			if err := rows.Scan(&tableName); err != nil {
				return err
			}
			fmt.Println(tableName)
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {

		responseString := "HIIIIII"

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		_, err := w.Write([]byte(responseString))

		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	})

	http.Handle("/auth", createAuthRoute())	

	http.ListenAndServe(":8080", nil)
}
