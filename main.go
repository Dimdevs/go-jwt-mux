package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypac/go-jwt-mux/controllers/authcontroller"
	"github.com/jeypac/go-jwt-mux/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/api/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/api/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/api/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":6000", r))
}
