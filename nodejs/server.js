var express = require('express');
var app = express();
var fs = require("fs");

// Mongo db declarations
var mongo = require("mongodb")
var MongoClient = require('mongodb').MongoClient;
var url = "mongodb://9.199.147.126/"

// ElasticSearch declarations
var elasticsearch = require('elasticsearch');
var client = new elasticsearch.Client({
    hosts: [ 'http://localhost:9200']
 });

var server = app.listen(8081, function () {
    var host = server.address().address
    var port = server.address().port
    console.log("Example app listening at http://%s:%s", host, port)
 })

// This responds with "Hello World" on the homepage
app.get('/', function (req, res) {
    console.log("Got a GET request");
    res.send('Hello GET');
 })

 // This responds a GET request with the information of person matching Id
 app.get('/:id',function (req, res) {
    var peopleId = req.params.id;
    console.log("Got a GET request with ID");

    // function to fetch the person details from the mongo Db
    var data = MongoClient.connect(url, function(err, db) {
        if (err) throw err;
        var dbo = db.db("mydb");
        var query = { id: peopleId };
        //Find the first document in the customers collection:
        dbo.collection("people").find(query).toArray(function(err, result) {
          if (err) throw err;
          console.log(result);
          res.end(JSON.stringify(result));
        });
      });
})

 /* This responds a POST request for the homepage. Id recieved in this post request is used to fetch 
    the details of the person which are then pushed to mongoDB and then deletes the same data from
    elastic search index */
app.post('/:id', function (req, res) {
    console.log("Got a POST request");
    var personId = req.params.id;
    console.log("now fetching record from elastic search")
    // ---------------------------------------------------------------------
    // function to retreive the person details from the elastic search index
    // ---------------------------------------------------------------------
    client.search({
        index: 'people',
        q: 'id:'+personId
    }).then(function(resp) {
        console.log("--- Hits ---");
        resp.hits.hits.forEach(function(hit){
        //console.log(hit._source);
        person = hit._source;
      })
    }, function(err) {
        console.trace(err.message);
    }),
    // ---------------------------------------------------------------------
    // function to upload the details of the person to mongoDB
    // ---------------------------------------------------------------------
    MongoClient.connect(url, function(err, db) {
        console.log("now inserting record into mongoDB")
        if (err) throw err;
        var dbo = db.db("mydb");        
        var myobj = { id: person.id, firstname : person.firstname ,lastname : person.lastname , address: { city : person.address.city , state : person.address.state} };
        dbo.collection("people").insertOne(myobj, function(err, res) {
          if (err) throw err;
          console.log("1 document inserted");
          db.close();
        });
      }),
      // ---------------------------------------------------------------------
      // function to delete the person details from the elastic search index
      // ---------------------------------------------------------------------
      client.deleteByQuery({
        index: 'people',        
        body: {
           query: {
               match: { id : personId }
           }
        }
    }, function (error, response) {
        console.log(response);
    });
    res.send();
 })
