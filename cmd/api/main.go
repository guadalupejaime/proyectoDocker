package main

import (
	"log"
	"os"

	"github.com/guadalupej/proyecto/pkg/characters"
	"github.com/guadalupej/proyecto/pkg/episodes"
	"github.com/guadalupej/proyecto/pkg/http"
	"github.com/guadalupej/proyecto/pkg/locations"
	"github.com/guadalupej/proyecto/pkg/mongo"
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

	controller := http.Controller{
		CharacterService: *characters.NewService(db),
		LocationService:  *locations.NewService(db),
		EpisodeService:   *episodes.NewService(db),
	}

	http.ListenAndServe(controller)
	log.Println("terminó")
}
