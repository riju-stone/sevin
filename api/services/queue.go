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

var taskQueue *amqp.Queue

func ConnectToRabbitMQ() (*RabbitMQClient, error) {
	l := utils.CustomLogger
	connUrl := os.Getenv("RABBITMQ_CONN_URL")
	l.Debugf("Connecting to RabbitMQ: %s", connUrl)

	// Connect to RabbitMQ
	conn, err := amqp.Dial(connUrl)
	if err != nil {
		return nil, err
	}

	// Open a RabbitMQ Channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	l.Infof("Connected to RabbitMQ")
	return &RabbitMQClient{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func (r *RabbitMQClient) InitTaskQueue() {
	l := utils.CustomLogger
	l.Debugf("Initializing task queue")
	q, err := r.Channel.QueueDeclare(
		"sevin_tasks",
		true,  // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		l.Fatalf("Failed to declare task queue: %v", err)
	}

	l.Infof("Task queue initialized")
	taskQueue = &q
}
