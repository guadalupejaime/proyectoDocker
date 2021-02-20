package mongo

import (
	"errors"

	"github.com/guadalupej/proyecto/pkg/models"
	"gopkg.in/mgo.v2"
)

func (r *Repository) InsertEpisode(episode models.Episode) error {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("episodes")
	err := com.Insert(&episode)
	if err != nil {
		if mgo.IsDup(err) {
			return errors.New("id already exists")
		}
		return err
	}
	return nil
}
