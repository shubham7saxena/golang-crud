package handler

import (
	"crud/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusBadRequest)
	}

	err = service.DeleteUserData(userID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
