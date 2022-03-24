package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() {
	log.Println("Start development server localhost:8000")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Hello).Methods("OPTIONS", "GET")
	router.HandleFunc("/user", CreateUser).Methods("OPTIONS", "POST")
	router.HandleFunc("/users", GetUsers).Methods("OPTIONS", "GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("OPTIONS", "GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
