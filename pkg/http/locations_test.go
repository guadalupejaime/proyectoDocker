package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func TestLocationController_List(t *testing.T) {
	locationsList := make([]models.Location, 5)
	for i := 0; i < 5; i++ {
		var lo models.Location
		gofakeit.Struct(&lo)

		locationsList = append(locationsList, lo)

	}

	tests := []struct {
		name             string
		locationService  LocationsService
		expectedStatus   int
		expectedResponse *models.Locations
		filters          models.LocationFilters
	}{
		{
			name:            "Status OK",
			locationService: &LocationsServiceMock{ListLocations: locationsList},
			expectedStatus:  http.StatusOK,
			expectedResponse: &models.Locations{
				Locations:     locationsList,
				TotalFound:    len(locationsList),
				TotalReturned: len(locationsList),
			},
		},
		{
			name: "Internal Server Error",
			locationService: &LocationsServiceMock{
				CodeError: 500,
				MsgError:  "Internal Server Error",
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LocationController{
				LocationsService: tt.locationService,
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
				var response *models.Locations
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("cannot unmarshal return value\n %s \n error: %v",
						rr.Body.String(), err)
				}

				// fmt.Println(response)
				if !reflect.DeepEqual(response, tt.expectedResponse) {
					// t.Errorf("response different = %v, want %v", response, tt.expectedResponse)
					t.Errorf("response different = , want ")
				}
			}
		})
	}
}

func TestLocationController_GetLocationByID(t *testing.T) {
	var location models.Location
	gofakeit.Struct(&location)

	tests := []struct {
		name             string
		locationService  LocationsService
		expectedStatus   int
		expectedResponse *models.Location
		id               int
	}{
		{
			name:             "200 - Status OK",
			locationService:  &LocationsServiceMock{ListLocations: []models.Location{location}},
			expectedStatus:   http.StatusOK,
			expectedResponse: &location,
			id:               1,
		},

		{
			name: "404 - Not Found",
			locationService: &LocationsServiceMock{
				CodeError: 404,
				MsgError:  "Not Found",
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "400 - Bad Request",
			locationService: &LocationsServiceMock{
				CodeError: 400,
				MsgError:  "Bad Request",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: " 500 - Internal Server Error",
			locationService: &LocationsServiceMock{
				CodeError: 500,
				MsgError:  "Internal Server Error",
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LocationController{
				LocationsService: tt.locationService,
			}
			req, err := http.NewRequest("GET", "", nil)
			if err != nil {
				t.Fatal(err)
			}
			ctx := context.WithValue(req.Context(), middleware.ContextKeyOffset, 10)
			req = req.WithContext(ctx)
			ctx = context.WithValue(req.Context(), middleware.ContextKeyLimit, 10)
			req = req.WithContext(ctx)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", fmt.Sprintf("%d", tt.id))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			rr := httptest.NewRecorder()

			render.SetContentType(render.ContentTypeJSON)(http.HandlerFunc(c.Get)).ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
				return
			}

			// Check return value
			if tt.expectedStatus == http.StatusOK {
				var response *models.Location
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

func TestLocationController_Insert(t *testing.T) {
	var location *models.Location
	gofakeit.Struct(&location)

	tests := []struct {
		name            string
		locationService LocationsService
		expectedStatus  int
		location        models.Location
	}{
		{
			name:            "200 - Status OK",
			locationService: &LocationsServiceMock{},
			expectedStatus:  http.StatusAccepted,
			location:        *location,
		},
		{
			name: "500 - Internal Server Error",
			locationService: &LocationsServiceMock{
				CodeError: 500,
				MsgError:  "Internal Server Error",
			},
			location:       *location,
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LocationController{
				LocationsService: tt.locationService,
			}

			location, err := json.Marshal(tt.location)
			if err != nil {
				t.Fatal(err)
			}
			body := bytes.NewBuffer(location)
			req, err := http.NewRequest("POST", "", body)
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
