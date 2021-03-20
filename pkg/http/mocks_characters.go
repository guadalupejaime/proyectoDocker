package http

import (
	"fmt"

	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/newerrors"
)

type CharactersServiceMock struct {
	List      []models.Character
	CodeError int
	MsgError  string
}

func (c CharactersServiceMock) GetCharacters(filters models.CharactersFilters) ([]models.Character, *int, error) {
	err := c.setError()
	if err != nil {
		return nil, nil, err
	}
	total := len(c.List)
	return c.List, &total, nil
}

func (c CharactersServiceMock) GetCharacterByID(id int) (*models.Character, error) {
	err := c.setError()
	if err != nil {
		return nil, err
	}
	return &c.List[0], nil
}

func (c CharactersServiceMock) InsertCharacter(episodes models.CharacterPayload) error {
	err := c.setError()
	if err != nil {
		return err
	}
	return nil
}

func (c CharactersServiceMock) setError() error {
	if c.CodeError == 500 {
		return fmt.Errorf(c.MsgError)
	}
	if c.CodeError == 404 {
		return newerrors.NewErrNotFound(c.MsgError)
	}
	if c.CodeError == 400 {
		return newerrors.NewErrBadRequest(c.MsgError)
	}
	return nil
}
