package main

import (
	///"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
//	"github/parasmamgain/rest-api/models"
	"github/parasmamgain/rest-api/restapi"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people", createPeople).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	restapi.GetPeople(w, r)
	
}

func createPeople(w http.ResponseWriter, r *http.Request) {
	restapi.CreatePeople(w, r)
}