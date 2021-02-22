package models

import (
	"net/http"
	"time"

	"github.com/guadalupej/proyecto/pkg/characters"
)

type Characters struct {
	Characters []Character `bson:"results" json:"results"`
}

type Character struct {
	ID       int          `bson:"_id" json:"id"`
	Name     string       `bson:"name" json:"name"`
	Status   string       `bson:"status" json:"status"`
	Species  string       `bson:"species" json:"species"`
	Type     string       `bson:"type" json:"type"`
	Gender   string       `bson:"gender" json:"gender"`
	Origin   OriginTiny   `bson:"origin" json:"origin"`
	Location LocationTiny `bson:"location" json:"location"`
	Image    string       `bson:"image" json:"image"`
	Episode  []string     `bson:"episode" json:"episode"`
	Created  time.Time    `bson:"created" json:"created"`
}

type OriginTiny struct {
	Name string `bson:"name" json:"name"`
	URL  string `bson:"url" json:"url"`
}
type LocationTiny struct {
	Name string `bson:"name" json:"name"`
	URL  string `bson:"url" json:"url"`
}

func (mt *Characters) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ToCharacterModel(character characters.Character) *Character {
	origin := OriginTiny{
		Name: character.Origin.Name,
		URL:  character.Origin.URL,
	}
	location := LocationTiny{
		Name: character.Location.Name,
		URL:  character.Location.URL,
	}
	return &Character{
		ID:       character.ID,
		Name:     character.Name,
		Status:   character.Name,
		Species:  character.Name,
		Type:     character.Name,
		Gender:   character.Name,
		Image:    character.Name,
		Episode:  character.Episode,
		Created:  character.Created,
		Origin:   origin,
		Location: location,
	}
}
