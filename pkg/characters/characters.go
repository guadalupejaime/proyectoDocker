package characters

import "time"

type Character struct {
	ID       int
	Name     string
	Status   string
	Species  string
	Type     string
	Gender   string
	Origin   OriginTiny
	Location LocationTiny
	Image    string
	Episode  []string
	Created  time.Time
}

type OriginTiny struct {
	Name string
	URL  string
}
type LocationTiny struct {
	Name string
	URL  string
}
type Filters struct {
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
