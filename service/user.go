package service

import (
	"crud/domain"
	"crud/repository"
	"encoding/json"
	"fmt"
	"strconv"
)

func GetUserData(userID int) (*domain.User, error) {
	key := getRedisKey(userID)
	rp := repository.GetNewRedisPool()
	userProfile := domain.User{}
	keyExists := rp.Exists(key)

	if keyExists == true {
		userSerialized, check := rp.Get(key)

		if check == true {
			err := json.Unmarshal(userSerialized, &userProfile)
			if err != nil {
				return &userProfile, err
			}
			return &userProfile, nil
		}
	}

	user, err := repository.NewUserRepository().GetUser(userID)

	if err != nil {
		fmt.Errorf("Error: %s", err.Error)
		return &userProfile, err
	}

	userRedis, err := marshalUserDataToBytes(user)

	if err != nil {
		fmt.Errorf("cannot convert user data to bytes")
	}

	rp.Set(getRedisKey(userID), userRedis)

	return user, nil
}

func DeleteUserData(userID int) error {
	rp := repository.GetNewRedisPool()
	err := rp.Delete(getRedisKey(userID))

	if err != nil {
		fmt.Errorf("cannot delete user from redis")
		return err
	}

	err = repository.NewUserRepository().DeleteUser(userID)

	if err != nil {
		fmt.Errorf("Error: %s", err.Error)
		return err
	}
	return nil
}

func InsertUserData(user *domain.User) error {
	err := repository.NewUserRepository().Insert(user)

	if err != nil {
		fmt.Errorf("Error: %s", err.Error)
		return err
	}
	return nil
}

func UpdateUserData(user *domain.User, userID int) error {
	rp := repository.GetNewRedisPool()
	err := repository.NewUserRepository().UpdateUser(user)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error)
		return err
	}

	userRedis, err := marshalUserDataToBytes(user)
	if err != nil {
		fmt.Errorf("cannot convert user data to bytes")
		return err
	}
	rp.Set(getRedisKey(userID), userRedis)
	return nil
}

func getRedisKey(id int) string {
	userIDstr := strconv.Itoa(id)
	userRedisKey := "user_" + userIDstr
	return userRedisKey
}

func marshalUserDataToBytes(user *domain.User) ([]byte, error) {
	userDatainBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error)
		return userDatainBytes, err
	}

	return userDatainBytes, nil

}
