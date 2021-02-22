package characters

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestGetCharacters_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	characters := make([]Character, 5)
	species := []string{"Human", "Alien"}
	status := []string{"Alive", "unknown", "Dead"}
	for i := 0; i < 5; i++ {
		gofakeit.ShuffleStrings(species)
		gofakeit.ShuffleStrings(status)
		characters[i] = Character{
			ID:      int(gofakeit.Int64()),
			Name:    gofakeit.FirstName(),
			Type:    gofakeit.LastName(),
			Gender:  gofakeit.Gender(),
			Image:   gofakeit.ImageURL(120, 120),
			Status:  status[0],
			Species: species[0],
			Origin: OriginTiny{
				Name: gofakeit.FirstName(),
				URL:  fmt.Sprintf("%d", gofakeit.Int64()),
			},
			Location: LocationTiny{
				Name: gofakeit.FirstName(),
				URL:  fmt.Sprintf("%d", gofakeit.Int64()),
			},
			Episode: []string{fmt.Sprintf("%d", gofakeit.Int64())},
		}
	}
	test := []struct {
		name    string
		fields  fields
		wantErr bool
		args    Filters
		want    []Character
	}{
		{
			name: "happy path",
			fields: fields{
				storage: storageMock{
					Characters: characters,
				},
			},
			want: characters,
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
			got, err := s.GetCharacters(Filters{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Service error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
