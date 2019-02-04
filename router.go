package main

import (
	"database/sql"

	"github.com/ecojuntak/gorb/controllers"
	"github.com/ecojuntak/gorb/middlewares"
	"github.com/ecojuntak/gorb/repositories"
	"github.com/gorilla/mux"
)

func LoadRouter(db *sql.DB) (r *mux.Router) {
	userRepo := repositories.NewUserRepo(db)
	userController := controllers.NewUserController(userRepo)

	r = mux.NewRouter()
	r.HandleFunc("/users", userController.Users).Methods("GET")
	r.HandleFunc("/users", userController.Create).Methods("POST")
	r.HandleFunc("/users/{id}", userController.User).Methods("GET")
	r.HandleFunc("/users/{id}", userController.Update).Methods("PATCH")
	r.HandleFunc("/users/{id}", userController.Delete).Methods("DELETE")

	r.Use(middlewares.LoggerMidldlware)
	return
}
