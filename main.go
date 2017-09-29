package main

import (
	app "crud/appcontext"
	"crud/server"
	"log"
	"net/http"

	sql "database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	app.InitDB()

	router := server.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
