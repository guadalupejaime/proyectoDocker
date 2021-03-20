package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/http/middleware"
	"github.com/guadalupej/proyecto/pkg/models"
)

func TestEpisodeController_List(t *testing.T) {
	episodesList := make([]models.Episode, 5)
	for i := 0; i < 5; i++ {
		var episode models.Episode
		gofakeit.Struct(&episode)

		episodesList[i] = episode
	}

	tests := []struct {
		name             string
		episodeService   EpisodesService
		expectedStatus   int
		expectedResponse *models.Episodes
		filters          models.EpisodesFilters
	}{
		{
			name:           "Status OK",
			episodeService: &EpisodesServiceMock{ListEpisodes: episodesList},
			expectedStatus: http.StatusOK,
			expectedResponse: &models.Episodes{
				Episodes:      episodesList,
				TotalFound:    len(episodesList),
				TotalReturned: len(episodesList),
			},
		},
		{
			name: "error 500",
			episodeService: &EpisodesServiceMock{
				CodeError: 500,
				MsgError:  "error in storage",
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &EpisodeController{
				EpisodesService: tt.episodeService,
			}
			req, err := http.NewRequest("GET", "", nil)
			if err != nil {
				t.Fatal(err)
			}
			ctx := context.WithValue(req.Context(), middleware.ContextKeyOffset, 10)
			req = req.WithContext(ctx)
			ctx = context.WithValue(req.Context(), middleware.ContextKeyLimit, 10)
			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()

			render.SetContentType(render.ContentTypeJSON)(http.HandlerFunc(c.List)).ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Check return value
			if tt.expectedStatus == http.StatusOK {
				var response *models.Episodes
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("cannot unmarshal return value\n %s \n error: %v",
						rr.Body.String(), err)
				}

				// fmt.Println(response)
				if !reflect.DeepEqual(response, tt.expectedResponse) {
					t.Errorf("response different = %v, want %v", response, tt.expectedResponse)
				}
			}
		})
	}
}

func TestEpisodeController_Get(t *testing.T) {
	var episode models.Episode
	gofakeit.Struct(&episode)

	tests := []struct {
		name             string
		episodeService   EpisodesService
		expectedStatus   int
		expectedResponse *models.Episode
		id               string
	}{
		{
			name:             "Status OK",
			episodeService:   &EpisodesServiceMock{ListEpisodes: []models.Episode{episode}},
			expectedStatus:   http.StatusOK,
			expectedResponse: &episode,
			id:               "1",
		},
		{
			name: "error 500",
			episodeService: &EpisodesServiceMock{
				CodeError: 500,
				MsgError:  "error in storage",
			},
			expectedStatus: http.StatusInternalServerError,
			id:             "1",
		},
		{
			name: "error 404",
			episodeService: &EpisodesServiceMock{
				CodeError: 404,
				MsgError:  "episode not found",
			},
			expectedStatus: http.StatusNotFound,
			id:             "1",
		},
		{
			name: "error 400",
			episodeService: &EpisodesServiceMock{
				CodeError: 400,
				MsgError:  "id most be number",
			},
			expectedStatus: http.StatusBadRequest,
			id:             "prueba",
		},
		{
			name: "error 400 missing id",
			episodeService: &EpisodesServiceMock{
				CodeError: 400,
				MsgError:  "missing field id",
			},
			expectedStatus: http.StatusBadRequest,
			id:             "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &EpisodeController{
				EpisodesService: tt.episodeService,
			}
			req, err := http.NewRequest("GET", "", nil)
			if err != nil {
				t.Fatal(err)
			}

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()

			render.SetContentType(render.ContentTypeJSON)(http.HandlerFunc(c.Get)).ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Check return value
			if tt.expectedStatus == http.StatusOK {
				var response *models.Episode
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("cannot unmarshal return value\n %s \n error: %v",
						rr.Body.String(), err)
				}

				// fmt.Println(response)
				if !reflect.DeepEqual(response, tt.expectedResponse) {
					t.Errorf("response different = %v, want %v", response, tt.expectedResponse)
				}
			}
		})
	}
}

func TestEpisodeController_Insert(t *testing.T) {
	var episode models.EpisodePayload
	gofakeit.Struct(&episode)

	tests := []struct {
		name           string
		episodeService EpisodesService
		expectedStatus int
		episode        models.EpisodePayload
	}{
		{
			name:           "Status OK",
			episodeService: &EpisodesServiceMock{},
			expectedStatus: http.StatusAccepted,
			episode:        episode,
		},
		{
			name: "error 500",
			episodeService: &EpisodesServiceMock{
				CodeError: 500,
				MsgError:  "error in storage",
			},
			episode:        episode,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "error 400",
			episodeService: &EpisodesServiceMock{
				CodeError: 400,
				MsgError:  "episode already exists",
			},
			episode:        episode,
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &EpisodeController{
				EpisodesService: tt.episodeService,
			}

			episodeFile, err := json.Marshal(tt.episode)
			if err != nil {
				t.Fatal(err)
			}
			buff := bytes.NewBuffer(episodeFile)
			req, err := http.NewRequest("POST", "", buff)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			render.SetContentType(render.ContentTypeJSON)(http.HandlerFunc(c.Create)).ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}
