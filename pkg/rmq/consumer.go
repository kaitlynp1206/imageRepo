package rmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	connection *amqp.Connection
	channel    *ampq.Channel
}

func SetupConsumer() *Consumer {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	return &Consumer{
		connection: conn,
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
