package publisher

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection

type QueueBroker struct {
	Q  amqp.Queue
	CH *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func InitRabbit(url string) QueueBroker {
	conn, err := amqp.Dial(url)
	if err != nil {
		for i := 0; i < 5 && err != nil; i++ {
			log.Println("Failed to connect to RabbitMQ")
			time.Sleep(5 * time.Second)
			conn, err = amqp.Dial(url)
		}
		failOnError(err, "Failed to connect to RabbitMQ")
	}

	connection = conn
	// defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	q, err := ch.QueueDeclare(
		"post_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	qb := QueueBroker{
		Q:  q,
		CH: ch,
	}

	return qb
}

func (qb QueueBroker) NewMessage(content []byte, collection string) error {

	err := qb.CH.Publish(
		"",        // exchange
		qb.Q.Name, // routing key
		false,     // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         content,
			Type:         collection,
		})
	if err != nil {
		log.Println("Failed to publish a message", err)
		return err
	}
	log.Printf(" [x] Sent %s", string(content))
	return nil
}

func (qb QueueBroker) Close() error {
	qb.CH.Close()
	connection.Close()
	return nil
}
