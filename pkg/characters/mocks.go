package characters

import (
	"errors"
)

type storageMock struct {
	Characters   []Character
	errorStorage bool
	errorMsg     string
}

func (s storageMock) GetCharacters(filters Filters) ([]Character, error) {
	if s.errorStorage {
		return nil, errors.New(s.errorMsg)
	}
	return s.Characters, nil
}
