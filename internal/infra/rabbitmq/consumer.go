package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan bool
}

func NewConsumer(amqpURI string) (*Consumer, error) {
	c := &Consumer{
		done: make(chan bool),
	}

	var err error
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Consumer) Consume(queueName string, handler func([]byte)) error {
	queue, err := c.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := c.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()

	log.Printf("Waiting for messages on queue: %s. To exit press CTRL+C", queueName)
	<-c.done

	return nil
}

func (c *Consumer) Close() error {
	close(c.done)
	if err := c.channel.Close(); err != nil {
		return err
	}
	if err := c.conn.Close(); err != nil {
		return err
	}
	return nil
}