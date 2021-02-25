package models

import "time"

type Locations struct {
	Locations []Location `bson:"results" json:"results"`
}

type Location struct {
	ID        int       `bson:"_id" json:"id" fake:"{number:1,10}"`
	Name      string    `bson:"name" json:"name" fake:"{firstname}"`
	Type      string    `bson:"type" json:"type" fake:"{lastname}"`
	Dimension string    `bson:"dimension" json:"dimension" fake:"{beername}"`
	Residents []string  `bson:"residents" json:"residents" fake:"{username}" fakesize:"2"`
	URL       string    `bson:"url" json:"url" fake:"{url}"`
	Created   time.Time `bson:"created" json:"created" fake:"{date}"`
}

type LocationFilters struct {
	Limit     int
	Offset    int
	Name      string
	Type      string
	Dimension string
}
