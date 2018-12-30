package restapi

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github/parasmamgain/rest-api/models"
	"github/parasmamgain/rest-api/datastore"
	"github/parasmamgain/rest-api/messaging"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method invoked in Get People")

	// getting unique Id from the rabbit-mq
	id := messaging.ReceivingMessageToQueue()

	//sending request to elastic search with ID received from rabbitmq
	fmt.Println("Id received : %s", id)
	p := elastic.PullFromElasticSearch(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)

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

	// sending data to elastic search
	elastic.PushToElasticSearch(p)

	// sending the unique Id of the individual in rabbit-mq
	messaging.SendMessageToQueue(p.ID)
}