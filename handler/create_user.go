package handler

import (
	app "crud/appcontext"
	"fmt"
	"net/http"
	"strconv"
)

const (
	insertQuery = "INSERT INTO users (age, name) VALUES($1, $2)"
)

func CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	params := r.URL.Query()
	name := params.Get("name")
	ageStr := params.Get("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(age)
	fmt.Println(name)
	_, err = db.Exec(insertQuery, age, name)
	if err != nil {
		fmt.Println("Error creating a new record")
	}
}
