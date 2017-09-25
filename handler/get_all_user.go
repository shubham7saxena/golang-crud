package handler

import (
	app "crud/appcontext"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	readAllQuery = "SELECT age, name, id from users limit 10"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	data, err := db.Query(readAllQuery)

	if err != nil {
		log.Fatal(err)
	}

	userData := User{}
	allUserData := make([]User, 0, 10)
	for i := 0; data.Next(); i++ {
		err := data.Scan(&userData.Age, &userData.Name, &userData.Id)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		allUserData = append(allUserData, userData)
	}

	response, err := json.Marshal(allUserData)

	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
