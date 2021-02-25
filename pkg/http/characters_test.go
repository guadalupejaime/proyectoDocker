package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/go-chi/render"
	"github.com/guadalupej/proyecto/pkg/characters"
	"github.com/guadalupej/proyecto/pkg/models"
)

func TestCharacterController_List(t *testing.T) {
	charactersList := make([]models.Character, 5)
	species := []string{"Human", "Alien"}
	status := []string{"Alive", "unknown", "Dead"}
	for i := 0; i < 5; i++ {
		gofakeit.ShuffleStrings(species)
		gofakeit.ShuffleStrings(status)
		charactersList[i] = models.Character{
			ID:      int(gofakeit.Int64()),
			Name:    gofakeit.FirstName(),
			Type:    gofakeit.LastName(),
			Gender:  gofakeit.Gender(),
			Image:   gofakeit.ImageURL(120, 120),
			Status:  status[0],
			Species: species[0],
			Origin: models.OriginTiny{
				Name: gofakeit.FirstName(),
				URL:  fmt.Sprintf("%d", gofakeit.Int64()),
			},
			Location: models.LocationTiny{
				Name: gofakeit.FirstName(),
				URL:  fmt.Sprintf("%d", gofakeit.Int64()),
			},
			Episode: []string{fmt.Sprintf("%d", gofakeit.Int64())},
		}
	}
	tests := []struct {
		name             string
		characterService CharactersService
		expectedStatus   int
		expectedResponse *models.Characters
		filters          characters.Filters
	}{
		{
			name:             "simple get",
			characterService: &CharactersServiceMock{},
			expectedStatus:   http.StatusOK,
			expectedResponse: &models.Characters{Characters: charactersList},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CharacterController{
				CharactersService: tt.characterService,
			}
			req, err := http.NewRequest("GET", "/vouchers", nil)
			if err != nil {
				t.Fatal(err)
			}
			q := req.URL.Query()
			q.Add("monto", "123")
			req.URL.RawQuery = q.Encode()

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
				if !reflect.DeepEqual(response, tt.expectedResponse) {
					t.Errorf("response different = %v, want %v", response, tt.expectedResponse)
				}
			}
		})
	}
}
