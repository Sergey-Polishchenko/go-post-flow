package inmemory

import (
	"testing"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	storage := NewStorage()
	user, _ := storage.CreateUser(model.UserInput{Name: "Author"})

	tests := []struct {
		name    string
		input   model.PostInput
		want    *model.Post
		wantErr error
	}{
		{
			name: "Create post successfully",
			input: model.PostInput{
				Title:         "Title",
				Content:       "Content",
				AuthorID:      user.ID,
				AllowComments: true,
			},
			want: &model.Post{
				ID:            "1",
				Title:         "Title",
				Content:       "Content",
				Author:        user,
				AllowComments: true,
			},
		},
		{
			name: "Author not found",
			input: model.PostInput{
				AuthorID: "non-existent",
			},
			wantErr: errors.ErrAuthorNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.CreatePost(tt.input)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Title, got.Title)
				assert.Equal(t, tt.want.Content, got.Content)
				assert.Equal(t, tt.want.Author.ID, got.Author.ID)
			}
		})
	}
}

func TestGetPosts(t *testing.T) {
	storage := NewStorage()
	user, _ := storage.CreateUser(model.UserInput{Name: "Author"})
	post1, _ := storage.CreatePost(model.PostInput{
		Title:         "Post1",
		Content:       "Content1",
		AuthorID:      user.ID,
		AllowComments: true,
	})
	post2, _ := storage.CreatePost(model.PostInput{
		Title:         "Post2",
		Content:       "Content2",
		AuthorID:      user.ID,
		AllowComments: true,
	})

	tests := []struct {
		name    string
		limit   *int
		offset  *int
		want    []*model.Post
		wantErr bool
	}{
		{
			name:   "Get all posts",
			limit:  nil,
			offset: nil,
			want:   []*model.Post{post2, post1},
		},
		{
			name:   "Get posts with limit",
			limit:  intPtr(1),
			offset: nil,
			want:   []*model.Post{post2},
		},
		{
			name:   "Get posts with offset",
			limit:  nil,
			offset: intPtr(1),
			want:   []*model.Post{post1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.GetPosts(tt.limit, tt.offset)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
