package queue

import (
	"fmt"
	"github.com/PB-Digital/ms-retail-products-info/properties"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func NewRabbitChannel() (*amqp.Channel, *amqp.Connection) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", properties.Props.RabbitMqUser,
		properties.Props.RabbitMqPass, properties.Props.RabbitMqHost, properties.Props.RabbitMqPort))
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return ch, conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
