package elastic

import (
	"fmt"
	"errors"
	"encoding/json"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"github/parasmamgain/rest-api/models"
	"github/parasmamgain/rest-api/configuration"
)

/*
Method that establishes connection with the elastic search server
*/
func createConnection(p models.Person) (client *elastic.Client, err error){
	fmt.Println("Creating Connection")
	properties, err := configuration.LoadConfigurationProperties()
	elasticSearchHost := string(properties.ElasticProperties.Host)
	elasticSearchPort := properties.ElasticProperties.Port

	url := fmt.Sprintf("http://%s:%d",elasticSearchHost,elasticSearchPort)
	fmt.Printf("Elastic search url is : %s /n",url)
	client, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		panic(err)
	}
	fmt.Println("Created Connection")
	if err != nil {
		panic(err)
	}
	return client, err
}

/*
Method that checks if the given indexname exists or not, If not then it creates an index with
indexName as configured in configuration.json
*/
func createIndexWithLogsIfDoesNotExist(client *elastic.Client,p models.Person) error {
	fmt.Println("Checkking if Index exists or not")
	properties, err := configuration.LoadConfigurationProperties()
	indexName := string(properties.ElasticProperties.IndexName)
	exists, err := client.IndexExists(indexName).Do(context.TODO())
	if err != nil {
		return err
	}
	
	if !exists {
		res, err := client.CreateIndex(indexName).Body("").Do(context.TODO())
		if err!= nil {
			return err
		}
		if !res.Acknowledged {
			return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
		}
		fmt.Println("Index Created")
	}
	if err != nil {
		return err
	}
	return err
}


func addLogsToIndex(client *elastic.Client, p models.Person) error {
	fmt.Println("Adding record")
	properties, err := configuration.LoadConfigurationProperties()
	indexName := string(properties.ElasticProperties.IndexName)
	//indexName := "peopleindex1"
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

		_, err = client.Index().Index(indexName).Type(docType).Id(p.ID).BodyJson(person).Do(context.TODO())
			fmt.Println("record added")
		if err != nil {
			return err
		}
		
	return nil
}

/*
This method pushes data to the elastic search index
*/
func PushToElasticSearch(p models.Person){
	client, err := createConnection(p)
	if err != nil {
		fmt.Printf("Connection could not be established")
		fmt.Printf("%s",err)
	}

	//check if the index exists or not, if not then create a new index
	err = createIndexWithLogsIfDoesNotExist(client,p)
	if err == nil{
		// push data to the index
		addLogsToIndex(client, p)
	}
}

/*
pulls the data from the elastic search index containing id is the unique value
*/
func PullFromElasticSearch(id string) (p models.Person){
	properties, err := configuration.LoadConfigurationProperties()
	indexName := string(properties.ElasticProperties.IndexName)
	client, err := createConnection(p)
	if err != nil {
		fmt.Printf("Connection could not be established")
		fmt.Printf("%s",err)
	}

	if err != nil {
		panic(err)
	}
	fmt.Println("Created Connection")
	get1, err := client.Get().
		Index(indexName).
		Id(id).
		Do(context.TODO())
	if err != nil {
		// Handle error
		panic(err)
	}
	
	
	if get1.Found {
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