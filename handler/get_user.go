package handler

import (
	app "crud/appcontext"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	readOneQuery = "SELECT age,name from users WHERE id='%d'"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	data, err := db.Query(fmt.Sprintf(readOneQuery, userID))

	if err != nil {
		log.Fatal(err)
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
