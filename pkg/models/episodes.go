package models

import "time"

type Episodes struct {
	Episodes []Episode `bson:"results" json:"results"`
}

type Episode struct {
	ID         int       `bson:"_id" json:"id"`
	Name       string    `bson:"name" json:"name"`
	AirDate    string    `bson:"air_date" json:"air_date"`
	Episode    string    `bson:"episode" json:"episode"`
	Characters []string  `bson:"characters" json:"characters"`
	Created    time.Time `bson:"created" json:"created"`
}
