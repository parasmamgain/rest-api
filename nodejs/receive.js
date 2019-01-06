#!/usr/bin/env node
var amqp = require('amqplib/callback_api');
var http = require("http");
var fs = require("fs")

let rawdata = fs.readFileSync('../configuration/configuration.json'); 
let configProp = JSON.parse(rawdata);
console.log("configuration properties loaded :" + JSON.stringify(configProp))
username = configProp.RabbitmqProperties.UserName;
password = configProp.RabbitmqProperties.Password;
rabbitMqHost = configProp.RabbitmqProperties.Host;
rabbitMqPort = configProp.RabbitmqProperties.Port;
queueName = configProp.RabbitmqProperties.QueueName;
amqp.connect('amqp://'+username+':'+password+'@'+rabbitMqHost+':'+rabbitMqPort, function(err, conn) {
    conn.createChannel(function(err, ch) {
        var q = queueName;

        ch.assertQueue(q, {durable: false});
        console.log(" [*] Waiting for messages in %s. To exit press CTRL+C", q);
        ch.consume(q, function(msg) {
        console.log(" [x] Received %s", msg.content.toString());
        personId = msg.content.toString();
        sendMessageReceivedToNodeApp(personId);
        }, {noAck: true});
    });
})


function sendMessageReceivedToNodeApp (personId) {
    console.log("ENtering the function with value" + personId)
      // An object of options to indicate where to post to
      console.log(personId);
        var post_options = {
            host: configProp.NodeJsProperties.Host,
            port: configProp.NodeJsProperties.Port,
            path: '/'+personId,
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                'Content-Length': Buffer.byteLength("")
            }
        };
        // Set up the request
        var post_req = http.request(post_options, function(res) {
            res.setEncoding('utf8');
            res.on('data', function (chunk) {
                console.log('Response: ' + chunk);
            });
        });
        // post the data
        post_req.write("");
        post_req.end();
}