package mq

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	ch *amqp.Channel
}

func GetRabbitMQ(conf *viper.Viper) *RabbitMQ {
	if !conf.GetBool("rabbitmq.enable") {
		return nil
	}
	username := conf.GetString("rabbitmq.username")
	password := conf.GetString("rabbitmq.password")
	host := conf.GetString("rabbitmq.host")
	port := conf.GetInt("rabbitmq.port")
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, port)
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return &RabbitMQ{
		ch: ch,
	}
}

func (mq *RabbitMQ) Send(b []byte) error {
	return mq.ch.Publish(
		"",
		"log",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b,
		})
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
