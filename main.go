package main

import (
	"log"
	"net/http"

	"task-5-vix-btpns/middlewares"

	"task-5-vix-btpns/controllers/authcontroller"
	"task-5-vix-btpns/controllers/photocontroller"
	"task-5-vix-btpns/database"

	"github.com/gorilla/mux"
)

func main() {

	database.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/photo", photocontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8888", r))
}
