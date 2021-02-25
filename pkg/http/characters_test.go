package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/http/middleware"
	"github.com/guadalupej/proyecto/pkg/models"
)

func TestCharacterController_List(t *testing.T) {
	charactersList := make([]models.Character, 0)

	for i := 0; i < 5; i++ {
		var ch models.Character
		gofakeit.Struct(&ch)

		charactersList = append(charactersList, ch)

	}

	tests := []struct {
		name             string
		characterService CharactersService
		expectedStatus   int
		expectedResponse *models.Characters
		filters          models.CharactersFilters
	}{
		{
			name:             "Status OK",
			characterService: &CharactersServiceMock{List: charactersList},
			expectedStatus:   http.StatusOK,
			expectedResponse: &models.Characters{Characters: charactersList},
		},
		{
			name:             "Internal Server Error",
			characterService: &CharactersServiceMock{List: charactersList, Error: true},
			expectedStatus:   http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CharacterController{
				CharactersService: tt.characterService,
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
				var response *models.Characters
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("cannot unmarshal return value\n %s \n error: %v",
						rr.Body.String(), err)
				}

				fmt.Println(response)
				if !reflect.DeepEqual(response, tt.expectedResponse) {
					t.Errorf("response different = %v, want %v", response, tt.expectedResponse)
				}
			}
		})
	}
}
