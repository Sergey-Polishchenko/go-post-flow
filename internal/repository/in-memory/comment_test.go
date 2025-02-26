package inmemory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Sergey-Polishchenko/go-post-flow/internal/delivery/graph/model"
	"github.com/Sergey-Polishchenko/go-post-flow/pkg/errors"
)

func TestCreateComment(t *testing.T) {
	storage := NewStorage()
	user, _ := storage.CreateUser(model.UserInput{Name: "Author"})
	post, _ := storage.CreatePost(model.PostInput{
		Title:         "Post",
		Content:       "Content",
		AuthorID:      user.ID,
		AllowComments: true,
	})

	tests := []struct {
		name    string
		input   model.CommentInput
		want    *model.Comment
		wantErr error
	}{
		{
			name: "Create comment successfully",
			input: model.CommentInput{
				Text:     "Comment",
				AuthorID: user.ID,
				PostID:   post.ID,
			},
			want: &model.Comment{
				ID:     "1",
				Text:   "Comment",
				Author: user,
				Post:   post,
			},
		},
		{
			name: "Author not found",
			input: model.CommentInput{
				AuthorID: "non-existent",
			},
			wantErr: errors.ErrAuthorNotFound,
		},
		{
			name: "Post not found",
			input: model.CommentInput{
				AuthorID: user.ID,
				PostID:   "non-existent",
			},
			wantErr: errors.ErrPostNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.CreateComment(tt.input)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Text, got.Text)
				assert.Equal(t, tt.want.Author.ID, got.Author.ID)
				assert.Equal(t, tt.want.Post.ID, got.Post.ID)
			}
		})
	}
}

func TestGetCommentsByIDs(t *testing.T) {
	storage := NewStorage()
	user, _ := storage.CreateUser(model.UserInput{Name: "Author"})
	post, _ := storage.CreatePost(model.PostInput{
		Title:         "Post",
		Content:       "Content",
		AuthorID:      user.ID,
		AllowComments: true,
	})
	comment1, _ := storage.CreateComment(model.CommentInput{
		Text:     "Comment1",
		AuthorID: user.ID,
		PostID:   post.ID,
	})
	comment2, _ := storage.CreateComment(model.CommentInput{
		Text:     "Comment2",
		AuthorID: user.ID,
		PostID:   post.ID,
	})

	tests := []struct {
		name    string
		ids     []string
		want    []*model.Comment
		wantErr bool
	}{
		{
			name: "Get comments by IDs",
			ids:  []string{comment1.ID, comment2.ID},
			want: []*model.Comment{comment1, comment2},
		},
		{
			name:    "Non-existent comment ID",
			ids:     []string{"non-existent"},
			want:    []*model.Comment{nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.GetCommentsByIDs(context.Background(), tt.ids)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetCommentsIDs(t *testing.T) {
	storage := NewStorage()
	user, _ := storage.CreateUser(model.UserInput{Name: "Author"})
	post, _ := storage.CreatePost(model.PostInput{
		Title:         "Post",
		Content:       "Content",
		AuthorID:      user.ID,
		AllowComments: true,
	})
	comment1, _ := storage.CreateComment(model.CommentInput{
		Text:     "Comment1",
		AuthorID: user.ID,
		PostID:   post.ID,
	})
	comment2, _ := storage.CreateComment(model.CommentInput{
		Text:     "Comment2",
		AuthorID: user.ID,
		PostID:   post.ID,
	})

	tests := []struct {
		name    string
		postID  string
		want    []string
		wantErr error
	}{
		{
			name:   "Get comment IDs",
			postID: post.ID,
			want:   []string{comment1.ID, comment2.ID},
		},
		{
			name:    "Post not found",
			postID:  "non-existent",
			wantErr: errors.ErrPostNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.GetCommentsIDs(tt.postID)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetChildrenIDs(t *testing.T) {
	storage := NewStorage()
	user, _ := storage.CreateUser(model.UserInput{Name: "Author"})
	post, _ := storage.CreatePost(model.PostInput{
		Title:         "Post",
		Content:       "Content",
		AuthorID:      user.ID,
		AllowComments: true,
	})
	parentComment, _ := storage.CreateComment(model.CommentInput{
		Text:     "Parent",
		AuthorID: user.ID,
		PostID:   post.ID,
	})
	childComment, _ := storage.CreateComment(model.CommentInput{
		Text:     "Child",
		AuthorID: user.ID,
		PostID:   post.ID,
		ParentID: &parentComment.ID,
	})

	tests := []struct {
		name      string
		commentID string
		want      []string
		wantErr   error
	}{
		{
			name:      "Get children IDs",
			commentID: parentComment.ID,
			want:      []string{childComment.ID},
		},
		{
			name:      "Comment not found",
			commentID: "non-existent",
			wantErr:   errors.ErrCommentNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := storage.GetChildrenIDs(tt.commentID)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
