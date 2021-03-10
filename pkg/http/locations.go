package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/http/middleware"
	"github.com/guadalupej/proyecto/pkg/models"
)

type LocationsService interface {
	GetLocations(filters models.LocationFilters) ([]models.Location, error)
	GetLocationByID(id int) (*models.Location, error)
	InsertLocation(location models.LocationPayload) error
}

type LocationController struct {
	LocationsService LocationsService
}

func NewLocationController(locations LocationsService) *LocationController {
	return &LocationController{
		LocationsService: locations,
	}
}

// Routes for permissions.
func (c *LocationController) Routes() chi.Router {
	r := chi.NewRouter()

	r.With(middleware.Paginate(100, 500, 0)).Get("/", c.List)
	return r
}

func (c *LocationController) List(w http.ResponseWriter, r *http.Request) {
	// Get param for query
	limit := r.Context().Value(middleware.ContextKeyLimit).(int)
	offset := r.Context().Value(middleware.ContextKeyOffset).(int)
	params := r.URL.Query()
	filters := models.LocationFilters{
		Limit:     limit,
		Offset:    offset,
		Name:      params.Get("name"),
		Type:      params.Get("type"),
		Dimension: params.Get("dimension"),
	}
	list, err := c.LocationsService.GetLocations(filters)
	if err != nil {
		checkError(err, w, r)
	}
	resp := &models.Locations{Locations: list}

	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
	return
}

func (c *LocationController) Get(w http.ResponseWriter, r *http.Request) {
	// Get param for query
	id := chi.URLParam(r, "id")
	if id == "" {
		render.Render(w, r, models.ErrInvalidRequest(errors.New("missing id")))
		return
	}

	characterIDint := stringToInt(id)
	if characterIDint == nil {
		render.Render(w, r, models.ErrInvalidRequest(errors.New("id most be number")))
		return
	}

	location, err := c.LocationsService.GetLocationByID(*characterIDint)
	if err != nil {
		checkError(err, w, r)
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, location)
	return
}

func (c *LocationController) Create(w http.ResponseWriter, r *http.Request) {

	// Unmarshal User Payload
	data := &models.LocationPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	err := c.LocationsService.InsertLocation(*data)
	if err != nil {
		checkError(err, w, r)
	}

	w.WriteHeader(http.StatusAccepted)
	return
}
