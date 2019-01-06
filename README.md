# Go-Node-js-Rest Api

This project is under progress and requires some more work to be done.

### Pre-Requisite

1. clone this project and make sure that you have following tools configuread, installed and running in you machine
2. Go lang
3. ElasticSearch 
4. Rabbitmq
5. Postman or any other similar tool
6. MongoDB
7. NodeJs

### Workflow

This application provides two endpoints
1. GET
2. POST

1. GET : `localhost:8081/{id}` : retreives the information of the person with id sent with the url from the MongoDB Database. This information is then shared in the response body.

   -  This request is being served by the Node Js module  

2. POST : `localhost:8000/people` : The POST methods accepts a body with the detail of a person as the request body. Sample of the request body is

	```
	{
	"id":"1",
	"firstname":"paras",
	"lastname":"mamgain",
	"address" :{
		"city":"delhi",
		"state":"New Delhi"
		}
	}
	```
   - This request is being served by the Go Lang module.

## Functionality of the GET-POST Requests

Once a POST request is being made with above obdy it pushes this data into the elastic search and at the same time send the `{id}` to the rabbitmq queue. This queue holds the id's of all the persons for which the POST request has been made. 

when a GET request is being made using a URL `localhost:8081/{id}` then it retrieves the person details belonging to `{id}` from the mongo returns it as a response.

1. /main.go - this is the initial point to start the application. `go run main.go`
2. /models/people.go - this file contains the entity object/structure of the person object.
3. /restapi/people.go - this file contains the method that handles the GET and POST request for the endpoint `localhost:8000/people`.
    - POST - is handled by the method `CreatePeople`
4. /datastore/elasticconnection.go - handles the code that pushes the data sent from the POST request to the elastic search index.
5. /messaging/rabbitmqconnection.go : this file handles the code that communicates with the rabbitmq. This sends and receives message from the queue present in the rabbitmq server.

6. /configuration/configuration.json :  a property file that can be used to configure the server hostname, server portnumber, datastore server  hostname/portnumber etc.


7. /nodeJS/*.js : a module that runs on the rabbitmq listeners and waits for the new notification. Once recieved it fetches this data from the elastic search index and pushes it to the mondo DB and then deletes the data from the Elastic Search index.
    
	- nodeJs/receive.js : waits and listens to the rabbit mq queue for the new entry in the elastic search index
	- nodejs/server.js : runs the rest api end point for GET request and various other modules that fecthces data from elastic search index and pushes it to the mongoDB and then deletes the same data from the elastic search index. 


## More Detailed Description

1.	A simple POST endpoint in Golang(i.e. `localhost:8000/people`) which accepts a data-body of your choice (in this case i have used People). 
2.	On receiving the data-body, the Golang service writes this data to Elastic Search Index.
3.	The Golang service then uses a messaging queue ( rabbit mq) to inform a Node service that there is new data in Elastic Search Index.
4.	The Node service takes the data-body from Elastic Search Index and writes it to a SQL / NoSQL database (Used MongoDB for the demo).
5.	Once that is done, the Node service deletes the data-body from Elastic Search Index.
6.	Write a GET endpoint( i.e. `localhost:8081/{id}`) in Node which fetches data from your Mongo DB database and serves it as a JSON response
7.	Deploy it on a server and send us the following:
	- a.	Link to your Git project
	- b.	POST and GET endpoints with details of the data to be posted


##  Starting application

You will have to start the application by opening three terminals. lets name them `terminal 1`, `terminal 2` and `terminal 3`. 
1. In `Terminal 1` go to directory `/rest-api/` and run command `go run main.go` . This commands starts the go modules.
2. In `Terminal 2` go to directory `/rest-api/nodejs/` and run command `node server.js` . This command starts the nodeJs modules available.
3. In `Terminal 3` go to directory `/rest-api/nodejs/` and run command `./receive.js`. This command starts the receives.js that continously listens on rabbitmq queue and expects a id to be pushed in the queue.