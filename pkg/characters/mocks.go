package characters

import (
	"errors"

	"github.com/guadalupej/proyecto/pkg/models"
)

type storageMock struct {
	Characters   []models.Character
	errorStorage bool
	errorMsg     string
}

func (s storageMock) GetCharacters(filters models.CharactersFilters) ([]models.Character, *int, error) {
	if s.errorStorage {
		return nil, nil, errors.New(s.errorMsg)
	}
	return s.Characters, nil, nil
}

func (s storageMock) GetCharacterByID(id int) (*models.Character, error) {
	if s.errorStorage {
		return nil, errors.New(s.errorMsg)
	}
	return &s.Characters[0], nil
}

func (s storageMock) InsertCharacter(characters models.Character) error {
	if s.errorStorage {
		return errors.New(s.errorMsg)
	}
	return nil
}
