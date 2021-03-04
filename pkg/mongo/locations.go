package mongo

import (
	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/newerrors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (r *Repository) InsertLocation(location models.Location) error {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("locations")
	err := com.Insert(&location)
	if err != nil {
		if mgo.IsDup(err) {
			return newerrors.NewErrBadRequest("id already exists")
		}
		return err
	}
	return nil
}

func (r *Repository) GetLocations(filters models.LocationFilters) ([]models.Location, error) {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("locations")
	query := bson.M{}
	if filters.Name != "" {
		query["name"] = filters.Name
	}
	if filters.Type != "" {
		query["type"] = filters.Type
	}
	if filters.Dimension != "" {
		query["dimension"] = filters.Dimension
	}
	locations := []models.Location{}
	err := com.Find(query).Limit(filters.Limit).Skip(filters.Limit).All(&locations)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *Repository) GetLocationByID(id int) (*models.Location, error) {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("locations")
	location := models.Location{}
	err := com.FindId(id).One(&location)
	if err != nil {
		if err.Error() == "not found" {
			return nil, newerrors.NewErrNotFound("location not found")
		}
		return nil, err
	}
	return &location, nil
}
