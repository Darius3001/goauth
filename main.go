package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"goauth.com/m/database"
	"goauth.com/m/userauth"
)

func main() {
	database.WaitForDatabase()

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

	http.HandleFunc("/testdb", database.TestDataBaseRoute)
	http.Handle("/", userauth.CreateAuthRoutes())
	http.ListenAndServe(":8080", nil)
}
