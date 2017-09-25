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
	deleteUserQuery = "DELETE from users where id='%d'"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf(deleteUserQuery, userID))

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}
