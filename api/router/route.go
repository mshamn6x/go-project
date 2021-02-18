package router

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	r.Use(MiddlewareHandler)
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	return r

}
