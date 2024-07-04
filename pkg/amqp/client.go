package amqp

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   *amqp.Queue
}

type ClientConfig struct {
	Protocol string
	Host     string
	User     string
	Password string
	Port     int
	VHost    string
}

func NewClient(config ClientConfig) *Client {
	uri := fmt.Sprintf("%s://%s:%s@%s:%d%s", config.Protocol, config.User, config.Password, config.Host, config.Port, config.VHost)

	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ", err)
	}

	client := &Client{
		conn: conn,
	}

	client.channel = client.CreateChannel()

	return client
}

func (a *Client) CreateChannel() *amqp.Channel {
	ch, err := a.conn.Channel()

	if err != nil {
		log.Fatal("Failed to open a channel")
	}

	return ch
}

func (a *Client) RegisterQueue(name string) *Client {
	q, err := a.channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("Failed to declare queue", err)
	}

	a.queue = &q
	return a
}

func (a *Client) Write(p []byte) (int, error) {
	err := a.channel.PublishWithContext(
		context.Background(),
		"",
		a.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        p,
		})

	if err != nil {
		return 0, err
	}

	return len(p), nil
}
