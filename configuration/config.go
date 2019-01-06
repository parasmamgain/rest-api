package configuration

import(
	"fmt"
	"os"
	"encoding/json"
)

type Configuration struct {
	ElasticProperties  ElasticSearch
	RabbitmqProperties Rabbitmq		
}

type ElasticSearch struct {
	Host      string
	Port      int
	IndexName string
}

type Rabbitmq struct {
	Host          string
	Port          int
	QueueName	  string
	UserName	  string
	Password	  string
}

type NodeJs struct {
	Host          string
	Port          int
}

type MongoDb struct {
	Host		   string
	Port		   int
	DbName		   string
	CollectionName string
}

func LoadConfigurationProperties() (configurationProperties Configuration, err error) {
	var configProperties Configuration
	//filename is the path to the json config file
	file, err := os.Open("./configuration/configuration.json") 
	if err != nil {
		  return configProperties, err 
	}  
	decoder := json.NewDecoder(file) 
	err = decoder.Decode(&configProperties) 
	if err != nil {  
		return configProperties, err 
	}
	fmt.Println("Properties loaded from configuration file")
	return configProperties, err
}