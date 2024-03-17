package main

import (
	"net/http"

	"restaurant-api/handlers"

	"github.com/gorilla/mux"
)

type Response struct {
	Users []User `json:"users"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")

	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":9000", router)
}
