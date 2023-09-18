package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func executeQuery(queryFunc func(*sql.DB) error) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "goserver", "thisisunsafe", "localhost", "3306", "openpager"))

	if err != nil {
		return err
	}

	defer db.Close()

	return queryFunc(db)
}
