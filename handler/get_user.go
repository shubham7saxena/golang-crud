package handler

import (
	domain "crud/domain"
	"crud/repository"
	"crud/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	user, err := getUser(params["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

func getUser(id string) (*domain.User, error) {
	key := "user_" + id
	rp := utils.GetNewRedisPool()
	userProfile := domain.User{}
	keyExists, err := rp.Exists(key)

	if err != nil {
		fmt.Println("error checking key in redis")
	}

	if keyExists == true {
		userSerialized, err, check := rp.Get(key)
		if err != nil {
			fmt.Println("error retreiving from redis")
		}

		if check == true {
			err = json.Unmarshal(userSerialized, &userProfile)
			if err != nil {
				return &userProfile, err
			}
			return &userProfile, nil
		}
	}

	userID, err := strconv.Atoi(id)

	if err != nil {
		fmt.Errorf("Error: %s", err)
		return &userProfile, err
	}

	user, err := repository.NewUserRepository().GetUser(userID)

	if err != nil {
		fmt.Errorf("Error: %s", err)
		return &userProfile, err
	}

	response, err := json.Marshal(user)

	if err != nil {
		return &userProfile, err
	}

	err = rp.Set(key, response)

	if err != nil {
		fmt.Errorf("error setting user key in redis")
	}

	return user, nil
}
