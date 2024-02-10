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

	headers := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				return
			}
			h.ServeHTTP(w, r)
		})
	}

	http.ListenAndServe(":8080", headers(r))

}
