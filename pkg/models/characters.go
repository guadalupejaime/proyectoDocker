package models

import "time"

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
