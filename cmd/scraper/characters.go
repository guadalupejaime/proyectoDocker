package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/guadalupej/proyecto/pkg/models"
)

func processCharacters() (*models.Characters, error) {
	response, err := getInfo("https://rickandmortyapi.com/api/character")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	buff := bytes.NewBuffer(response)
	characters := &models.Characters{}
	err = json.NewDecoder(buff).Decode(characters)
	if err != nil {
		log.Println("error in decode ", err)
		return nil, err
	}
	for _, character := range characters.Characters {
		if character.Location.URL != "" {
			character.Location.URL = strings.Split(character.Location.URL, "location/")[1]
		}
		if character.Origin.URL != "" {
			character.Origin.URL = strings.Split(character.Origin.URL, "location/")[1]
		}
		for i, episode := range character.Episode {
			if episode != "" {
				character.Episode[i] = strings.Split(episode, "episode/")[1]
			}
		}
	}
	return characters, nil
}
