package mongo

import (
	"errors"

	"github.com/guadalupej/proyecto/pkg/models"
	"gopkg.in/mgo.v2"
)

func (r *Repository) InsertLocation(location models.Location) error {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("locations")
	err := com.Insert(&location)
	if err != nil {
		if mgo.IsDup(err) {
			return errors.New("id already exists")
		}
		return err
	}
	return nil
}
