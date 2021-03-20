package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/guadalupej/proyecto/pkg/models"
)

func processEpisodes() (*models.Episodes, error) {

	episodesResp := &models.Episodes{
		Episodes: make([]models.Episode, 0),
	}

	episodes := &models.Episodes{
		Info: models.Info{
			Next: "https://rickandmortyapi.com/api/episode",
		},
	}

	for {
		response, err := getInfo(episodes.Info.Next)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		buff := bytes.NewBuffer(response)

		err = json.NewDecoder(buff).Decode(episodes)
		if err != nil {
			log.Println("error in decode ", err)
			return nil, err
		}

		for _, episode := range episodes.Episodes {
			for j, character := range episode.Characters {
				if character != "" {
					episode.Characters[j] = strings.Split(character, "character/")[1]
				}
			}

			episodesResp.Episodes = append(episodesResp.Episodes, episode)
		}

		if len(episodesResp.Episodes) >= episodes.Info.Count {
			break
		}

	}
	return episodesResp, nil
}
