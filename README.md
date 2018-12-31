# rest-api

This project is under progress and requires some more work to be done.

This Go lang project provides rest api endpoints for the URL `localhost:8000/people` for GET and POST methods. The POST methods accepts a body with the detail of a person as the request body. Sample of the request body is

```{
"id":"1",
"firstname":"paras",
"lastname":"mamgain",
"address" :{
	"city":"delhi",
	"state":"New Delhi"
	}
}
```

Once a POST request is being made with above obdy it pushes this data into the elastic search and at the same time send the `id` to the rabbitmq queue. This queue holds the id's of all the persons for which the POST request has been made. 
when a GET request is being made using a URL `localhost:8000/people` then it retrieves the `id` from the rabbitmq queue and then it uses this `id` to retrieve the entire data object from elastic search data store and displays it as a response.

1. /main.go - this is the initial point to start the application. `go run main.go`
2. /models/people.go - this file contains the entity object/structure of the person object.
3. /restapi/people.go - this file contains the method that handles the GET and POST request for the endpoint `localhost:8000/people`.
    - GET - is handled by the method `GetPeople`
    - POST - is handled by the method `CreatePeople`
4. /datastore/elasticconnection.go - handles the code that pushes the data sent from the POST request to the elastic search index.
5. /messaging/rabbitmqconnection.go : this file handles the code that communicates with the rabbitmq. This sends and receives message from the queue present in the rabbitmq server.


### todo

6. /property.env :  a property file that can be used to configure the server hostname, server portnumber, datastore server  hostname/portnumber etc.
7. /nodeJS/*.go : a module that runs on the rabbitmq/ibmmq listeners and waits for the new notification. Once recieved it deletes the data from the datastore.

### Pre-Requisite

1. clone this project and make sure that you have following tools configuread, installed and running in you machine
2. Go lang
3. ElasticSearch 
4. Rabbitmq
5. Postman or any other similar tool

### Workflow

This application provides two endpoints