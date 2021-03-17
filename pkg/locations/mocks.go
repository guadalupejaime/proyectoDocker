package locations

import (
	"errors"

	"github.com/guadalupej/proyecto/pkg/models"
)

type storageMock struct {
	Locations    []models.Location
	errorStorage bool
	errorMsg     string
}

func (s storageMock) GetLocations(filters models.LocationFilters) ([]models.Location, *int, error) {
	if s.errorStorage {
		return nil, nil, errors.New(s.errorMsg)
	}
	return s.Locations, nil, nil
}

func (s storageMock) GetLocationByID(id int) (*models.Location, error) {
	if s.errorStorage {
		return nil, errors.New(s.errorMsg)
	}
	return &s.Locations[0], nil
}

func (s storageMock) InsertLocation(location models.Location) error {
	if s.errorStorage {
		return errors.New(s.errorMsg)
	}
	return nil
}
