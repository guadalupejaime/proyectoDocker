package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/models"
)

type ContextKey struct {
	Name string
}

var (
	ContextKeyLimit  = &ContextKey{"Limit"}
	ContextKeyOffset = &ContextKey{"Offset"}
)

func Paginate(defaultLimit, maxLimit, defaultOffset int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var err error
			limitInt := defaultLimit
			limit := r.URL.Query().Get("limit")
			if limit != "" {
				limitInt, err = strconv.Atoi(limit)
				if err != nil {
					render.Render(w, r, models.ErrInvalidRequest(errors.New("limit must be an integer")))
					return
				}
			}
			if limitInt > maxLimit {
				render.Render(w, r, models.ErrInvalidRequest(fmt.Errorf("limit must be lower than %d", maxLimit)))
				return
			}

			offsetInt := defaultOffset
			offset := r.URL.Query().Get("offset")
			if offset != "" {
				offsetInt, err = strconv.Atoi(offset)
				if err != nil {
					render.Render(w, r, models.ErrInvalidRequest(errors.New("offset must be an integer")))
					return
				}
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, ContextKeyLimit, limitInt)
			ctx = context.WithValue(ctx, ContextKeyOffset, offsetInt)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
