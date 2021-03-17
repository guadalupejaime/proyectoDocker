package episodes

import (
	"errors"

	"github.com/guadalupej/proyecto/pkg/models"
)

type storageMock struct {
	Episodes     []models.Episode
	errorStorage bool
	errorMsg     string
}

func (s storageMock) GetEpisodes(filters models.EpisodesFilters) ([]models.Episode, *int, error) {
	if s.errorStorage {
		return nil, nil, errors.New(s.errorMsg)
	}
	return s.Episodes, nil, nil
}

func (s storageMock) GetEpisodeByID(id int) (*models.Episode, error) {
	if s.errorStorage {
		return nil, errors.New(s.errorMsg)
	}

	for _, ep := range s.Episodes {
		if id == ep.ID {
			return &ep, nil
		}
	}

	return &s.Episodes[0], nil
}

func (s storageMock) InsertEpisode(episodes models.Episode) error {
	if s.errorStorage {
		return errors.New(s.errorMsg)
	}
	return nil
}
