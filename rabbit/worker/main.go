package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/mongo"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	dbURL := os.Getenv("MONGOURL")
	dbName := os.Getenv("ME_CONFIG_MONGODB_AUTH_DATABASE")
	user := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	pass := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	log.Println("connect storage...")
	db, err := mongo.NewStorage(dbURL, dbName, user, pass)
	if err != nil {
		log.Println(err)
		return
	}

	rabbitURL := os.Getenv("RABBITMQURL")
	conn, err := amqp.Dial(rabbitURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"post_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
			SaveMessage(d.Type, d.Body, db)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func SaveMessage(collection string, content []byte, db *mongo.Repository) error {
	switch collection {
	case "characters":
		data := models.Character{}
		err := json.Unmarshal(content, &data)
		if err != nil {
			log.Println(err)
			return err
		}
		data.Created = time.Now()
		db.InsertCharacter(data)
	case "episodes":
		data := models.Episode{}
		err := json.Unmarshal(content, &data)
		if err != nil {
			log.Println(err)
			return err
		}
		data.Created = time.Now()
		db.InsertEpisode(data)
	case "locations":
		data := models.Location{}
		err := json.Unmarshal(content, &data)
		if err != nil {
			log.Println(err)
			return err
		}
		data.Created = time.Now()
		db.InsertLocation(data)
	}
	return nil
}
