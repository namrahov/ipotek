package queue

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type IMessageSender interface {
	SendMessage(message interface{}, queueName string) error
}
type MessageSender struct {
}

func (*MessageSender) SendMessage(message interface{}, queueName string) error {
	log.Infof("RabbitMQ.SendQMessage.start %s", message)

	ch, conn := NewRabbitChannel()
	defer ch.Close()
	defer conn.Close()
	b, err := json.Marshal(message)
	if err != nil {
		log.Error(err)
		return err
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
	if err != nil {
		log.Errorf("Failed to publish a message %v", err.Error())
		return err
	}

	log.Infof("RabbitMQ.SendQMessage.success %s into queue %s", message, queueName)

	return nil
}
