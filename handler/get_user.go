package handler

import (
	app "crud/appcontext"
	"fmt"
	"net/http"
)

const (
	readOneQuery = "SELECT from users WHERE id='%d'"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	data, err := db.Query(readOneQuery)

	if err != nil {
		fmt.Println("error in reading the user from database")
	}

	var age int
	var name string
	data.Scan(&age, &name)
	fmt.Printf("age: %d, name: %s \n", age, name)
}
