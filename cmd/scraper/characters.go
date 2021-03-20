package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/guadalupej/proyecto/pkg/models"
)

func processCharacters() (*models.Characters, error) {

	characterResp := &models.Characters{
		Characters: make([]models.Character, 0),
	}

	characters := &models.Characters{
		Info: models.Info{
			Next: "https://rickandmortyapi.com/api/character",
		},
	}

	for {
		response, err := getInfo(characters.Info.Next)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		buff := bytes.NewBuffer(response)
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
			for j, episode := range character.Episode {
				if episode != "" {
					character.Episode[j] = strings.Split(episode, "episode/")[1]
				}
			}

			characterResp.Characters = append(characterResp.Characters, character)

		}

		if len(characterResp.Characters) >= characters.Info.Count {
			break
		}
	}

	return characterResp, nil
}
