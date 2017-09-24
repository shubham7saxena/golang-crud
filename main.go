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

type user struct {
	name string
	age  int
}

var db *sql.DB

const (
	readOneQuery = "SELECT from users WHERE name='%s' AND age='%d'"
)

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
	router.HandleFunc("/user/id", getUserHandler).Methods("GET")
	router.HandleFunc("/user/id", deleteUserHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {

}
