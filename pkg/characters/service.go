package characters

import "log"

// storage stores all the
type storage interface {
	GetCharacters(filters Filters) ([]Character, error)
}
type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetCharacters(filters Filters) ([]Character, error) {
	characters, err := s.storage.GetCharacters(filters)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return characters, nil
}
