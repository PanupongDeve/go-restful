package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("Get")
	router.HandleFunc("/people/${id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("5555")
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

}

func CreatePerson(w http.ResponseWriter, r *http.Request) {

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {

}
