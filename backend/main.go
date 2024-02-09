// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rihabcheberli/App-NextJS-Go/backend/database"
	"github.com/rihabcheberli/App-NextJS-Go/backend/handlers"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users", handlers.GetAllUsersHandler).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
