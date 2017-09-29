package domain

import "crud/contract"

type User struct {
	Id   int    `json:"id,omitempty"`
	Age  int    `json:"age,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewUser(s *contract.User) *User {
	return &User{
		Id:   s.Id,
		Age:  s.Age,
		Name: s.Name,
	}
}
