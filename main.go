package main

import (
	app "crud/appcontext"
	"crud/config"
	"crud/server"
	"log"
	"net/http"
)

func main() {
	config.Load()
	app.Initiate()

	router := server.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}
