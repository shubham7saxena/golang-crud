package main

import (
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
	insertQuery  = "INSERT INTO users (age, name) VALUES($1, $2)"
	readAllQuery = "SELECT * from users limit 10"
	readOneQuery = "SELECT from users WHERE name='%s' AND age='%d'"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/user", createNewUserHandler).Methods("PUT")
	router.HandleFunc("/users", getAllUsersHandler).Methods("GET")
	router.HandleFunc("/user/id", getUserHandler).Methods("GET")
	router.HandleFunc("/user/id", deleteUserHandler).Methods("DELETE")
	//		db connections

	var err error

	db, err = sql.Open("postgres", dbConnectionString())

	if err != nil {
		fmt.Println("couldn't connect to the database")
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createNewUserHandler(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec(insertQuery, 20, "lovee")
	if err != nil {
		fmt.Println("Error creating a new record")
	}
}

func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	data, err := db.Query(readAllQuery)

	if err != nil {
		fmt.Println("error in reading all the queries from database")
	}

	var age int
	var name string

	for data.Next() {
		data.Scan(&age, &name)
		fmt.Printf("age: %d, name: %s \n", age, name)
	}

}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {

}

func dbConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' sslmode=disable", "test", "postgres", "s7saxena")
}
