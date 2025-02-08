package inmemory

import (
	"reflect"
	"testing"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
)

func TestInMemoryStorage_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		input   model.UserInput
		want    *model.User
		wantErr bool
	}{
		{
			name:  "create user successfully",
			input: model.UserInput{Name: "John Doe"},
			want:  &model.User{ID: "user_1", Name: "John Doe"},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				storage := NewStorage()
				got, err := storage.CreateUser(tt.input)

				if (err != nil) != tt.wantErr {
					t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				}

				if got.Name != tt.want.Name {
					t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestInMemoryStorage_GetUser(t *testing.T) {
	tests := []struct {
		name          string
		existingUsers map[string]*model.User
		id            string
		want          *model.User
		wantErr       bool
	}{
		{
			name: "get existing user",
			existingUsers: map[string]*model.User{
				"user_1": {ID: "user_1", Name: "Alice"},
			},
			id:   "user_1",
			want: &model.User{ID: "user_1", Name: "Alice"},
		},
		{
			name:          "get non-existing user",
			existingUsers: map[string]*model.User{},
			id:            "user_99",
			want:          nil,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				storage := &InMemoryStorage{
					Users: tt.existingUsers,
				}

				got, err := storage.GetUser(tt.id)

				if (err != nil) != tt.wantErr {
					t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetUser() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
