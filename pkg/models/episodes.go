package models

import (
	"errors"
	"net/http"
	"time"
)

type Episodes struct {
	Info          Info      `bson:"info,omitempty" json:"info,omitempty"`
	Episodes      []Episode `bson:"results" json:"results"`
	TotalFound    int       `bson:"total_found" json:"total_found"`
	TotalReturned int       `bson:"total_returned" json:"total_returned"`
}

func (mt *Episodes) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Episode struct {
	ID         int       `bson:"_id" json:"id" fake:"{number:1,10}"`
	Name       string    `bson:"name" json:"name" fake:"{firstname}"`
	AirDate    string    `bson:"air_date" json:"air_date" fake:"{firstname}"`
	Episode    string    `bson:"episode" json:"episode" fake:"{firstname}"`
	Characters []string  `bson:"characters" json:"characters" fake:"{username}" fakesize:"2"`
	Created    time.Time `bson:"created" json:"created" fake:"{date}"`
}

func (mt *Episode) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type EpisodesFilters struct {
	Limit   int
	Offset  int
	Name    string
	Episode string
}

type EpisodePayload struct {
	Name       string    `bson:"name" json:"name" fake:"{firstname}"`
	AirDate    string    `bson:"air_date" json:"air_date" fake:"{firstname}"`
	Episode    string    `bson:"episode" json:"episode" fake:"{firstname}"`
	Characters []string  `bson:"characters" json:"characters" fake:"{username}" fakesize:"2"`
	Created    time.Time `bson:"created" json:"created" fake:"{date}"`
}

func (e *EpisodePayload) validate() (err error) {
	if e.Name == "" {
		return errors.New("missing field name")
	}
	if e.AirDate == "" {
		return errors.New("missing field air_date")
	}
	if e.Episode == "" {
		return errors.New("missing field episode")
	}
	if len(e.Characters) == 0 {
		return errors.New("missing field characters")
	}
	return
}

// Bind Func to use User as a payload
func (e *EpisodePayload) Bind(r *http.Request) error {
	if err := e.validate(); err != nil {
		return err
	}
	return nil
}
