package appcontext

import (
	sql "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", dbConnectionString())

	if err != nil {
		fmt.Println("Error connecting to the database")
	}

	if err = db.Ping(); err != nil {
		fmt.Printf("Ping to database host failed: %s \n", err)
	}
	return db, err
}

func dbConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' sslmode=disable", "test", "postgres", "s7saxena")
}

func GetDB() *sql.DB {
	return db
}
