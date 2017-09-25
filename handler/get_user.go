package handler

import (
	app "crud/appcontext"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	readOneQuery = "SELECT * from users WHERE id='%d'"
)

type User struct {
	Id   int    `json:"id,omitempty"`
	Age  int    `json:"age,omitempty"`
	Name string `json:"name,omitempty"`
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	data := db.QueryRow(fmt.Sprintf(readOneQuery, userID))

	userData := User{}

	err = data.Scan(&userData.Age, &userData.Name, &userData.Id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	response, err := json.Marshal(userData)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}
