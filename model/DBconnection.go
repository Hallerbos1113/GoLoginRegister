package model

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host			= "localhost"
	port			= 5432
	user			= "postgres"
	password	=	"1234"
	dbname		= "testdb"
)

var DB *sql.DB

func DBConnect() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	
		// Open the database connection
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
			log.Fatal(err)
	}
	// defer DB.Close()

	// Check the connection
	err = DB.Ping()
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

}