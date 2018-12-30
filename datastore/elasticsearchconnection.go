package elastic

import (
	"gopkg.in/olivere/elastic.v5"
	"github/parasmamgain/rest-api/models"
	"fmt"
	"errors"
	"golang.org/x/net/context"
	"encoding/json"
)

func PushToElasticSearch(p models.Person){
	createConnection(p)
}

func createConnection(p models.Person) {
	fmt.Println("Creating Connection")
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Created Connection")
	err = createIndexWithLogsIfDoesNotExist(client,p)
	if err != nil {
		panic(err)
	}
}

func createIndexWithLogsIfDoesNotExist(client *elastic.Client,p models.Person) error {
	fmt.Println("Creating Index")
	indexName := "peopleindex1"
	exists, err := client.IndexExists(indexName).Do(context.TODO())
	if err != nil {
		return err
	}

	if exists {
		return addLogsToIndex(client, p)
	}
	fmt.Println("Index Created")
	// sending data to elastic search index
	res, err := client.CreateIndex(indexName).Body("").Do(context.TODO())

	if err != nil {
		return err
	}
	if !res.Acknowledged {
		return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}

	return addLogsToIndex(client, p)
}

func addLogsToIndex(client *elastic.Client, p models.Person) error {
	fmt.Println("Adding record")
	indexName := "peopleindex1"
	docType := "personaldetails"
		address := models.Address{
			 City :p.Address.City,
			 State:p.Address.State, 
		}
		person := models.Person{
			ID: p.ID,
			Firstname: p.Firstname,
			Lastname: p.Lastname,
			Address: &address ,
		}

		_, err := client.Index().Index(indexName).Type(docType).Id(p.ID).BodyJson(person).Do(context.TODO())
			fmt.Println("record added")
		if err != nil {
			return err
		}
		
	return nil
}

func PullFromElasticSearch(id string) (p models.Person){
	fmt.Println("Creating Connection")
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Created Connection")
	indexName := "peopleindex1"
	
	// creating a query

	//docType := "personaldetails"
	get1, err := client.Get().
		Index(indexName).
		Id(id).
		Do(context.TODO())
	if err != nil {
		// Handle error
		panic(err)
	}
	
	
	if get1.Found {
		//fmt.Printf("%s", get1.Source)
		fmt.Println()
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		err := json.Unmarshal(*get1.Source, &p)
		
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("Data for User:%s with UserId:%s received.",p.ID, p.Firstname)
	}
	return p
}