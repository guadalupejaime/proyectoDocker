package main

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/guadalupej/proyecto/pkg/models"
)

func processLocations() (*models.Locations, error) {
	response, err := getInfo("https://rickandmortyapi.com/api/location")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	buff := bytes.NewBuffer(response)
	locations := &models.Locations{}
	err = json.NewDecoder(buff).Decode(locations)
	if err != nil {
		log.Println("error in decode ", err)
		return nil, err
	}
	for i, location := range locations.Locations {
		for j, character := range location.Residents {
			if character != "" {
				locations.Locations[i].Residents[j] = strings.Split(character, "character/")[1]
			}
		}
	}
	return locations, nil
}
