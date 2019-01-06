var express = require('express');
var app = express();
var fs = require("fs");
let rawdata = fs.readFileSync('../configuration/configuration.json'); 
let configProp = JSON.parse(rawdata);
console.log("configuration properties loaded :" + JSON.stringify(configProp))


// Mongo db declarations
var mongo = require("mongodb")
var MongoClient = require('mongodb').MongoClient;
var mongoHost = configProp.MongoProperties.Host;
var mongoPort = configProp.MongoProperties.Port;
var mongoDbName = configProp.MongoProperties.DbName;
var mongoCollectionName = configProp.MongoProperties.CollectionName;
var url = "mongodb://"+ mongoHost +"/"

// ElasticSearch declarations
var elasticsearch = require('elasticsearch');
var elasticHost = configProp.ElasticProperties.Host;
var elasticPort = configProp.ElasticProperties.Port;
var elasticIndex = configProp.ElasticProperties.IndexName;
var client = new elasticsearch.Client({
    hosts: [ 'http://' + elasticHost + ':' + elasticPort ]
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
        var dbo = db.db(mongoDbName);
        var query = { id: peopleId };
        //Find the first document in the customers collection:
        dbo.collection(mongoCollectionName).find(query).toArray(function(err, result) {
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
        index: elasticIndex,
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
        var dbo = db.db(mongoDbName);        
        var myobj = { id: person.id, firstname : person.firstname ,lastname : person.lastname , address: { city : person.address.city , state : person.address.state} };
        dbo.collection(mongoCollectionName).insertOne(myobj, function(err, res) {
          if (err) throw err;
          console.log("1 document inserted");
          db.close();
        });
      }),
      // ---------------------------------------------------------------------
      // function to delete the person details from the elastic search index
      // ---------------------------------------------------------------------
      client.deleteByQuery({
        index: elasticIndex,        
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
