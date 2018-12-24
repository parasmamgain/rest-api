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

1. /main.go - this is the initial point to start the application. `go run main.go`
2. /models/people.go - this file contains the entity object/structure of the person object.
3. /restapi/people.go - this file contains the method that handles the GET and POST request for the endpoint `localhost:8000/people`.
    - GET - is handled by the method `GetPeople`
    - POST - is handled by the method `CreatePeople`
4. /datastore/elasticconnection.go - handles the code that pushes the data sent from the POST request to the elastic search index.

### todo

5. /messaging/*.go : this folder should contain the file that handles the request/response to ibmmq, rabbitmq or any other messaging application.
6. /property.env :  a property file that can be used to configure the server hostname, server portnumber, datastore server  hostname/portnumber etc.
