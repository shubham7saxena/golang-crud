package handler

import (
	app "crud/appcontext"
	"fmt"
	"net/http"
)

const (
	readAllQuery = "SELECT * from users limit 10"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
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
