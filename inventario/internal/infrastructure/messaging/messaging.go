package messaging

import (
	"github.com/streadway/amqp"
)

type RabbitMQClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQClient(url string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQClient{conn: conn, ch: ch}, nil
}

func (client *RabbitMQClient) GetChannel() *amqp.Channel {
	return client.ch
}

func (client *RabbitMQClient) Consume(queueName string, handler func([]byte)) error {
	messages, err := client.ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for message := range messages {
		handler(message.Body)
	}
	return nil
}

func (client *RabbitMQClient) Close() {
	client.ch.Close()
	client.conn.Close()
}
