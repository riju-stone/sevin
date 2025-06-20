package services

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/riju-stone/sevin/api/utils"
)

type RabbitMQClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

var TaskQueue *RabbitMQClient

func InitTaskQueue() {
	l := utils.CustomLogger
	connUrl := os.Getenv("RABBITMQ_CONN_URL")
	l.Debugf("Connecting to RabbitMQ: %s", connUrl)

	// Connect to RabbitMQ
	conn, err := amqp.Dial(connUrl)
	if err != nil {
		l.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	// Open a RabbitMQ Channel
	ch, err := conn.Channel()
	if err != nil {
		l.Fatalf("Failed to open a channel: %v", err)
	}

	l.Infof("Connected to RabbitMQ")

	// Store the connection and channel in the TaskQueue variable
	TaskQueue = &RabbitMQClient{
		Conn:    conn,
		Channel: ch,
	}
}
