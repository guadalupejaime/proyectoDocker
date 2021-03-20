package mongo

import (
	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/newerrors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (r *Repository) InsertCharacter(character models.Character) error {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("characters")
	id, err := r.getCounterCharacters()
	if err != nil {
		return err
	}
	character.ID = *id
	err = com.Insert(&character)
	if err != nil {
		if mgo.IsDup(err) {
			return newerrors.NewErrBadRequest("id already exists")
		}
		return err
	}
	return nil
}

func (r *Repository) GetCharacters(filters models.CharactersFilters) ([]models.Character, *int, error) {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("characters")

	searchQuery := getParameters(filters)

	characters := []models.Character{}
	err := com.Find(searchQuery).Limit(filters.Limit).Skip(filters.Offset).All(&characters)
	if err != nil {
		return nil, nil, err
	}
	n, err := com.Find(searchQuery).Count()
	if err != nil {
		return nil, nil, err
	}
	return characters, &n, nil
}

func (r *Repository) GetCharacterByID(id int) (*models.Character, error) {
	session := r.Session.Copy()
	defer session.Close()
	com := session.DB(r.DatabaseName).C("characters")
	characters := models.Character{}
	err := com.FindId(id).One(&characters)
	if err != nil {
		if err.Error() == "not found" {
			return nil, newerrors.NewErrNotFound("character not found")
		}
		return nil, err
	}
	return &characters, nil
}

func getParameters(filters models.CharactersFilters) *bson.M {
	searchQuery := bson.M{}
	if filters.Name != "" {
		searchQuery["name"] = filters.Name
	}
	if filters.Status != "" {
		searchQuery["status"] = filters.Status
	}
	if filters.Species != "" {
		searchQuery["species"] = filters.Species
	}
	if filters.Gender != "" {
		searchQuery["gender"] = filters.Gender
	}
	if filters.Origin != "" {
		searchQuery["origin.name"] = filters.Origin
	}
	if filters.Location != "" {
		searchQuery["location.name"] = filters.Location
	}
	if filters.Episode != "" {
		searchQuery["episode"] = filters.Episode
	}
	return &searchQuery
}
