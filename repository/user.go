package repository

import (
	"crud/appcontext"
	"crud/domain"
	"database/sql"
	"fmt"
)

const (
	readOneQuery    = "SELECT * from users WHERE id='%d'"
	insertQuery     = "INSERT INTO users (age, name) VALUES($1, $2)"
	deleteUserQuery = "DELETE from users where id='%d'"
	updateUserQuery = "Update users set age='%d', name='%s' where id='%d'"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() *userRepository {
	return &userRepository{
		db: appcontext.GetDB(),
	}
}

func (ur *userRepository) Insert(u *domain.User) error {
	_, err := ur.db.Exec(insertQuery, u.Age, u.Name)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (ur *userRepository) GetUser(userID int) (*domain.User, error) {
	data := ur.db.QueryRow(fmt.Sprintf(readOneQuery, userID))
	userData := domain.User{}
	err := data.Scan(&userData.Age, &userData.Name, &userData.Id)

	if err != nil {
		return &domain.User{}, err
	}

	return &userData, nil
}

func (ur *userRepository) DeleteUser(userID int) error {
	_, err := ur.db.Exec(fmt.Sprintf(deleteUserQuery, userID))

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (ur *userRepository) UpdateUser(u *domain.User) error {
	_, err := ur.db.Exec(fmt.Sprintf(updateUserQuery, u.Age, u.Name, u.Id))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
