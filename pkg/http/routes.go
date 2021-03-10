package http

import (
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/characters"
	"github.com/guadalupej/proyecto/pkg/episodes"
	"github.com/guadalupej/proyecto/pkg/locations"
	"github.com/guadalupej/proyecto/pkg/models"
	"github.com/guadalupej/proyecto/pkg/newerrors"
)

type Controller struct {
	CharacterService characters.Service
	LocationService  locations.Service
	EpisodeService   episodes.Service
}

// ListenAndServe starts the server
func ListenAndServe(controller Controller) {

	r := chi.NewRouter()

	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "File-Name"},
		ExposedHeaders:   []string{"Link", "File-Name", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// Mount "characters" controller
	r.Mount("/characters", NewCharacterController(controller.CharacterService).Routes())

	// Mount "locations" controller
	r.Mount("/locations", NewLocationController(controller.LocationService).Routes())

	// Mount "episodes" controller
	r.Mount("/episodes", NewEpisodeController(controller.EpisodeService).Routes())

	// Start service
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("error: %s", err.Error())
	}
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

func stringToInt(s string) *int {
	resp, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &resp
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
