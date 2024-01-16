package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vispro-project/BackEnd-LombaKu.git/controllers/lombacontroller"
	"github.com/vispro-project/BackEnd-LombaKu.git/controllers/lombacontroller/authcontroller"
	"github.com/vispro-project/BackEnd-LombaKu.git/middlewares"
	"github.com/vispro-project/BackEnd-LombaKu.git/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/lomba", lombacontroller.Index).Methods("GET")
	api.HandleFunc("/lomba/{id}", lombacontroller.Show).Methods("GET")
	api.HandleFunc("/lomba", lombacontroller.Create).Methods("POST")
	api.HandleFunc("/lomba/{id}", lombacontroller.Index).Methods("PUT")
	api.HandleFunc("/lomba/{id}", lombacontroller.Delete).Methods("DELETE")
	api.Use(middlewares.JWTMiddleware)

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
