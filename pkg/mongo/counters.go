package mongo

import (
	"errors"

	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/newerrors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (r *Repository) InsertCounters(counters models.Counters) error {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("counters")
	err := com.Insert(&counters)
	if err != nil {
		if mgo.IsDup(err) {
			return newerrors.NewErrBadRequest("id already exists")
		}
		return err
	}
	return nil
}

func (r *Repository) getCounterCharacters() (*int, error) {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("counters")
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count_character": 1}},
		ReturnNew: true,
		Upsert:    false,
	}
	counterM := models.Counters{}
	info, err := com.FindId(0).Apply(change, &counterM)
	if err != nil {
		return nil, err
	}
	if info == nil || info.Matched == 0 {
		return nil, errors.New("document with the increment ID not found")
	}

	return &counterM.CountCharacters, nil
}

func (r *Repository) getCounterLocation() (*int, error) {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("counters")
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count_location": 1}},
		ReturnNew: true,
		Upsert:    false,
	}
	counterM := models.Counters{}
	info, err := com.FindId(0).Apply(change, &counterM)
	if err != nil {
		return nil, err
	}
	if info == nil || info.Matched == 0 {
		return nil, errors.New("document with the increment ID not found")
	}

	return &counterM.CountLocation, nil
}

func (r *Repository) getCounterEpisodes() (*int, error) {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("counters")
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count_episode": 1}},
		ReturnNew: true,
		Upsert:    false,
	}
	counterM := models.Counters{}
	info, err := com.FindId(0).Apply(change, &counterM)
	if err != nil {
		return nil, err
	}
	if info == nil || info.Matched == 0 {
		return nil, errors.New("document with the increment ID not found")
	}

	return &counterM.CountEpisode, nil
}
