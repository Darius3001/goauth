package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"openpager.com/m/database"
)

func main() {
	dbSetupErr := database.SetupDatabase()
	if dbSetupErr != nil {
		return
	}



	err := database.ExecuteQuery(func(db *sql.DB) error {
		
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
