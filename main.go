package main

import (
	///"encoding/json"
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github/parasmamgain/rest-api/restapi"
	"github/parasmamgain/rest-api/configuration"
)

func main() {
	router := mux.NewRouter()
	_, err := configuration.LoadConfigurationProperties()
	if err != nil {
		fmt.Printf("Error while loading the properties %s /n", err.Error)
	}
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