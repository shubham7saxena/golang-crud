package handler

import (
	"crud/contract"
	domain "crud/domain"
	"crud/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateNewUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest contract.User
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("error : %s", err.Error()), http.StatusBadRequest)
		return
	}

	user := domain.NewUser(&userRequest)

	err = service.InsertUserData(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
