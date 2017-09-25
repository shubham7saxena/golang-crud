package handler

import (
	app "crud/appcontext"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	readOneQuery = "SELECT age,name from users WHERE id='%d'"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	params := r.URL.Query()
	userIDStr := params.Get("id")
	userID, err := strconv.Atoi(userIDStr)
	data, err := db.Query(fmt.Sprintf(readOneQuery, userID))

	if err != nil {
		fmt.Println("error in reading the user from database")
	}

	var age int
	var name string

	for data.Next() {

		err = data.Scan(&age, &name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("age: %d, name: %s \n", age, name)
	}
}
