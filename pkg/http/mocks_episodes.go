package http

import (
	"fmt"

	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/newerrors"
)

type EpisodesServiceMock struct {
	CodeError    int
	MsgError     string
	ListEpisodes []models.Episode
}

func (e EpisodesServiceMock) GetEpisodes(filters models.EpisodesFilters) ([]models.Episode, *int, error) {
	err := e.setError()
	if err != nil {
		return nil, nil, err
	}
	total := len(e.ListEpisodes)
	return e.ListEpisodes, &total, nil
}

func (e EpisodesServiceMock) GetEpisodeByID(id int) (*models.Episode, error) {
	err := e.setError()
	if err != nil {
		return nil, err
	}
	return &e.ListEpisodes[0], nil
}

func (e EpisodesServiceMock) InsertEpisode(episodes models.EpisodePayload) error {
	err := e.setError()
	if err != nil {
		return err
	}
	return nil
}

func (e EpisodesServiceMock) setError() error {
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
