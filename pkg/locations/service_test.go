package locations

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/guadalupej/proyecto/pkg/models"
)

func TestGetLocations_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	locations := make([]models.Location, 0)
	for i := 0; i < 5; i++ {
		var lc models.Location
		gofakeit.Struct(&lc)

		locations = append(locations, lc)
	}
	test := []struct {
		name    string
		fields  fields
		wantErr bool
		args    models.LocationFilters
		want    []models.Location
	}{
		{
			name: "happy path",
			fields: fields{
				storage: storageMock{
					Locations: locations,
				},
			},
			want: locations,
		},
		{
			name: "error storage",
			fields: fields{
				storage: storageMock{
					errorStorage: true,
					errorMsg:     "error storage",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.storage)
			got, err := s.GetLocations(models.LocationFilters{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Service error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocationByID_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	locations := make([]models.Location, 0)
	var location models.Location
	gofakeit.Struct(&location)
	locations = append(locations, location)

	test := []struct {
		name    string
		fields  fields
		wantErr bool
		args    models.LocationFilters
		want    *models.Location
		id      int
	}{
		{
			name: "happy path",
			fields: fields{
				storage: storageMock{
					Locations: locations,
				},
			},
			want: &locations[0],
			id:   1,
		},
		{
			name: "error storage",
			fields: fields{
				storage: storageMock{
					errorStorage: true,
					errorMsg:     "error storage",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.storage)
			got, err := s.GetLocationByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetLocationByID(%v) = %v, want %v", tt.id, got, tt.want)
			}
		})
	}
}

func TestInsertLocation_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	test := []struct {
		name    string
		fields  fields
		wantErr bool
		arg     models.LocationPayload
	}{
		{
			name: "happy path",
			fields: fields{
				storage: storageMock{},
			},
		},
		{
			name: "error storage",
			fields: fields{
				storage: storageMock{
					errorStorage: true,
					errorMsg:     "error storage",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(tt.fields.storage)
			err := s.InsertLocation(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
