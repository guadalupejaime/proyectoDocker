package http

import (
	"errors"

	"github.com/guadalupej/proyecto/pkg/models"
)

type CharactersServiceMock struct {
	List  []models.Character
	Error bool
}

func (c CharactersServiceMock) GetCharacters(filters models.CharactersFilters) ([]models.Character, error) {
	if c.Error {
		return nil, errors.New("error in storage")
	}
	return c.List, nil
}
