package repository

import (
	"crud/appcontext"
	"crud/domain"
	"database/sql"
	"fmt"
)

const (
	readOneQuery = "SELECT * from users WHERE id='%d'"
	insertQuery  = "INSERT INTO users (age, name) VALUES($1, $2)"
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
