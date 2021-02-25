package models

import (
	"net/http"
	"time"
)

type Episodes struct {
	Episodes []Episode `bson:"results" json:"results"`
}

type Episode struct {
	ID         int       `bson:"_id" json:"id" fake:"{number:1,10}"`
	Name       string    `bson:"name" json:"name" fake:"{firstname}"`
	AirDate    string    `bson:"air_date" json:"air_date" fake:"{firstname}"`
	Episode    string    `bson:"episode" json:"episode" fake:"{firstname}"`
	Characters []string  `bson:"characters" json:"characters" fake:"{username}" fakesize:"2"`
	Created    time.Time `bson:"created" json:"created" fake:"{date}"`
}

func (mt *Episodes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type EpisodesFilters struct {
	Limit   int
	Offset  int
	Name    string
	Episode string
}
