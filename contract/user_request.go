package contract

type User struct {
	Id   int    `json:"id,omitempty"`
	Age  int    `json:"age,omitempty"`
	Name string `json:"name,omitempty"`
}

func (u *User) IsInvalid() bool {
	return u.Age < 0
}
