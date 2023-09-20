package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	dbUserName = "goserver"
	dbPassword = "thisisunsafe"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "openpager"
)

var dbUrl = fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s", 
	dbUserName, 
	dbPassword, 
	dbHost, 
	dbPort,
	dbName)

func ExecuteQuery(queryFunc func(*sql.DB) error) error {
	db, err := sql.Open("mysql", dbUrl) 

	if err != nil {
		return err
	}

	defer db.Close()

	return queryFunc(db)
}

func SetupDatabase() error {
	db, err := sql.Open("mysql", dbUrl) 

	if err != nil {
		log.Fatal("Could not open mysql database\n", err)	
		return err
	}
	defer db.Close()

	err = createTablesIfNotExists(db)


	return nil
}

func createTablesIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(30) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password_hash VARCHAR(60) NOT NULL,
			CHECK (
				CHAR_LENGTH(username) >= 3 
				AND email REGEXP '^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$'
			)
		)
	`)
	
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}