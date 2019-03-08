package config

import (
	"database/sql"
	"fmt"
	// Needed only for initialization.
	_ "github.com/lib/pq"
)

// Database is the object uses by the models for accessing
// database tables and executing queries.
var Database *sql.DB

func init() {
	var err error
	Database, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/currency?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = Database.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connection successful.")
}
