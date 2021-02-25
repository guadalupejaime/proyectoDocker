package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/guadalupej/proyecto/pkg/models"
)

func processEpisodes() (*models.Episodes, error) {
	response, err := getInfo("https://rickandmortyapi.com/api/episode")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	buff := bytes.NewBuffer(response)
	episodes := &models.Episodes{}
	err = json.NewDecoder(buff).Decode(episodes)
	if err != nil {
		log.Println("error in decode ", err)
		return nil, err
	}
	for i, episode := range episodes.Episodes {
		for j, character := range episode.Characters {
			if character != "" {
				episodes.Episodes[i].Characters[j] = strings.Split(character, "character/")[1]
			}
		}
	}
	return episodes, nil
}
