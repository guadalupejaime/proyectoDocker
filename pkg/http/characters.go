package http

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/characters"
	"github.com/guadalupej/proyecto/pkg/http/middleware"
	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/newerrors"
)

type CharactersService interface {
	GetCharacters(filters characters.Filters) ([]characters.Character, error)
}
type CharacterController struct {
	CharactersService CharactersService
}

func NewCharacterController(characters CharactersService) *CharacterController {
	return &CharacterController{
		CharactersService: characters,
	}
}

// Routes for permissions.
func (c *CharacterController) Routes() chi.Router {
	r := chi.NewRouter()

	r.With(middleware.Paginate(100, 500, 0)).Get("/", c.List)
	return r
}

func (c *CharacterController) List(w http.ResponseWriter, r *http.Request) {
	// Get param for query
	params := r.URL.Query()
	limit := r.Context().Value(middleware.ContextKeyLimit).(int)

	offset := r.Context().Value(middleware.ContextKeyOffset).(int)

	filters := characters.Filters{
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
	list, err := c.CharactersService.GetCharacters(filters)
	if err != nil {
		checkError(err, w, r)
	}
	resp := &models.Characters{}
	for _, char := range list {
		resp.Characters = append(resp.Characters, *models.ToCharacterModel(char))
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, resp)
	return
}

func paramToInt(param string, val url.Values) *int {
	if s := val.Get(param); s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &i
	}
	return nil
}

func checkError(err error, w http.ResponseWriter, r *http.Request) {
	if _, ok := err.(*newerrors.ErrNotFound); ok {
		render.Render(w, r, models.ErrNotFound(err))

	} else if _, ok := err.(*newerrors.ErrBadRequest); ok {
		render.Render(w, r, models.ErrInvalidRequest(err))

	} else if _, ok := err.(*newerrors.ErrUnauthorized); ok {
		render.Render(w, r, models.ErrUnauthorized(err))

	} else if _, ok := err.(*newerrors.ErrForbidden); ok {
		render.Render(w, r, models.ErrForbidden(err))
	} else if _, ok := err.(*newerrors.ErrConflict); ok {
		render.Render(w, r, models.ErrConflict(err))

	} else {
		render.Render(w, r, models.ErrInternalServer(err))
	}
}
