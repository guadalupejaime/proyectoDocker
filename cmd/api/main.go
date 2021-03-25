package main

import (
	"log"
	"os"

	"github.com/guadalupej/proyecto/pkg/characters"
	"github.com/guadalupej/proyecto/pkg/episodes"
	"github.com/guadalupej/proyecto/pkg/http"
	"github.com/guadalupej/proyecto/pkg/locations"
	"github.com/guadalupej/proyecto/pkg/mongo"
	"github.com/guadalupej/proyecto/rabbit/publisher"
)

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
	log.Println(rabbitURL)
	qb := publisher.InitRabbit(rabbitURL)

	defer qb.Close()

	controller := http.Controller{
		CharacterService: *characters.NewService(db),
		LocationService:  *locations.NewService(db),
		EpisodeService:   *episodes.NewService(db),
	}

	http.ListenAndServe(controller, qb)
	log.Println("terminó")
}
