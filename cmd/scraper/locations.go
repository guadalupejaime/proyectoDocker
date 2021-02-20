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
	for _, location := range locations.Locations {
		for i, character := range location.Residents {
			location.Residents[i] = strings.Split(character, "character/")[1]
		}
	}
	return locations, nil
}
