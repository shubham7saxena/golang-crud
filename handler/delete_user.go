package handler

import (
	app "crud/appcontext"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	deleteUserQuery = "DELETE from users where id='%d'"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db := app.GetDB()
	params := r.URL.Query()
	userIDStr := params.Get("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(fmt.Sprintf(deleteUserQuery, userID))

	if err != nil {
		log.Fatal(err)
	}
}
