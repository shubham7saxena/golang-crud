package main

import (
	app "crud/appcontext"
	"crud/server"
	"fmt"
	"log"
	"net/http"

	sql "database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error

	db, err = app.InitDB()

	if err != nil {
		fmt.Println("couldn't connect to the database")
	}

	router := server.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
