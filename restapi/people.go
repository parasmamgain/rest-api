package restapi

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github/parasmamgain/rest-api/models"
	"github/parasmamgain/rest-api/datastore"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method invoked in Get People")
}

func CreatePeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method invoked in Create People Method")
	decoder := json.NewDecoder(r.Body)
	var p models.Person
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Data for User:%s with UserId:%s received.",p.ID, p.Firstname)
	fmt.Println("Sending data to elastic search server")
	elastic.PushToElasticSearch(p)


}