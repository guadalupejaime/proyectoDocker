package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/http/middleware"
	"github.com/guadalupej/proyecto/pkg/models"
)

type QueueBrokerService interface {
	NewMessage(content []byte, collection string) error
}

type CharactersService interface {
	GetCharacters(filters models.CharactersFilters) ([]models.Character, *int, error)
	GetCharacterByID(id int) (*models.Character, error)
	InsertCharacter(episodes models.CharacterPayload) error
}
type CharacterController struct {
	CharactersService CharactersService
	QueueBroker       QueueBrokerService
}

func NewCharacterController(characters CharactersService, qb QueueBrokerService) *CharacterController {
	return &CharacterController{
		CharactersService: characters,
		QueueBroker:       qb,
	}
}

// Routes for permissions.
func (c *CharacterController) Routes() chi.Router {
	r := chi.NewRouter()

	r.With(middleware.Paginate(100, 500, 0)).Get("/", c.List)
	r.Get("/{id}", c.Get)
	r.Post("/", c.Create)
	r.Post("/rabbit", c.CreateWithRabbit)
	return r
}

func (c *CharacterController) List(w http.ResponseWriter, r *http.Request) {
	// Get param for query
	limit := r.Context().Value(middleware.ContextKeyLimit).(int)
	offset := r.Context().Value(middleware.ContextKeyOffset).(int)
	params := r.URL.Query()
	filters := models.CharactersFilters{
		Limit:    limit,
		Offset:   offset,
		Name:     params.Get("name"),
		Status:   params.Get("status"),
		Species:  params.Get("species"),
		Gender:   params.Get("gender"),
		Origin:   params.Get("origin"),
		Location: params.Get("location"),
		Episode:  params.Get("episode"),
	}
	list, total, err := c.CharactersService.GetCharacters(filters)
	if err != nil {
		checkError(err, w, r)
		return
	}
	resp := &models.Characters{
		Characters:    list,
		TotalFound:    *total,
		TotalReturned: len(list),
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
	return
}

func (c *CharacterController) Get(w http.ResponseWriter, r *http.Request) {
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

	character, err := c.CharactersService.GetCharacterByID(*characterIDint)
	if err != nil {
		checkError(err, w, r)
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, character)
	return
}

func (c *CharacterController) Create(w http.ResponseWriter, r *http.Request) {

	// Unmarshal User Payload
	data := &models.CharacterPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}
	// character post via rabbit
	err := c.CharactersService.InsertCharacter(*data)
	if err != nil {
		checkError(err, w, r)
	}

	w.WriteHeader(http.StatusAccepted)
	return
}

func (c *CharacterController) CreateWithRabbit(w http.ResponseWriter, r *http.Request) {

	// Unmarshal User Payload
	data := &models.CharacterPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	// character post via rabbit
	content, err := json.Marshal(data)
	if err != nil {
		render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}
	err = c.QueueBroker.NewMessage(content, "characters")
	if err != nil {
		render.Render(w, r, models.ErrInternalServer(err))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	return
}
