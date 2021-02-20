package mongo

import (
	"errors"

	"github.com/guadalupej/proyecto/pkg/models"
	"gopkg.in/mgo.v2"
)

func (r *Repository) InsertCharacter(character models.Character) error {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("characters")
	err := com.Insert(&character)
	if err != nil {
		if mgo.IsDup(err) {
			return errors.New("id already exists")
		}
		return err
	}
	return nil
}
