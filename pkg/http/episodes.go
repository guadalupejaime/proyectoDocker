package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/http/middleware"
	"github.com/guadalupej/proyecto/pkg/models"
)

type EpisodesService interface {
	GetEpisodes(filters models.EpisodesFilters) ([]models.Episode, *int, error)
	GetEpisodeByID(id int) (*models.Episode, error)
	InsertEpisode(episodes models.EpisodePayload) error
}
type EpisodeController struct {
	EpisodesService EpisodesService
}

func NewEpisodeController(episodeService EpisodesService) *EpisodeController {
	return &EpisodeController{
		EpisodesService: episodeService,
	}
}

// Routes for permissions.
func (c *EpisodeController) Routes() chi.Router {
	r := chi.NewRouter()

	r.With(middleware.Paginate(100, 500, 0)).Get("/", c.List)
	r.Get("/{id}", c.Get)
	r.Post("/", c.Create)
	return r
}

func (c *EpisodeController) List(w http.ResponseWriter, r *http.Request) {
	// Get param for query
	limit := r.Context().Value(middleware.ContextKeyLimit).(int)
	offset := r.Context().Value(middleware.ContextKeyOffset).(int)
	params := r.URL.Query()
	filters := models.EpisodesFilters{
		Limit:   limit,
		Offset:  offset,
		Name:    params.Get("name"),
		Episode: params.Get("episode"),
	}
	list, total, err := c.EpisodesService.GetEpisodes(filters)
	if err != nil {
		checkError(err, w, r)
		return
	}
	resp := &models.Episodes{
		Episodes:      list,
		TotalFound:    *total,
		TotalReturned: len(list),
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
	return
}

func (c *EpisodeController) Get(w http.ResponseWriter, r *http.Request) {
	// Get param for query
	id := chi.URLParam(r, "id")
	if id == "" {
		render.Render(w, r, models.ErrInvalidRequest(errors.New("missing id")))
		return
	}

	episodeIDint := stringToInt(id)
	if episodeIDint == nil {
		render.Render(w, r, models.ErrInvalidRequest(errors.New("id most be number")))
		return
	}

	episode, err := c.EpisodesService.GetEpisodeByID(*episodeIDint)
	if err != nil {
		checkError(err, w, r)
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, episode)
	return
}

func (c *EpisodeController) Create(w http.ResponseWriter, r *http.Request) {

	// Unmarshal User Payload
	data := &models.EpisodePayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	err := c.EpisodesService.InsertEpisode(*data)
	if err != nil {
		checkError(err, w, r)
	}

	w.WriteHeader(http.StatusAccepted)
	return
}
