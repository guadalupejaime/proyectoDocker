package models

import (
	"errors"
	"net/http"
	"time"
)

type Locations struct {
	Locations []Location `bson:"results" json:"results"`
}

func (lm *Locations) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
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

func (lm *Location) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type LocationFilters struct {
	Limit     int
	Offset    int
	Name      string
	Type      string
	Dimension string
}

type LocationPayload struct {
	Name      string   `bson:"name" json:"name" fake:"{firstname}"`
	Type      string   `bson:"type" json:"type" fake:"{lastname}"`
	Dimension string   `bson:"dimension" json:"dimension" fake:"{beername}"`
	Residents []string `bson:"residents" json:"residents" fake:"{username}" fakesize:"2"`
	URL       string   `bson:"url" json:"url" fake:"{url}"`
}

func (lp *LocationPayload) validate() (err error) {
	if lp.Name == "" {
		return errors.New("missing field name")
	}

	if lp.Dimension == "" {
		return errors.New("missing field dimension")
	}

	if lp.Type == "" {
		return errors.New("missing field type")
	}

	if len(lp.Residents) == 0 {
		return errors.New("missing field residents")
	}
	return
}

// Bind Func to use User as a payload
func (e *LocationPayload) Bind(r *http.Request) error {
	if err := e.validate(); err != nil {
		return err
	}
	return nil

}
