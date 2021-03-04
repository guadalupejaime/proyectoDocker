package models

import (
	"net/http"
	"time"
)

type Characters struct {
	Characters []Character `bson:"results" json:"results"`
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
