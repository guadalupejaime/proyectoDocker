package characters

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/guadalupej/proyecto/pkg/models"
)

func TestGetCharacters_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	characters := make([]models.Character, 0)
	for i := 0; i < 5; i++ {
		var ch models.Character
		gofakeit.Struct(&ch)

		characters = append(characters, ch)
	}
	test := []struct {
		name    string
		fields  fields
		wantErr bool
		args    models.CharactersFilters
		want    []models.Character
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
			got, err := s.GetCharacters(models.CharactersFilters{})
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

func TestGetCharacter_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	characters := make([]models.Character, 0)
	for i := 0; i < 1; i++ {
		var ch models.Character
		gofakeit.Struct(&ch)

		characters = append(characters, ch)
	}
	test := []struct {
		name    string
		fields  fields
		wantErr bool
		want    *models.Character
		id      int
	}{
		{
			name: "happy path",
			fields: fields{
				storage: storageMock{
					Characters: characters,
				},
			},
			want: &characters[0],
			id:   2,
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
			got, err := s.GetCharacterByID(tt.id)
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

func testinsertCharacter_test(t *testing.T) {
	type fields struct {
		storage storage
	}
	test := []struct {
		name    string
		fields  fields
		wantErr bool
		arg     models.Character
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
			err := s.InsertCharacter(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
