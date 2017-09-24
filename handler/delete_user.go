package handler

import "net/http"

const (
	deleteUserQuery = "DELETE from users where id='%d'"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}
