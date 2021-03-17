package models

import (
	"errors"
	"net/http"
	"time"
)

type Characters struct {
	Characters    []Character `bson:"results" json:"results"`
	TotalFound    int         `bson:"total_found" json:"total_found"`
	TotalReturned int         `bson:"total_returned" json:"total_returned"`
}

type Character struct {
	ID       int          `bson:"_id" json:"id" fake:"{number:1,10}"`
	Name     string       `bson:"name" json:"name" fake:"{firstname}"`
	Status   string       `bson:"status" json:"status" fake:"{randomstring:[Alive,unknown,Dead]}"`
	Species  string       `bson:"species" json:"species" fake:"{randomstring:[Human,Alien]}"`
	Type     string       `bson:"type" json:"type" fake:"{lastname}"`
	Gender   string       `bson:"gender" json:"gender" fake:"{gender}"`
	Origin   OriginTiny   `bson:"origin" json:"origin" fake:"{struct}"`
	Location LocationTiny `bson:"location" json:"location" fake:"{struct}"`
	Image    string       `bson:"image" json:"image" fake:"{imageurl:[20,20]}"`
	Episode  []string     `bson:"episode" json:"episode" fake:"{name}" fakesize:"3"`
	Created  time.Time    `bson:"created" json:"created" fake:"{date}"`
}

func (mt *Character) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type OriginTiny struct {
	Name string `bson:"name" json:"name" fake:"{firstname}"`
	URL  string `bson:"url" json:"url" fake:"{imageurl:[20,20]}"`
}

type LocationTiny struct {
	Name string `bson:"name" json:"name" fake:"{firstname}"`
	URL  string `bson:"url" json:"url" fake:"{imageurl:[20,20]}"`
}

func (mt *Characters) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type CharactersFilters struct {
	Limit    int
	Offset   int
	Name     string
	Status   string
	Species  string
	Gender   string
	Origin   string
	Location string
	Episode  string
}

type CharacterPayload struct {
	Name     string       `bson:"name" json:"name" fake:"{firstname}"`
	Status   string       `bson:"status" json:"status" fake:"{randomstring:[Alive,unknown,Dead]}"`
	Species  string       `bson:"species" json:"species" fake:"{randomstring:[Human,Alien]}"`
	Type     string       `bson:"type" json:"type" fake:"{lastname}"`
	Gender   string       `bson:"gender" json:"gender" fake:"{gender}"`
	Origin   OriginTiny   `bson:"origin" json:"origin" fake:"{struct}"`
	Location LocationTiny `bson:"location" json:"location" fake:"{struct}"`
	Image    string       `bson:"image" json:"image" fake:"{imageurl:[20,20]}"`
	Episode  []string     `bson:"episode" json:"episode" fake:"{name}" fakesize:"3"`
	Created  time.Time    `bson:"created" json:"created" fake:"{date}"`
}

func (e *CharacterPayload) validate() (err error) {
	if e.Name == "" {
		return errors.New("missing field name")
	}
	if e.Status == "" {
		return errors.New("missing field status")
	}
	if e.Species == "" {
		return errors.New("missing field species")
	}
	if len(e.Episode) == 0 {
		return errors.New("missing field episode")
	}
	return
}

// Bind Func to use User as a payload
func (e *CharacterPayload) Bind(r *http.Request) error {
	if err := e.validate(); err != nil {
		return err
	}
	return nil

}
