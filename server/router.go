package server

import (
	"crud/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() http.Handler {

	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
	router.HandleFunc("/user", handler.CreateNewUserHandler).Methods("POST")
	router.HandleFunc("/users", handler.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", handler.GetUserHandler).Methods("GET")
	router.HandleFunc("/user/{id}", handler.DeleteUserHandler).Methods("DELETE")

	return router
}
