package main

import (
	app "crud/appcontext"
	"crud/server"
	"log"
	"net/http"
)

func main() {
	app.Initiate()

	router := server.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
