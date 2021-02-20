package main

import (
	"log"
	"net/http"
	"os"

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
	log.Println("process characters...")
	characters, err := processCharacters()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("insert characters...")
	for _, character := range characters.Characters {
		err = db.InsertCharacter(character)
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("process locations...")
	locations, err := processLocations()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("insert locations...")
	for _, location := range locations.Locations {
		err = db.InsertLocation(location)
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("process episodes...")
	episodes, err := processEpisodes()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("insert episodes...")
	for _, episode := range episodes.Episodes {
		err = db.InsertEpisode(episode)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func getInfo(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("error in request ", err)
		return nil, err
	}

	var body []byte
	_, err = response.Body.Read(body)
	if err != nil {
		log.Println("error in Read ", err)
		return nil, err
	}
	return body, nil
}
