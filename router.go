package main

import (
	"database/sql"

	"github.com/ecojuntak/gorb/controllers"
	"github.com/ecojuntak/gorb/repositories"
	"github.com/gorilla/mux"
)

func LoadRouter(db *sql.DB) (r *mux.Router) {
	userRepo := repositories.NewUserRepo(db)
	userController := controllers.NewUserController(userRepo)

	r = mux.NewRouter()
	r.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.GetById).Methods("GET")
	r.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PATCH")
	r.HandleFunc("/users/{id}", userController.Delete).Methods("DELETE")

	return
}
