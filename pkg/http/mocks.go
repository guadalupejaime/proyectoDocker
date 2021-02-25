package http

import (
	"github.com/guadalupej/proyecto/pkg/characters"
)

type CharactersServiceMock struct {
	List []characters.Character
}

func (c CharactersServiceMock) GetCharacters(filters characters.Filters) ([]characters.Character, error) {
	return c.List, nil
}
