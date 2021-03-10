package episodes

import (
	"log"
	"time"

	"github.com/guadalupej/proyecto/pkg/models"
)

// storage stores all the
type storage interface {
	// episodes
	GetEpisodes(filters models.EpisodesFilters) ([]models.Episode, error)
	GetEpisodeByID(id int) (*models.Episode, error)
	InsertEpisode(episodes models.Episode) error
}

type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetEpisodes(filters models.EpisodesFilters) ([]models.Episode, error) {
	episodes, err := s.storage.GetEpisodes(filters)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return episodes, nil
}

func (s Service) GetEpisodeByID(id int) (*models.Episode, error) {
	episodes, err := s.storage.GetEpisodeByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return episodes, nil
}

func (s Service) InsertEpisode(episodes models.EpisodePayload) error {

	newEpisode := models.Episode{
		Name:       episodes.Name,
		AirDate:    episodes.AirDate,
		Episode:    episodes.Episode,
		Characters: episodes.Characters,
		Created:    time.Now(),
	}

	err := s.storage.InsertEpisode(newEpisode)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
