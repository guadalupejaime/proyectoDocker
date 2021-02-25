package episodes

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/guadalupej/proyecto/pkg/models"
)

func TestGetEpisodes_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	episodes := make([]models.Episode, 0)
	for i := 0; i < 5; i++ {
		var ch models.Episode
		gofakeit.Struct(&ch)

		episodes = append(episodes, ch)
	}
	test := []struct {
		name    string
		fields  fields
		wantErr bool
		args    models.EpisodesFilters
		want    []models.Episode
	}{
		{
			name: "happy path",
			fields: fields{
				storage: storageMock{
					Episodes: episodes,
				},
			},
			want: episodes,
		},
		{
			name: "error storage",
			fields: fields{
				storage: storageMock{
					errorStorage: true,
					errorMsg:     "error storage",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.storage)
			got, err := s.GetEpisodes(models.EpisodesFilters{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Service error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetEpisodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEpisodeByID_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	episodes := make([]models.Episode, 0)

	var episode models.Episode

	for i := 0; i < 5; i++ {
		var ch models.Episode
		gofakeit.Struct(&ch)

		if episode.ID == 0 && gofakeit.Int8()%2 == 0 {
			episode = ch
		}

		episodes = append(episodes, ch)
	}
	test := []struct {
		name    string
		fields  fields
		wantErr bool
		argsID  int
		want    *models.Episode
	}{
		{
			name: "happy path",
			fields: fields{
				storage: storageMock{
					Episodes: episodes,
				},
			},
			argsID: episode.ID,
			want:   &episode,
		},
		{
			name: "error storage",
			fields: fields{
				storage: storageMock{
					errorStorage: true,
					errorMsg:     "error storage",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.storage)
			got, err := s.GetEpisodeByID(tt.argsID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetEpisodeByID(%d) = %v, want %v", episode.ID, got, tt.want)
			}
		})
	}
}
