# rest-api

This project is under progress and requires some more work to be done.

This Go lang project provides rest api endpoints for the URL `localhost:8000/people` for GET and POST methods. The POST methods accepts a body withthe the detail of a person as the request body. Sample of the request body is 
```{
"id":"1",
"firstname":"paras",
"lastname":"mamgain",
"address" :{
	"city":"delhi",
	"state":"New Delhi"
	}
}```