package services

import (
	"github.com/streadway/amqp"
	"order-service/src/rabbit"
)

var RabbitChannel = rabbit.Init()

func PushMessage(msg string) string {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	}

	if err := RabbitChannel.Publish(
		"",
		"test-queue",
		false,
		false,
		message,
	); err != nil {
		return "Can't push message"
	}

	return "Successfully pushed"
}
