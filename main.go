package main

import (
	app "crud/appcontext"
	"crud/handler"
	"fmt"
	"log"
	"net/http"

	sql "database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error

	db, err = app.InitDB()

	if err != nil {
		fmt.Println("couldn't connect to the database")
	}

	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/user", handler.CreateNewUserHandler).Methods("PUT")
	router.HandleFunc("/users", handler.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/user", handler.GetUserHandler).Methods("GET")
	router.HandleFunc("/user", handler.DeleteUserHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
