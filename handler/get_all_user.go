package handler

import (
	app "crud/appcontext"
	"fmt"
	"log"
	"net/http"
)

const (
	readAllQuery = "SELECT age, name from users limit 10"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	data, err := db.Query(readAllQuery)

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()
	var age int
	var name string

	for data.Next() {
		err := data.Scan(&age, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("age: %d, name: %s \n", age, name)
	}
}
