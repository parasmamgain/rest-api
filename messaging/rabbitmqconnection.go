package messaging

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
	"github/parasmamgain/rest-api/configuration"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

/* This method sends a message to rabbit mq queue containing the person id as a data in body*/
func SendMessageToQueue(message string) {
	properties, err := configuration.LoadConfigurationProperties()
	rabbitMqHost := string(properties.RabbitmqProperties.Host)
	rabbitMqPort := properties.RabbitmqProperties.Port
	username := string(properties.RabbitmqProperties.UserName)
	password := string(properties.RabbitmqProperties.Password)
	queueName := string(properties.RabbitmqProperties.QueueName)
	url := fmt.Sprintf("amqp://%s:%s@%s:%d",username, password, rabbitMqHost, rabbitMqPort)

	fmt.Printf("Rabbitmq url is : %s /n",url)

	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	
	q, err := ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

/* This method fetches a message from rabbit mq queue*/
func ReceivingMessageToQueue() (id string){
	fmt.Println("Retriving message from rabbit MQ")
	properties, err := configuration.LoadConfigurationProperties()
	rabbitMqHost := string(properties.RabbitmqProperties.Host)
	rabbitMqPort := properties.RabbitmqProperties.Port
	username := string(properties.RabbitmqProperties.UserName)
	password := string(properties.RabbitmqProperties.Password)
	queueName := string(properties.RabbitmqProperties.QueueName)
	url := fmt.Sprintf("amqp://%s:%s@%s:%d",username, password, rabbitMqHost, rabbitMqPort)

	fmt.Printf("Rabbitmq url is : %s /n",url)


	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	for d := range msgs {
		if d.Body != nil {
			id = string(d.Body)
			break 
		}
	}
	fmt.Printf("Id %s : received from rabbitmq queue", id)
	return id
}