package models

import "time"

type Locations struct {
	Locations []Location `bson:"results" json:"results"`
}

type Location struct {
	ID        int       `bson:"_id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Type      string    `bson:"type" json:"type"`
	Dimension string    `bson:"dimension" json:"dimension"`
	Residents []string  `bson:"residents" json:"residents"`
	URL       string    `bson:"url" json:"url"`
	Created   time.Time `bson:"created" json:"created"`
}
