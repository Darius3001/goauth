package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

const (
	dbUserName = "goserver"
	dbPassword = "thisisunsafe"
	dbHost     = "goauth_mysql"
	dbPort     = "3306"
	dbName     = "goauth"
)

var dbUrl = fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s",
	dbUserName,
	dbPassword,
	dbHost,
	dbPort,
	dbName)

func WaitForDatabase() {
	maxTries := 15
	retryInterval := 3 * time.Second

	for try := 0; try < maxTries; try++ {

		db, err := sql.Open("mysql", dbUrl)

		if err != nil {
			log.Println("Could not open Database")
			time.Sleep(retryInterval)
			continue
		}

		if err := db.Ping(); err == nil {
			log.Println("Successfully pinged Database")
			return
		}
		log.Printf("Could not ping Database: Try %d\n", try)
		time.Sleep(retryInterval)
	}

	log.Fatalf("Database could not be reached after %d tries", maxTries)
}

func TestDataBaseRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TESTROUTE")
	_, err := sql.Open("mysql", dbUrl)
	if err != nil {
		http.Error(w, "Database not connected.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

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

	return err
}

func createTablesIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(30) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password_hash VARCHAR(60) NOT NULL,
			CONSTRAINT username_min_len_3 CHECK (
				CHAR_LENGTH(username) >= 3 
			),
			CONSTRAINT email_min_len CHECK (
				CHAR_LENGTH(email) >= 5
			),
			CONSTRAINT pw_hash_len_chk CHECK (
				CHAR_LENGTH(password_hash) = 60
			)	
		)
	`)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
