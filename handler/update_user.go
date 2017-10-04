package handler

import (
	"crud/contract"
	domain "crud/domain"
	"crud/repository"
	"crud/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	existingUser, err := repository.NewUserRepository().GetUser(userID)

	if *existingUser == (domain.User{}) {
		http.Error(w, fmt.Sprintf("The user with the given ID does not exist"), http.StatusBadRequest)
		return
	}

	var userRequest contract.User
	err = json.NewDecoder(r.Body).Decode(&userRequest)

	if err != nil {
		http.Error(w, fmt.Sprintf("error : %s", err.Error()), http.StatusBadRequest)
		return
	}

	if userRequest.IsInvalid() {
		http.Error(w, "User's name and age are mandatory. Age cannot be negative", http.StatusBadRequest)
		return
	}

	updatedUser := domain.NewUser(&userRequest)

	err = repository.NewUserRepository().UpdateUser(updatedUser)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err.Error()), http.StatusInternalServerError)
	}

	userRedisKey := "user_" + params["id"]
	fmt.Println(userRedisKey)
	rp := utils.GetNewRedisPool()

	keyExists, err := rp.Exists(userRedisKey)

	if keyExists == true {
		response, err := json.Marshal(updatedUser)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshalling user"), http.StatusInternalServerError)
			return
		}

		err = rp.Set(userRedisKey, response)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error setting the updated user to redis"), http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusOK)
	return
}
