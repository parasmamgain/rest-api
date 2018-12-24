package elastic

import (
	"gopkg.in/olivere/elastic.v5"
	"github/parasmamgain/rest-api/models"
	"fmt"
	"errors"
	"golang.org/x/net/context"
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

		_, err := client.Index().Index(indexName).Type(docType).BodyJson(person).Do(context.TODO())
			fmt.Println("record added")
		if err != nil {
			return err
		}
		
	return nil
}
