package handler

import (
	"crud/contract"
	domain "crud/domain"
	"crud/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	var userRequest contract.User
	err = json.NewDecoder(r.Body).Decode(&userRequest)

	if err != nil {
		http.Error(w, fmt.Sprintf("error : %s", err.Error()), http.StatusBadRequest)
	}

	if !userExists(userID) {
		http.Error(w, fmt.Sprintf("The user with the given ID does not exist"), http.StatusBadRequest)
	}

	updatedUser := domain.NewUser(&userRequest)

	err = service.UpdateUserData(updatedUser, userID)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func userExists(id int) bool {
	existingUser, err := service.GetUserData(id)
	if err != nil {
		fmt.Errorf("Error checking user in database")
	}

	if *existingUser == (domain.User{}) {
		return false
	}
	return true
}
