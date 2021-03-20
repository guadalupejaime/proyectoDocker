package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/guadalupej/proyecto/pkg/models"
)

func processLocations() (*models.Locations, error) {

	locationResp := &models.Locations{
		Locations: make([]models.Location, 0),
	}

	locations := &models.Locations{
		Info: models.Info{
			Next: "https://rickandmortyapi.com/api/location",
		},
	}

	for {

		response, err := getInfo(locations.Info.Next)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		buff := bytes.NewBuffer(response)

		err = json.NewDecoder(buff).Decode(locations)
		if err != nil {
			log.Println("error in decode ", err)
			return nil, err
		}
		for _, location := range locations.Locations {
			for j, character := range location.Residents {
				if character != "" {
					location.Residents[j] = strings.Split(character, "character/")[1]
				}
			}

			locationResp.Locations = append(locationResp.Locations, location)
		}

		if len(locationResp.Locations) >= locations.Info.Count {
			break
		}

	}
	return locationResp, nil
}
