package locations

import (
	"log"

	"github.com/guadalupej/proyecto/pkg/models"
)

// storage stores all the locations
type storage interface {
	// locations
	GetLocations(filters models.LocationFilters) ([]models.Location, error)
	GetLocationByID(id int) (*models.Location, error)
	InsertLocation(location models.Location) error
}

type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetLocations(filters models.LocationFilters) ([]models.Location, error) {
	locations, err := s.storage.GetLocations(filters)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return locations, nil
}

func (s Service) GetLocationByID(id int) (*models.Location, error) {
	location, err := s.storage.GetLocationByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return location, nil
}

func (s Service) InsertLocation(location models.Location) error {
	err := s.storage.InsertLocation(location)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
