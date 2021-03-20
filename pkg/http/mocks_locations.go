package http

import (
	"fmt"

	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/newerrors"
)

type LocationsServiceMock struct {
	ListLocations []models.Location
	CodeError     int
	MsgError      string
}

func (s LocationsServiceMock) GetLocations(filters models.LocationFilters) ([]models.Location, *int, error) {
	err := s.setError()
	if err != nil {
		return nil, nil, err
	}
	total := len(s.ListLocations)
	return s.ListLocations, &total, nil
}

func (s LocationsServiceMock) GetLocationByID(id int) (*models.Location, error) {
	err := s.setError()
	if err != nil {
		return nil, err
	}
	return &s.ListLocations[0], nil
}

func (s LocationsServiceMock) InsertLocation(location models.LocationPayload) error {
	err := s.setError()
	if err != nil {
		return err
	}
	return nil
}

func (e LocationsServiceMock) setError() error {
	if e.CodeError == 500 {
		return fmt.Errorf(e.MsgError)
	}
	if e.CodeError == 404 {
		return newerrors.NewErrNotFound(e.MsgError)
	}
	if e.CodeError == 400 {
		return newerrors.NewErrBadRequest(e.MsgError)
	}
	return nil
}
