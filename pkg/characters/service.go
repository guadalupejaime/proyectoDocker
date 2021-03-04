package characters

import (
	"log"

	"github.com/guadalupej/proyecto/pkg/models"
)

// storage stores all the
type storage interface {
	// characters
	GetCharacters(filters models.CharactersFilters) ([]models.Character, error)
	GetCharacterByID(id int) (*models.Character, error)
	InsertCharacter(characters models.Character) error
}
type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetCharacters(filters models.CharactersFilters) ([]models.Character, error) {
	characters, err := s.storage.GetCharacters(filters)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return characters, nil
}

func (s Service) GetCharacterByID(id int) (*models.Character, error) {
	characters, err := s.storage.GetCharacterByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return characters, nil
}

func (s Service) InsertCharacter(characters models.Character) error {
	err := s.storage.InsertCharacter(characters)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
